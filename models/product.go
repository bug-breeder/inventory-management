package models

type Product struct {
	ID          int     `json:"id"`
	ProductCode string  `json:"product_code"`
	ProductName string  `json:"product_name"`
	UnitPrice   float64 `json:"unit_price"`
	Unit        string  `json:"unit"`
	Weight      float64 `json:"weight"`
	CategoryID  int     `json:"category_id"`
	Status      bool    `json:"status"`
	ImageURL    string  `json:"image_url"`
	Description string  `json:"description"`
}

