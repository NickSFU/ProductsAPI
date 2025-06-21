package models

type Measure struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Product struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	UnitCost float64 `json:"unit_cost"`
	Measure  uint    `json:"measure"`
}
