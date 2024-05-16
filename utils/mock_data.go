package utils

import (
	"database/sql"
)

func InsertMockData(db *sql.DB) error {
	categories := []struct {
		ID           int
		CategoryName string
	}{
		{1, "Điện tử"},
		{2, "Thời trang"},
		{3, "Văn phòng phẩm"},
	}

	products := []struct {
		ProductName string
		UnitPrice   float64
		Unit        string
		Weight      float64
		CategoryID  int
		Status      bool
	}{
		{"Laptop Acer Swift 3", 699.99, "cái", 2.0, 1, true},
		{"Pc Phong Vũ 1", 1299.99, "cái", 10, 1, true},
		{"Aó khoác nam", 9.99, "cái", 0.2, 2, true},
		{"Sổ tay Teko", 19.99, "quyển", 0.5, 3, true},
	}

	for _, category := range categories {
		_, err := db.Exec("INSERT INTO product_categories (id, category_name) VALUES ($1, $2) ON CONFLICT DO NOTHING", category.ID, category.CategoryName)
		if err != nil {
			return err
		}
	}

	for _, product := range products {
		_, err := db.Exec("INSERT INTO products (product_name, unit_price, unit, weight, category_id, status) VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT DO NOTHING",
			product.ProductName, product.UnitPrice, product.Unit, product.Weight, product.CategoryID, product.Status)
		if err != nil {
			return err
		}
	}

	return nil
}

// Delete all data from the database
func DeleteMockData(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM products")
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM product_categories")
	if err != nil {
		return err
	}

	return nil
}
