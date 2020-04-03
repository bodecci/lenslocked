package main

import (
	"fmt"
	"github.com/bodecci/use_golang/lenslocked/views"
	_ "html/template"
	"log"

	"github.com/gorilla/mux"
	"net/http"
)

var (
	homeView    *views.View
	contactView *views.View
	signupView *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
}

func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(signupView.Render(w, nil))
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `These are the FAQs for this site
					<ul>
					<li>something here</li>
					<li>something2 here</li></ul>`)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>We couldn't find the page you're looking for</h1>")
}

func main() {

	homeView = views.NewView("bootstrap","views/home.gohtml")
	contactView = views.NewView("bootstrap","views/contact.gohtml")
	signupView = views.NewView("bootstrap", "views/signup.gohtml")

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	r.HandleFunc("/upload", signup)
	fmt.Println("Server starting 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
