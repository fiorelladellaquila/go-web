package producto

import (
	"context"
	"errors"
)

// Errores
var (
	ErrEmptyList = errors.New("the list is empty")
	ErrNotFound  = errors.New("product not found")
	ErrStatement = errors.New("error preparing statement")
	ErrExec      = errors.New("error exect statement")
	ErrLastId    = errors.New("error getting last id")
)

type Repository interface {
	Create(ctx context.Context, producto Producto) (Producto, error)
	GetAll(ctx context.Context) ([]Producto, error)
	GetByID(ctx context.Context, id int) (Producto, error)
	Update(ctx context.Context, producto Producto) (Producto, error)
	Delete(ctx context.Context, id int) error
}
