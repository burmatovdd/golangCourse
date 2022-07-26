package server

import (
	"fmt"
	"homework/models"
	"homework/userInfo"
	"html/template"
	"net/http"
)

type Service struct{}

type Server interface {
	Handler()
}

type ViewData struct {
	models.User
}

var testTemplate *template.Template

func (service *Service) Handler() {

	var err error
	testTemplate, err = template.ParseFiles("templates/index.html")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", getInfo)
	fmt.Println("server is starting...")
	http.ListenAndServe(":1234", nil)
}

func getInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	method := userInfo.Service{}

	user := method.AddUser()

	vd := ViewData{user}

	err := testTemplate.Execute(w, vd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
