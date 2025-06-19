package models

type Measure struct {
	ID   uint
	Name string
}

type Product struct {
	ID       uint
	Name     string
	Quantity int
	UnitCost float64
	Measure  uint
}
