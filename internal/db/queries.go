package db

const (
	InsertProductSQL     = "INSERT INTO products (name, description, price) VALUES (?, ?, ?)"
	SelectProductsSQL    = "SELECT * FROM products"
	SelectProductByIDSQL = "SELECT * FROM products WHERE id = ?"
	UpdateProductSQL     = "UPDATE products SET name = ?, description = ?, price = ? WHERE id = ?"
	DeleteProductSQL     = "DELETE FROM products WHERE id = ?"

	InsertUserSQL = "INSERT INTO users (username, password, first_name, last_name, email, phone, date_created, role) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	SelectUserSQL = "SELECT id, password, email, role FROM users WHERE username = ?"
	CountUserSQL  = "SELECT COUNT(*) FROM users WHERE username = ?"
)
