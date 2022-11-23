package models

type Targeting struct {
	Type     string
	Value    int64
	Children *[]Node
}
