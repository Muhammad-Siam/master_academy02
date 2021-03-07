package main

import (
	"fmt"
	"net/http"
)

/*type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}*/

func main() {
	//handler func(ResponseWriter, *Request)
	//var x string
	//var handler func(ResponseWriter, *Request)

	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("img")))) //curent dir
	//http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/address", address)
	http.ListenAndServe("localhost:8888", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my first golang web page")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my about page")
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my first golang web page/contact")
}

func address(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html;charset= UTF-8")
	fmt.Fprintf(w, "<img src=\"siam.png\" />")
}
