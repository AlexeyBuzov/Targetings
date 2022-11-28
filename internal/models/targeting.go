package models

type Targeting struct {
	Type      string
	Value     []int64
	Exception *[]int64
}
