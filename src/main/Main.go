package main

import (
	"controller"
	"net/http"
)

func main() {
	http.Handle("/css/", http.FileServer(http.Dir("../view")))
	http.Handle("/js/", http.FileServer(http.Dir("../view")))
	http.Handle("/images/", http.FileServer(http.Dir("../view")))
	http.Handle("/html/", http.FileServer(http.Dir("../view")))
	http.HandleFunc("/", controller.IndexHandle)
	http.HandleFunc("/tools", controller.ToolsHandle)
	http.ListenAndServe(":8080", nil)
}
