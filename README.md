# Beverage API - XML/JSON Encoding Demo

Simple REST API demonstrating Go's `encoding/xml` and `encoding/json` packages with struct composition.

## Features

- Content negotiation (Accept/Content-Type headers)
- Pre-seeded with 25 beverages
- Demonstrates struct composition (`Beverage` → `BeveragePack`)
- Single codebase, dual serialization formats

## Prerequisites

- Go 1.24.1+
- curl (for testing)

## Running the Server
```bash
go mod download
go run main.go
```

Server starts on `http://localhost:8080`

## API Endpoints

### GET /beverages
Retrieve all beverages

**JSON Response:**
```bash
curl -H "Accept: application/json" http://localhost:8080/beverages
```

**XML Response:**
```bash
curl -H "Accept: application/xml" http://localhost:8080/beverages
```

### GET /beverages/{id}
Retrieve single beverage by ID

**JSON Response:**
```bash
# Replace {id} with actual beverage ID from GET /beverages
curl -H "Accept: application/json" http://localhost:8080/beverages/{id}
```

**XML Response:**
```bash
curl -H "Accept: application/xml" http://localhost:8080/beverages/{id}
```

### POST /beverages
Create new beverage

**JSON Request:**
```bash
curl -X POST \
  -H "Content-Type: application/json" \
  -H "Accept: application/json" \
  -d '{"name":"New Beverage","style":"IPA","brewery":"Local Brewery"}' \
  http://localhost:8080/beverages
```

**XML Request:**
```bash
curl -X POST \
  -H "Content-Type: application/xml" \
  -H "Accept: application/xml" \
  -d '<?xml version="1.0" encoding="UTF-8"?><Beverage><name>New Beverage</name><style>IPA</style><brewery>Local Brewery</brewery></Beverage>' \
  http://localhost:8080/beverages
```

## Project Structure
```
.
├── main.go                          # Entry point, server setup, database seeding
├── http/rest/handlers.go            # HTTP handlers, encoding/decoding logic
├── models/
│   ├── json_beverage/beverage.go    # JSON-tagged structs
│   └── xml_beverage/beverage.go     # XML-tagged structs
├── go.mod
└── go.sum
```

## Key Concepts Demonstrated

### Struct Tags
```go
type Beverage struct {
    ID      string `json:"id" xml:"id"`
    Name    string `json:"name" xml:"name"`
    Brewery string `json:"brewery,omitempty" xml:"brewery,omitempty"`
}
```

### Struct Composition (Not Inheritance)
```go
type BeveragePack struct {
    Beverage              // Embedded struct
    Quantity int          `json:"quantity"`
    PackType string       `json:"pack_type"`
}
```

### One-Line Encoding/Decoding
```go
// Decode
json.NewDecoder(r.Body).Decode(&bev)

// Encode
json.NewEncoder(w).Encode(bev)
```

## Branches

- `main`: Manual XML wrapper handling (explicit)
- `feature/xml-wrapper-struct`: Struct-based XML wrapper (elegant)

## Dependencies

- `github.com/google/uuid` - UUID generation for beverage IDs