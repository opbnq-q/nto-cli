package model

import (
	"github.com/opbnq-q/nto-cli/utils"
	"go/ast"
	"go/parser"
	"go/token"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/kuzgoga/fogg"
)

func ParseModelsPackage(modelsPkgDir string) ([]Model, error) {
	var models []Model
	fileset := token.NewFileSet()

	files, err := os.ReadDir(modelsPkgDir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		filepathPath := filepath.Join(modelsPkgDir, file.Name())
		fileAst, err := parser.ParseFile(fileset, filepathPath, nil, parser.DeclarationErrors)

		if err != nil {
			return nil, err
		}

		models = append(models, parseFileStructs(fileset, fileAst.Decls)...)
	}

	ParseRelatedModels(&models)

	return models, nil
}

func parseFileStructs(fileSet *token.FileSet, decls []ast.Decl) []Model {
	var models []Model
	for _, decl := range decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}

		if genDecl.Tok != token.TYPE {
			continue
		}

		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			if _, ok := typeSpec.Type.(*ast.StructType); ok && typeSpec.Name != nil {
				models = append(models, parseModelDecl(fileSet, typeSpec))
			}
		}

	}
	return models
}

func parseModelDecl(fileset *token.FileSet, decl *ast.TypeSpec) Model {
	name := decl.Name.Name
	structType, _ := decl.Type.(*ast.StructType)
	var fields []Field

	for _, fieldDecl := range structType.Fields.List {
		fieldPos := fileset.Position(fieldDecl.Pos()).String()
		if len(fieldDecl.Names) == 0 {
			slog.Error("%s Embedded structure isn't supported", fieldPos)
			continue
		}

		var tag string
		if fieldDecl.Tag != nil {
			tag = fieldDecl.Tag.Value[1 : len(fieldDecl.Tag.Value)-1]
		}

		storage, err := fogg.Parse(tag)
		if err != nil {
			slog.Error("%s Struct tag parsing error: %s", fieldPos, err)
		}

		var metadata FieldMetadata

		if storage.HasTag("ui") {
			uiTag := storage.GetTag("ui")
			metadata.Hidden = uiTag.HasOption("hidden")
			metadata.Label = uiTag.GetParamOr("label", "")
			metadata.IsSlice = utils.IsSlice(fieldDecl.Type)
			metadata.IsPrimitiveType = utils.IsPrimitiveType(fieldDecl.Type)
			metadata.Readonly = uiTag.HasOption("readonly")
			metadata.Datatype = uiTag.GetParamOr("datatype", "")
			if uiTag.HasParam("field") {
				metadata.RelatedFields = strings.Split(uiTag.GetParam("field").Value, ".")
			}
		} else {
			slog.Warn("%s Field does not have a UI tag", fieldPos)
		}

		field := Field{
			Name:     fieldDecl.Names[0].Name,
			Type:     fieldDecl.Type,
			Tag:      tag,
			Metadata: metadata,
		}
		fields = append(fields, field)
	}
	return Model{
		Name:   name,
		Fields: fields,
	}
}

func ParseRelatedModels(models *[]Model) {
	for i := range *models {
		model := &(*models)[i]
		for j := range model.Fields {
			field := &model.Fields[j]
			if field.Metadata.IsPrimitiveType {
				continue
			}

			relatedModelName := utils.ResolveBaseType(field.Type)
			if relatedModelName == nil {
				slog.Error("Failed to resolve base type for field `%s` in model `%s`", field.Name, model.Name)
				continue
			}

			found := false
			for _, m := range *models {
				if m.Name == *relatedModelName {
					field.Metadata.RelatedModel = m.Name
					field.Metadata.IsRelatedModel = true
					found = true
					break
				}
			}

			if !found {
				slog.Error("Cannot classify field type `%s` in model `%s`, type `%s`",
					field.Name, model.Name, *relatedModelName)
				continue
			}
		}
	}
}
