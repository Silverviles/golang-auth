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

	result, err := DB.Exec(db.InsertProductSQL, product.Name, product.Description, product.Price)
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

	rows, err := DB.Query(db.SelectProductsSQL)
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

	row := DB.QueryRow(db.SelectProductByIDSQL, productID)

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

	_, err = DB.Exec(db.UpdateProductSQL, product.Name, product.Description, product.Price, product.ID)
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

	_, err = DB.Exec(db.DeleteProductSQL, productID)
	if err != nil {
		return err
	}

	err = transaction.Commit()
	if err != nil {
		return err
	}

	return nil
}
