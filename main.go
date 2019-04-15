package main

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"dlser/handler"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("服务启动")
	router := httprouter.New()
	router.POST("/download", handler.RequestHandler)
	router.POST("/getTaskList", handler.GetTaskList)

	err := http.ListenAndServe(":9898", router)
	if err != nil {
		fmt.Println(err)
	}
}