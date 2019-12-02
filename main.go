package main

import (
	"fmt"

	"github.com/gorilla/mux"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my  Super Awesome site!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `To get in touch, please send an email to <a
		href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>`)

}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `These are the FAQs for this site
					<ul>
					<li>something here</li>
					<li>something2 here</li></ul>`)
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	fmt.Println("Server starting")
	http.ListenAndServe(":3000", r)
}
