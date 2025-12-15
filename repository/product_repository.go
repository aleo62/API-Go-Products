package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
	"strconv"
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

func (pr *ProductRepository) DeleteProduct(id int) (bool, error) {
	query, err := pr.connection.Prepare("DELETE FROM product WHERE id = $1")
	if err != nil {
		print(err)
		return false, err
	}

	result, err := query.Exec(id)
	if err != nil {
		print(err)
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		print(err)
		return false, err
	}

	return rowsAffected > 0, nil
}

func (pr *ProductRepository) UpdateProduct(id int, fields map[string]interface{}) (bool, error) {
	var queryString string
	var index int
	args := []interface{}{}
	
	print(len(fields))

	allowed := []string{"product_name", "price"}

	for _, k := range allowed {
		v, ok := fields[k]
		if !ok {
			continue
		}
		if index > 0 {
			queryString += ", "
		}

		queryString += fmt.Sprintf("%s = $%d", k, index+1)
		index++
		args = append(args, v)
	}

	query, err := pr.connection.Prepare("UPDATE product SET " + queryString + " WHERE id = $" + strconv.Itoa(index+1))
	print(index)
	if err != nil {
		print(err)
		return false, err
	}

	args = append(args, id)
	result, err := query.Exec(args...)
	if err != nil {
		print(err)
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		print(err)
		return false, err
	}

	return rowsAffected > 0, nil
}
