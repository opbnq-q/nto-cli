package model

import "go/ast"

type Model struct {
	Name   string
	Fields []Field
}

type Field struct {
	Name     string
	Type     ast.Expr
	Tag      string
	Metadata FieldMetadata
}

type FieldMetadata struct {
	Hidden          bool
	Readonly        bool
	Label           string
	IsSlice         bool
	IsPrimitiveType bool
	IsRelatedModel  bool
	RelatedModel    string
	Datatype        string
	RelatedFields   []string
}
