package main

import (
	"net/http"
)

type Cluster struct{
	Name string
	Email string
	Lang string
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET"{
	Tmpl.Execute(w, nil)
	return
	}
	if r.Method == "POST"{
		member := Cluster{
		Name: r.FormValue("fullname"),
		Email: r.FormValue("email"),
		Lang: r.FormValue("lang"),
	}
	if err := Suc.Execute(w,  member); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
}
