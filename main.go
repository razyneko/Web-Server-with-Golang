package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler (w http.ResponseWriter,r *http.Request) {
	// error handling
	if err := r.ParseForm(); err != nil{
		fmt.Fprintf(w,"Parseform() err : %v", err)
		return
	}
	fmt.Fprint(w,"POST request successful")
	// accessing name and address fields from the form
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler (w http.ResponseWriter,r *http.Request) {
	// checking correct path
	if r.URL.Path != "/hello"{
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	// Restricting POST request on "/hello"
	if r.Method != "GET"{
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w,"hello !")
}

func main() {
	// telling golang to check "/static" directory
	fileServer := http.FileServer(http.Dir("./static"))
	// handling root route
	http.Handle("/", fileServer)
	// handling form route
	http.HandleFunc("/form", formHandler)
	// handling hello route
	http.HandleFunc("/hello", helloHandler)
	// using Printf to use \n
	fmt.Printf("Starting server at port 8080\n")
	// handling errors
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}