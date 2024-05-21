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
	categories, err := GetCategories(h.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(categories)
}

// GetCategories fetches categories from the cache or the database
func GetCategories(db *sql.DB) ([]models.ProductCategory, error) {
	// Check if categories are in the cache
	if cachedCategories, found := utils.Cache.Get("categories"); found {
		return cachedCategories.([]models.ProductCategory), nil
	}

	// Fetch categories from the database
	rows, err := db.Query("SELECT id, category_name FROM product_categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.ProductCategory
	for rows.Next() {
		var c models.ProductCategory
		if err := rows.Scan(&c.ID, &c.CategoryName); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	// Cache the categories
	utils.Cache.Set("categories", categories, cache.DefaultExpiration)
	return categories, nil
}
