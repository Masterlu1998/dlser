package main

import (
	"dlser/handler"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	fmt.Println("服务启动")
	router := httprouter.New()
	router.POST("/download", handler.RequestHandler)
	router.POST("/getTaskList", handler.GetTaskListHandler)
	router.POST("/getFile", handler.GetFileHandler)
	router.POST("/deleteFile", handler.DeleteFileHandler)
	router.ServeFiles("/download/file/*filepath", http.Dir("./file"))

	err := http.ListenAndServe(":9898", router)
	if err != nil {
		fmt.Println(err)
	}
}
