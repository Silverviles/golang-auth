package services

import (
	"go-app/internal/db"
	"go-app/internal/models"
)

func AddProduct(product *models.Product) (*models.Product, error) {
	DB := db.GetDatabaseConnection()

	transaction, err := DB.Begin()
	if err != nil {
		return nil, err
	}

	insertProductSQL := "INSERT INTO products (name, description, price) VALUES (?, ?, ?)"
	result, err := DB.Exec(insertProductSQL, product.Name, product.Description, product.Price)
	if err != nil {
		return nil, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	err = transaction.Commit()
	if err != nil {
		return nil, err
	}

	product.ID = int(lastInsertedID)
	return product, nil
}

func GetProducts() ([]models.Product, error) {
	DB := db.GetDatabaseConnection()

	selectProductsSQL := "SELECT * FROM products"
	rows, err := DB.Query(selectProductsSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Close the rows if the execution fails in the middle

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func GetProductByID(productID int) (*models.Product, error) {
	DB := db.GetDatabaseConnection()

	selectProductSQL := "SELECT * FROM products WHERE id = ?"
	row := DB.QueryRow(selectProductSQL, productID)

	var product models.Product
	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func UpdateProduct(product *models.Product) (*models.Product, error) {
	DB := db.GetDatabaseConnection()

	transaction, err := DB.Begin()
	if err != nil {
		return nil, err
	}

	updateProductSQL := "UPDATE products SET name = ?, description = ?, price = ? WHERE id = ?"
	_, err = DB.Exec(updateProductSQL, product.Name, product.Description, product.Price, product.ID)
	if err != nil {
		return nil, err
	}

	err = transaction.Commit()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func DeleteProduct(productID int) error {
	DB := db.GetDatabaseConnection()

	transaction, err := DB.Begin()
	if err != nil {
		return err
	}

	deleteProductSQL := "DELETE FROM products WHERE id = ?"
	_, err = DB.Exec(deleteProductSQL, productID)
	if err != nil {
		return err
	}

	err = transaction.Commit()
	if err != nil {
		return err
	}

	return nil
}
