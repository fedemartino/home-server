package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
func homeHandler(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("html/index.html")
	t.Execute(w, "")
}

func main() {
	var port = ":9090"

	// Serve files from the "static" directory
	jsFileServer := http.FileServer(http.Dir("./html/js"))
	imgFileServer := http.FileServer(http.Dir("./html/img"))
	http.Handle("/js/", http.StripPrefix("/js/", jsFileServer))
	http.Handle("/img/", http.StripPrefix("/img/", imgFileServer))

	//handle dynamic urls
	http.HandleFunc("/", handler)
	http.HandleFunc("/home/", homeHandler)

	fmt.Println("Listening on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
