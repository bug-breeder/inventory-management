package handlers

import (
	"database/sql"
	"net/http"
	"product_app/models"
	"product_app/utils"
	"strconv"
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
		Title      string
		Template   string
		Products   []models.Product
		Categories []models.ProductCategory
	}{
		Title:      "Product List",
		Template:   "list-products",
		Products:   products,
		Categories: categories,
	}

	utils.RenderTemplate(w, "base.html", data)
}

func (h *ProductHandler) ShowAddProductPage(w http.ResponseWriter, r *http.Request) {
	var categories []models.ProductCategory
	rows, err := h.DB.Query("SELECT * FROM product_categories")
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
		Title      string
		Template   string
		Categories []models.ProductCategory
	}{
		Title:      "Add Product",
		Template:   "add-product",
		Categories: categories,
	}

	utils.RenderTemplate(w, "base.html", data)
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	p.ProductName = r.FormValue("product_name")
	p.UnitPrice, _ = strconv.ParseFloat(r.FormValue("unit_price"), 64)
	p.Unit = r.FormValue("unit")
	p.Weight, _ = strconv.ParseFloat(r.FormValue("weight"), 64)
	p.CategoryID, _ = strconv.Atoi(r.FormValue("category_id"))
	p.Status = r.FormValue("status") == "true"

	sqlStatement := `INSERT INTO products (product_name, unit_price, unit, weight, category_id, status) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := h.DB.Exec(sqlStatement, p.ProductName, p.UnitPrice, p.Unit, p.Weight, p.CategoryID, p.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
