package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"product_app/models"
	"product_app/utils"

	"github.com/patrickmn/go-cache"
)

type CategoryHandler struct {
	DB *sql.DB
}

func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	cachedCategories, found := utils.Cache.Get("categories")
	if found {
		json.NewEncoder(w).Encode(cachedCategories)
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

	utils.Cache.Set("categories", categories, cache.DefaultExpiration)
	json.NewEncoder(w).Encode(categories)
}
