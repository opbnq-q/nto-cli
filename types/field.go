package types

type Field struct {
	Name string
	Type string
	Medatada []Medatada
}

type Medatada struct {
	Name string
	Values []string
}