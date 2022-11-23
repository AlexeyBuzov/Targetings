package models

type Node struct {
	Type     string
	Logic    *string
	Value    *Targeting
	Children *[]Node
}
