package rest

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/przant/relevant-talks-20251022/models/json_beverage"
	"github.com/przant/relevant-talks-20251022/models/xml_beverage"
)

// In-memory storage
var beverageDB = make(map[string]interface{})

// GetBeverageDB returns the database map (for seeding)
func GetBeverageDB() map[string]interface{} {
	return beverageDB
}

// GetAllBeverages handles GET /beverages
func GetAllBeverages(w http.ResponseWriter, r *http.Request) {
	acceptHeader := r.Header.Get("Accept")

	// Collect all beverages
	var beverages []interface{}
	for _, bev := range beverageDB {
		beverages = append(beverages, bev)
	}

	// Respond based on Accept header
	if strings.Contains(acceptHeader, "application/xml") {
		w.Header().Set("Content-Type", "application/xml")
		w.Write([]byte(xml.Header))
		w.Write([]byte(`<Beverages>`))
		xml.NewEncoder(w).Encode(beverages)
		w.Write([]byte(`</Beverages>`))
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(beverages)
	}
}

// GetBeverageByID handles GET /beverages/{id}
func GetBeverageByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/beverages/")

	bev, exists := beverageDB[id]
	if !exists {
		http.Error(w, "Beverage not found", http.StatusNotFound)
		return
	}

	acceptHeader := r.Header.Get("Accept")

	if strings.Contains(acceptHeader, "application/xml") {
		w.Header().Set("Content-Type", "application/xml")
		w.Write([]byte(xml.Header))
		xml.NewEncoder(w).Encode(bev)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(bev)
	}
}

// CreateBeverage handles POST /beverages
func CreateBeverage(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	id := uuid.New().String()

	if strings.Contains(contentType, "application/xml") {
		var bev xml_beverage.Beverage
		if err := xml.NewDecoder(r.Body).Decode(&bev); err != nil {
			http.Error(w, fmt.Sprintf("Invalid XML: %v", err), http.StatusBadRequest)
			return
		}
		bev.ID = id
		beverageDB[id] = bev

		w.Header().Set("Content-Type", "application/xml")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(xml.Header))
		xml.NewEncoder(w).Encode(bev)
	} else {
		var bev json_beverage.Beverage
		if err := json.NewDecoder(r.Body).Decode(&bev); err != nil {
			http.Error(w, fmt.Sprintf("Invalid JSON: %v", err), http.StatusBadRequest)
			return
		}
		bev.ID = id
		beverageDB[id] = bev

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(bev)
	}
}
