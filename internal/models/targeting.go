package models

type Targeting struct {
	Location      *LocationTargeting
	Microcategory *MicrocategoryTargeting
	Interest      *InterestTargeting
}

type LocationTargeting struct {
	Include []int64
	Exclude *[]int64
}

type MicrocategoryTargeting struct {
	Include []int64
	Exclude *[]int64
}

type InterestTargeting struct {
	Include []Interest
	Exclude *[]Interest
}

type Interest struct {
	Id    int64
	Power string
}
