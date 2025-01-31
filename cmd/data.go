package main

type Product struct {
	ID        int    `json:"id"`
	SKU       int    `json:"sku"`
	Name      string `json:"name"`
	Category  string `json:"category"`
	Price     int    `json:"price"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
