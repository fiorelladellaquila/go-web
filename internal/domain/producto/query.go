package producto

var (
	QueryInsertProduct = `INSERT INTO my_db.products(name,quantity,code_value,expiration,is_published, price)
	VALUES(?,?,?,?,?,?)`
	QueryGetAllProducts = `SELECT id, name, quantity, code_value, expiration, is_published, price 
	FROM my_db.products`
	QueryDeleteProduct  = `DELETE FROM my_db.products WHERE id = ?`
	QueryGetProductById = `SELECT id, name, quantity, code_value, expiration, is_published, price
	FROM my_db.products WHERE id = ?`
	QueryUpdateProduct = `UPDATE my_db.products SET name = ?, quantity = ?, code_value = ?, expiration = ?, is_published = ?, price = ?
	WHERE id = ?`
)
