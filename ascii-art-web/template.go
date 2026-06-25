package main

import "html/template"

var Tmpl = template.Must(
	template.ParseFiles("form/form.html"),
)
var Suc = template.Must(
	template.ParseFiles("form/success.html"),
)