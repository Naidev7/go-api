package main

import (
	"net/http"
)

func (app *application) routes() http.Handler{

	mux := http.NewServeMux()

	mux.HandleFunc("/products/", app.productsHandler)


	return mux

}
