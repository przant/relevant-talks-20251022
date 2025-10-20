package xml_beverage

// Beverage represents a single Beverage with XML tags
type Beverage struct {
	ID      string `xml:"id"`
	Name    string `xml:"name"`
	Style   string `xml:"style"`
	Brewery string `xml:"brewery,omitempty"`
}

// BeveragePack represents a pack of Beverages (composition example)
type BeveragePack struct {
	Beverage         // Embedded Beverage struct (composition!)
	Quantity int     `xml:"quantity"`
	PackType string  `xml:"pack_type"`
	Price    float64 `xml:"price,omitempty"`
}
