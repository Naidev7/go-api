package main

import (
	"database/sql"
)

type Store interface {
	GetProducts(category string, price, limit, offset int) ([]Product, error)
}

type postgresStore struct {
	DB *sql.DB
}

func NewPostgresStore(db *sql.DB) *postgresStore {
	return &postgresStore{
		DB: db,
	}
}

func (s *postgresStore) GetProducts(category string, price, limit, offset int) ([]Product, error) {
	var q string
	var rows *sql.Rows
	var err error
	if category != "" {
		q = "SELECT * FROM  products WHERE category = $1  OFFSET $2 LIMIT $3"
		rows, err = s.DB.Query(q, category, offset, limit)
	}
	if price > 0 {
		q = "SELECT * FROM products WHERE price < $1 OFFSET $2 LIMIT $3"
		rows, err = s.DB.Query(q, price, offset, limit)
	}
	if q == "" {
		q = "SELECT * FROM products OFFSET $1 LIMIT $1"
		rows, err = s.DB.Query(q, offset, limit)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.SKU, &p.Name, &p.Category, &p.Price, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
