package controller

import (
	"html/template"
	"log"
	"net/http"
)

type ToolsController struct {
}

func ToolsHandle(w http.ResponseWriter, r *http.Request) {
	log.Println("ToolsHandler")
	tools := &ToolsController{}
	tools.IndexAction(w, r)

}

func (this *ToolsController) IndexAction(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../view/html/tools.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}
