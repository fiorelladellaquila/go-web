package producto

import (
	"context"
	"time"
)

// Base de datos en memoria.
var (
	productos = []Producto{
		{
			ID:          1,
			Name:        "Banana",
			CodeValue:   "AABBCCC",
			Quantity:    10,
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       10.0,
		},
		{
			ID:          2,
			Name:        "Manzana",
			CodeValue:   "AABBDDD",
			Quantity:    5,
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       5.0,
		},
		{
			ID:          3,
			Name:        "Pera",
			CodeValue:   "AAZZZCCC",
			Quantity:    8,
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       8.0,
		},
	}
)

type repositoryMemory struct {
	db []Producto
}

func NewRepositoryMemory() Repository {
	return &repositoryMemory{
		db: productos,
	}
}

// Create creates a new product.
func (r *repositoryMemory) Create(ctx context.Context, producto Producto) (Producto, error) {
	r.db = append(r.db, producto)

	return producto, nil
}

// GetAll returns all products.
func (r *repositoryMemory) GetAll(ctx context.Context) ([]Producto, error) {
	if len(r.db) < 1 {
		return []Producto{}, ErrEmptyList
	}

	return r.db, nil
}

// GetByID returns a product by its ID.
func (r *repositoryMemory) GetByID(ctx context.Context, id int) (Producto, error) {
	for _, producto := range r.db {
		if producto.ID == id {
			return producto, nil
		}
	}

	return Producto{}, ErrNotFound
}

// Delete deletes a product.
func (r *repositoryMemory) Delete(ctx context.Context, id int) error {

	for key, producto := range r.db {
		if producto.ID == id {
			r.db = append(r.db[:key], r.db[key+1:]...)
			return nil
		}

	}

	return ErrNotFound
}
