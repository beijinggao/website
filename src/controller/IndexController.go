package controller

import (
	"html/template"
	"log"
	"net/http"
)

type IndexController struct {
}

func IndexHandle(w http.ResponseWriter, r *http.Request) {
	log.Println("indexHandler")
	/*pathInfo := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(pathInfo, "/")
	var action = ""
	if len(parts) > 1 {
		action = strings.Title(parts[1]) + "Action"
	}

	index := &IndexController{}
	controller := reflect.ValueOf(index)
	method := controller.MethodByName(action)
	if !method.IsValid() {
		method = controller.MethodByName(strings.Title("index") + "Action")
	}
	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)
	method.Call([]reflect.Value{responseValue, requestValue})*/
	index := &IndexController{}
	index.IndexAction(w, r)

}

func (this *IndexController) IndexAction(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../view/index.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}
