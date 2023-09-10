package producto

import (
	"context"
	"database/sql"
)

type repository struct {
	db *sql.DB
}

// NewRepositoryMySql creates a new repository.
func NewRepositoryMySql(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

// Create creates a new product.
func (r *repository) Create(ctx context.Context, producto Producto) (Producto, error) {

	statement, err := r.db.Prepare(QueryInsertProduct)

	if err != nil {
		return Producto{}, ErrStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		producto.Name,
		producto.Quantity,
		producto.CodeValue,
		producto.Expiration,
		producto.IsPublished,
		producto.Price,
	)

	if err != nil {
		return Producto{}, ErrExec
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return Producto{}, ErrLastId
	}

	producto.ID = int(lastId)

	return producto, nil
}

// GetAll returns all products.
func (r *repository) GetAll(ctx context.Context) ([]Producto, error) {
	rows, err := r.db.Query(QueryGetAllProducts)
	if err != nil {
		return []Producto{}, err
	}

	defer rows.Close()

	var productos []Producto

	for rows.Next() {
		var producto Producto
		err := rows.Scan(
			&producto.ID,
			&producto.Name,
			&producto.Quantity,
			&producto.CodeValue,
			&producto.Expiration,
			&producto.IsPublished,
			&producto.Price,
		)
		if err != nil {
			return []Producto{}, err
		}

		productos = append(productos, producto)
	}

	if err := rows.Err(); err != nil {
		return []Producto{}, err
	}

	return productos, nil
}

// GetByID returns a product by its ID.
func (r *repository) GetByID(ctx context.Context, id int) (Producto, error) {
	row := r.db.QueryRow(QueryGetProductById, id)

	var producto Producto
	err := row.Scan(
		&producto.ID,
		&producto.Name,
		&producto.Quantity,
		&producto.CodeValue,
		&producto.Expiration,
		&producto.IsPublished,
		&producto.Price,
	)

	if err != nil {
		return Producto{}, err
	}

	return producto, nil
}

// Delete deletes a product.
func (r *repository) Delete(ctx context.Context, id int) error {
	result, err := r.db.Exec(QueryDeleteProduct, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return ErrNotFound
	}

	return nil

}
