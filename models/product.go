package models

type Product struct {
	ID          int     `json:"id"`
	ProductName string  `json:"product_name"`
	UnitPrice   float64 `json:"unit_price"`
	Unit        string  `json:"unit"`
	Weight      float64 `json:"weight"`
	CategoryID  int     `json:"category_id"`
	Status      bool    `json:"status"`
}
