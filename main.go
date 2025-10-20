package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/przant/relevant-talks-20251022/http/rest"
	"github.com/przant/relevant-talks-20251022/models/json_beverage"
	"github.com/przant/relevant-talks-20251022/models/xml_beverage"
)

func main() {
	// Pre-seed the database with 25 beverages
	seedDatabase()

	// Route handler
	http.HandleFunc("/beverages", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			rest.GetAllBeverages(w, r)
		} else if r.Method == http.MethodPost {
			rest.CreateBeverage(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/beverages/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			rest.GetBeverageByID(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func seedDatabase() {
	beverages := []struct {
		name    string
		style   string
		brewery string
	}{
		{"Pale Ale", "IPA", "Stone Brewing"},
		{"Lager Light", "Lager", "Corona"},
		{"Amber Ale", "Amber", "New Belgium"},
		{"Stout Porter", "Stout", "Guinness"},
		{"Wheat Beer", "Wheat", "Blue Moon"},
		{"Pilsner Classic", "Pilsner", "Pilsner Urquell"},
		{"Red Ale", "Red Ale", "Smithwick's"},
		{"Brown Ale", "Brown", "Newcastle"},
		{"Belgian White", "Witbier", "Hoegaarden"},
		{"Saison Farmhouse", "Saison", "Saison Dupont"},
		{"Kolsch Light", "Kolsch", "Reissdorf"},
		{"Bock Strong", "Bock", "Shiner"},
		{"Hefeweizen", "Wheat", "Weihenstephaner"},
		{"Double IPA", "DIPA", "Pliny the Elder"},
		{"Sour Ale", "Sour", "Rodenbach"},
		{"Barleywine", "Barleywine", "Sierra Nevada"},
		{"Porter Dark", "Porter", "Anchor"},
		{"Scotch Ale", "Scotch Ale", "Founders"},
		{"Tripel Belgian", "Tripel", "Westmalle"},
		{"Quadrupel", "Quad", "Rochefort"},
		{"Marzen Oktoberfest", "Marzen", "Paulaner"},
		{"Schwarzbier", "Dark Lager", "Kostritzer"},
		{"Cream Ale", "Cream Ale", "Genesee"},
		{"Golden Ale", "Golden", "Victory"},
		{"Session IPA", "Session", "Founders"},
	}

	for _, bev := range beverages {
		id := uuid.New().String()

		// Alternate between JSON and XML models for variety
		if len(rest.GetBeverageDB())%2 == 0 {
			jsonBev := json_beverage.Beverage{
				ID:      id,
				Name:    bev.name,
				Style:   bev.style,
				Brewery: bev.brewery,
			}
			rest.GetBeverageDB()[id] = jsonBev
		} else {
			xmlBev := xml_beverage.Beverage{
				ID:      id,
				Name:    bev.name,
				Style:   bev.style,
				Brewery: bev.brewery,
			}
			rest.GetBeverageDB()[id] = xmlBev
		}
	}

	fmt.Printf("Database seeded with %d beverages\n", len(rest.GetBeverageDB()))
}
