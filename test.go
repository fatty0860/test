package main

import "net/http"
import "html/template"

type Person struct {
	Name string
	Age  string
}

func indexHdl(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	var obj Person

	if r.Method == "GET" {
		obj.Name = r.FormValue("Name")
		t, _ := template.ParseFiles("web/index.html")
		t.Execute(w, obj)

	} else if r.Method == "POST" {
		obj.Name = r.FormValue("Name")
		obj.Age = r.FormValue("Age")
		t, _ := template.ParseFiles("web/index.html")
		t.Execute(w, obj)
	}
}

func main() {
	http.Handle("/css/", http.FileServer(http.Dir("web")))
	http.Handle("/script/", http.FileServer(http.Dir("web")))
	http.HandleFunc("/", indexHdl)

	http.ListenAndServe(":8888", nil)
}
