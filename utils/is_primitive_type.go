package utils

import "go/ast"

var primitives = map[string]bool{
	"bool":       true,
	"string":     true,
	"int":        true,
	"int8":       true,
	"int16":      true,
	"int32":      true,
	"int64":      true,
	"uint":       true,
	"uint8":      true,
	"uint16":     true,
	"uint32":     true,
	"uint64":     true,
	"uintptr":    true,
	"byte":       true,
	"rune":       true,
	"float32":    true,
	"float64":    true,
	"complex64":  true,
	"complex128": true,
}

func IsPrimitiveType(expr ast.Expr) bool {
	if ident, ok := expr.(*ast.Ident); ok {
		return primitives[ident.Name]
	}

	if _, ok := expr.(*ast.ArrayType); ok {
		// Arrays and slices are not primitives
		return false
	}

	if _, ok := expr.(*ast.MapType); ok {
		// Maps are not primitives
		return false
	}

	if _, ok := expr.(*ast.ChanType); ok {
		// Channels are not primitives
		return false
	}

	if _, ok := expr.(*ast.StructType); ok {
		// Structs are not primitives
		return false
	}

	if _, ok := expr.(*ast.StarExpr); ok {
		// Handle pointers
		return IsPrimitiveType(expr.(*ast.StarExpr).X)
	}

	if _, ok := expr.(*ast.SelectorExpr); ok {
		// Handle selector expressions (like pkg.Type)
		return false
	}

	return false
}
