package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"product_app/models"
)

type CategoryHandler struct {
	DB *sql.DB
}

func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
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

	json.NewEncoder(w).Encode(categories)
}
