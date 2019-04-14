package main

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"dlser/handler"
)

func main() {
	router := httprouter.New()
	
	router.POST("/download", handler.RequestHandler)

	err := http.ListenAndServe(":9898", router)
	if err != nil {
		fmt.Println(err)
	}
}