package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type Product struct {
    ID        int    `json:"id"`
    SKU       int    `json:"sku"`
    Name      string `json:"name"`
    Category  string `json:"category"`
    Price     int    `json:"price"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
}

func processData(p *url.URL, app *application) (*sql.Rows, error) {
    var rows *sql.Rows
    var err error

    category := p.Query().Get("category")
    priceLessThan := p.Query().Get("priceLessThan")


    // mejorar la logica de aplicacion de filtros y descuentos, actualmente solo se aplica uno a la vez
    if category != "" {
        rows, err = app.DB.Query("SELECT * FROM products WHERE category = $1", category)
        if err != nil {
            log.Println("error en boots",err)
            return nil, err
        }
    } else if priceLessThan != "" {        
        rows, err = app.DB.Query("SELECT * FROM products WHERE price < $1", priceLessThan)
        if err != nil {
            log.Println("error en pricessLess",err)
            return nil, err
        }
    } else {
        rows, err = app.DB.Query("SELECT * FROM products")
        if err != nil {
            return nil, err
        }
    }

    return rows, nil
}


func (app *application) productsHandler(w http.ResponseWriter, r *http.Request) {
    rows, err := processData(r.URL, app)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var products []Product
    for rows.Next() {
        var p Product
        err := rows.Scan(&p.ID, &p.SKU, &p.Name, &p.Category, &p.Price, &p.CreatedAt, &p.UpdatedAt)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        products = append(products, p)

    }

    if err = rows.Err(); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    // mirar de settear correctamente el body de la response agregando al price otro subcampo con el precio con descuento, sin descuento ....
    json.NewEncoder(w).Encode(products)
}

