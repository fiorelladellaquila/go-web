package producto

import "time"

// Producto describes a product.
type Producto struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Quantity    int       `json:"quantity"`
	CodeValue   string    `json:"code_value"`
	Expiration  time.Time `json:"expiration"`
	IsPublished bool      `json:"is_published"`
	Price       float64   `json:"price"`
}

// RequestProducto describes the data needed to create a new product.
type RequestProducto struct {
	Name        string    `json:"name"`
	Quantity    int       `json:"quantity"`
	CodeValue   string    `json:"code_value"`
	Expiration  time.Time `json:"expiration"`
	IsPublished bool      `json:"is_published"`
	Price       float64   `json:"price"`
}
