package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"product_app/models"
	"product_app/utils"
)

type ProductHandler struct {
	DB *sql.DB
}

func (h *ProductHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	rows, err := h.DB.Query("SELECT * FROM products")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.ProductName, &p.UnitPrice, &p.Unit, &p.Weight, &p.CategoryID, &p.Status); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		products = append(products, p)
	}

	var categories []models.ProductCategory
	rows, err = h.DB.Query("SELECT * FROM product_categories")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var c models.ProductCategory
		if err := rows.Scan(&c.ID, &c.CategoryName); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		categories = append(categories, c)
	}

	data := struct {
		Products   []models.Product
		Categories []models.ProductCategory
	}{
		Products:   products,
		Categories: categories,
	}

	utils.RenderTemplate(w, "list_products.html", data)
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	rows, err := h.DB.Query("SELECT * FROM products")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.ProductName, &p.UnitPrice, &p.Unit, &p.Weight, &p.CategoryID, &p.Status); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		products = append(products, p)
	}

	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sqlStatement := `INSERT INTO products (product_name, unit_price, unit, weight, category_id, status) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := h.DB.Exec(sqlStatement, p.ProductName, p.UnitPrice, p.Unit, p.Weight, p.CategoryID, p.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
