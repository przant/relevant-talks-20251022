package xml_beverage

import "encoding/xml"

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

// BeveragesResponse wraps multiple beverages for XML serialization
type BeveragesResponse struct {
	XMLName   xml.Name      `xml:"Beverages"`
	Beverages []interface{} `xml:"Beverage"`
}
