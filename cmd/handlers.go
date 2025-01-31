package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (app *application) strToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}
func (app *application) productsHandler(w http.ResponseWriter, r *http.Request) {

	req := struct {
		Category string `json:"category"`
		Price    int    `json:"price"`
		Limit    int    `json:"limit"`
		Offset   int    `json:"offset"`
	}{
		Category: r.URL.Query().Get("category"),
		Price:    app.strToInt(r.URL.Query().Get("price")),
		Limit:    app.strToInt(r.URL.Query().Get("limit")),
	}
	if req.Limit == 0 {
		req.Limit = 10
	}
	req.Offset = req.Offset * req.Limit

	products, err := app.Store.GetProducts(req.Category, req.Price, req.Limit, req.Offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// mirar de settear correctamente el body de la response agregando al price otro subcampo con el precio con descuento, sin descuento ....
	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
