package main

import (
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// fmt.Fprint(w, r.URL.Path) // shows the address path in the DOM

	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my  Super Awesome site!</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, `To get in touch, please send an email to <a
		href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>`)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `<h1>We could not find the page you're looking for</h1> <p>Please email us if you keep being sent to an invalid page</p>`)
	}
}

// func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	fmt.Fprint(w, "Welcome!\n")
//}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", handlerFunc)
	fmt.Println("Server starting")
	http.ListenAndServe(":3000", r)
}
