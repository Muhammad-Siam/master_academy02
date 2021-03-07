package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/Features", Features)
	http.HandleFunc("/Docs", Docs)

	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe("localhost:8888", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)
	//fmt.Fprintf(w, "Welcome to siam")
}
func Features(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("wpage/features.gohtml")

	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)
	//fmt.Fprintf(w, "Welcome to siam")
}

func Docs(w http.ResponseWriter, r *http.Request) {
	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("wpage/docs.gohtml")

	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)
	//fmt.Fprintf(w, "Welcome to siam")
}
