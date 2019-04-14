package main

import (
	// "fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"dlser/handler"
)

func main() {
	router := httprouter.New()

	router.POST("/download", handler.RequestHandler)

	http.ListenAndServe(":7080", router)
}