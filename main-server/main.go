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
func homeHandler(w http.ResponseWriter) {

	fmt.Println("Serving home html")
	if w == nil {
		fmt.Println("Responser Writer is nil")
	} else {
		fmt.Println("Response writer is NOT nil")
	}
	t, err := template.ParseFiles("html/index.html")
	if err != nil {
		fmt.Println("Error getting home html")
		fmt.Println(err.Error())
	} else {
		fmt.Println("no error getting template")
	}

	t.Execute(w, "")
}

func genericFileHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Generic file Handler executed")
	fmt.Println(r.URL.Path)
	file := r.URL.Path[1:]
	fmt.Println("requested file", file)
	if file == "" {
		fmt.Println("requested path is root /")
		homeHandler(w)
	} else {
		t, _ := template.ParseFiles("html/" + file + ".html")
		t.Execute(w, "")
	}
}

func main() {
	var port = ":9090"

	// Serve files from the "static" directory
	jsFileServer := http.FileServer(http.Dir("./html/js"))
	cssFileServer := http.FileServer(http.Dir("./html/css"))
	imgFileServer := http.FileServer(http.Dir("./html/img"))
	http.Handle("/js/", http.StripPrefix("/js/", jsFileServer))
	http.Handle("/img/", http.StripPrefix("/img/", imgFileServer))
	http.Handle("/css/", http.StripPrefix("/css/", cssFileServer))

	//handle dynamic urls
	http.HandleFunc("/", genericFileHandler)
	http.HandleFunc("/message", handler)

	fmt.Println("Listening on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
