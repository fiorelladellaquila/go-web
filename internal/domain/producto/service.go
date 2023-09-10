package producto

import (
	"context"
	"errors"
	"log"
)

type service struct {
	repository Repository
}

type Service interface {
	Create(ctx context.Context, requestProduct RequestProducto) (Producto, error)
	GetAll(ctx context.Context) ([]Producto, error)
	GetByID(ctx context.Context, id int) (Producto, error)
	Update(ctx context.Context, requestProduct RequestProducto, id int) (Producto, error)
	Delete(ctx context.Context, id int) error
}

// NewService creates a new product service.
func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

// Create creates a new product.
func (s *service) Create(ctx context.Context, requestProduct RequestProducto) (Producto, error) {
	producto := requestToProducto(requestProduct)
	response, err := s.repository.Create(ctx, producto)
	if err != nil {
		log.Println("error en servicio. Metodo create")
		return Producto{}, errors.New("error en servicio. Metodo create")
	}

	return response, nil
}

// GetAll returns all products.
func (s *service) GetAll(ctx context.Context) ([]Producto, error) {
	productos, err := s.repository.GetAll(ctx)
	if err != nil {
		log.Println("log de error en service de producto", err.Error())
		return []Producto{}, ErrEmptyList
	}

	return productos, nil
}

// GetByID returns a product by its ID.
func (s *service) GetByID(ctx context.Context, id int) (Producto, error) {
	producto, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("log de error en service de producto", err.Error())
		return Producto{}, errors.New("error en servicio. Metodo GetByID")
	}

	return producto, nil
}

// Update updates a product.
func (s *service) Update(ctx context.Context, requestProduct RequestProducto, id int) (Producto, error) {
	producto := requestToProducto(requestProduct)
	producto.ID = id
	response, err := s.repository.Update(ctx, producto)
	if err != nil {
		log.Println("log de error en service de producto", err.Error())
		return Producto{}, errors.New("error en servicio. Metodo Update")
	}

	return response, nil
}

// Delete deletes a product.
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("log de error borrado de producto", err.Error())
		return errors.New("error en servicio. Metodo Delete")
	}

	return nil
}

func requestToProducto(requestProduct RequestProducto) Producto {
	var producto Producto
	producto.Name = requestProduct.Name
	producto.Quantity = requestProduct.Quantity
	producto.CodeValue = requestProduct.CodeValue
	producto.Expiration = requestProduct.Expiration
	producto.IsPublished = requestProduct.IsPublished
	producto.Price = requestProduct.Price

	return producto
}
