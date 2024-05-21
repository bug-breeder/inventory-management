package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"product_app/models"
	"product_app/utils"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
)

type ProductHandler struct {
	DB *sql.DB
}

func (h *ProductHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
	// Read query parameters
	productIDStr := r.URL.Query().Get("product_id")
	productName := r.URL.Query().Get("product_name")
	categoryIDStr := r.URL.Query().Get("category_id")
	statusStr := r.URL.Query().Get("status")

	// Construct the base query
	query := `SELECT id, product_name, unit_price, unit, weight, category_id, status FROM products WHERE 1=1`
	var args []interface{}
	var conditions []string

	// Add conditions based on the provided filters
	if productIDStr != "" {
		productID, err := strconv.ParseInt(productIDStr, 10, 32)
		if err != nil {
			log.Printf("Error converting status to int: %v\n", err)
		} else {
			conditions = append(conditions, "id = $")
			args = append(args, productID)
		}
	}
	if productName != "" {
		conditions = append(conditions, "product_search @@ plainto_tsquery('simple', unaccent($")
		args = append(args, "%"+productName+"%")
	}
	if categoryIDStr != "" {
		categoryID, err := strconv.Atoi(categoryIDStr)
		if err != nil {
			log.Printf("Error converting category_id to int: %v\n", err)
		} else {
			conditions = append(conditions, "category_id = $")
			args = append(args, categoryID)
		}
	}
	if statusStr != "" {
		status, err := strconv.ParseBool(statusStr)
		if err != nil {
			log.Printf("Error converting status to bool: %v\n", err)
		} else {
			conditions = append(conditions, "status = $")
			args = append(args, status)
		}
	}

	// Combine the base query with the conditions
	if len(conditions) > 0 {
		for i, condition := range conditions {
			query += " AND " + condition + strconv.Itoa(i+1)
			// special case for full-text search
			if condition == "product_search @@ plainto_tsquery('simple', unaccent($" {
				query += "))"
			}
		}
	}

	rows, err := h.DB.Query(query, args...)
	if err != nil {
		log.Printf("Error querying products: %v\n", err)
		log.Printf("Query: %s\n", query)
		log.Printf("Args: %+v\n", args)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.ProductName, &p.UnitPrice, &p.Unit, &p.Weight, &p.CategoryID, &p.Status); err != nil {
			log.Printf("Error scanning product: %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		products = append(products, p)
	}

	var categories []models.ProductCategory
	cachedCategories, found := utils.Cache.Get("categories")
	if found {
		categories = cachedCategories.([]models.ProductCategory)
	} else {
		rows, err = h.DB.Query("SELECT id, category_name FROM product_categories")
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
		utils.Cache.Set("categories", categories, cache.DefaultExpiration)
	}

	data := struct {
		Title      string
		Template   string
		Products   []models.Product
		Categories []models.ProductCategory
		Filters    struct {
			ProductID   string
			ProductName string
			CategoryID  string
			Status      string
		}
	}{
		Title:      "Product List",
		Template:   "list-products",
		Products:   products,
		Categories: categories,
	}

	data.Filters.ProductID = productIDStr
	data.Filters.ProductName = productName
	data.Filters.CategoryID = categoryIDStr
	data.Filters.Status = statusStr

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

func (h *ProductHandler) ShowEditProductPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var product models.Product
	err = h.DB.QueryRow("SELECT id, product_name, unit_price, unit, weight, category_id, status from products WHERE id = $1", id).Scan(&product.ID, &product.ProductName, &product.UnitPrice, &product.Unit, &product.Weight, &product.CategoryID, &product.Status)
	if err != nil {
		log.Printf("Error querying product: %v\n", err)
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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
		Product    models.Product
		Categories []models.ProductCategory
	}{
		Title:      "Edit Product",
		Template:   "edit-product",
		Product:    product,
		Categories: categories,
	}

	utils.RenderTemplate(w, "base.html", data)
}

func (h *ProductHandler) EditProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var product models.Product
	product.ProductName = r.FormValue("product_name")
	product.UnitPrice, _ = strconv.ParseFloat(r.FormValue("unit_price"), 64)
	product.Unit = r.FormValue("unit")
	product.Weight, _ = strconv.ParseFloat(r.FormValue("weight"), 64)
	product.CategoryID, _ = strconv.Atoi(r.FormValue("category_id"))
	product.Status = r.FormValue("status") == "true"

	sqlStatement := `UPDATE products SET product_name = $1, unit_price = $2, unit = $3, weight = $4, category_id = $5, status = $6 WHERE id = $7`
	_, err = h.DB.Exec(sqlStatement, product.ProductName, product.UnitPrice, product.Unit, product.Weight, product.CategoryID, product.Status, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
