package json_beverage

// Beverage represents a single beverage with JSON tags
type Beverage struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Style   string `json:"style"`
	Brewery string `json:"brewery,omitempty"`
}

// BeveragePack represents a pack of beverages (composition example)
type BeveragePack struct {
	Beverage         // Embedded Beverage struct (composition!)
	Quantity int     `json:"quantity"`
	PackType string  `json:"pack_type"`
	Price    float64 `json:"price,omitempty"`
}
