package repository

import (
	"database/sql"
	"go-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) *ProductRepository {
	return &ProductRepository{connection: connection}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT * FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		panic(err)
	}

	var products []model.Product

	for rows.Next() {
		product := model.Product{}
		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			panic(err)
		}

		products = append(products, product)
	}
	rows.Close()

	return products, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {

	var id int
	query, err := pr.connection.Prepare("INSERT INTO product" +
		" (product_name, price)" +
		"VALUES ($1, $2) RETURNING id")
	if err != nil {
		print(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		print(err)
		return 0, err
	}

	return id, nil
}
