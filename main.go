package main

import (
	"fmt"
	"log"
	"net/http"
)

// handler for form submission
func FormHandler(w http.ResponseWriter, r *http.Request) {
	// Parse form and handle error
	if err := r.ParseForm(); err != nil{
		http.Error(w, fmt.Sprintf("Parseform() error : %v", err), http.StatusInternalServerError)
		// fmt.Sprintf() returns a formatted string not prints
		return
	}

	// Retrieve Form Values
	name := r.FormValue("name")
	address := r.FormValue("address")

	// Resposne with submitted data
	fmt.Fprintf(w, "Thank You for submitting the form :)\n")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)

	//fmt.Fprintf is often used to send a formatted response to the client via an http.ResponseWriter.(w) 
	//it returns the number of bytes written
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// check for correct url path
	if r.URL.Path != "/hello"{
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	// Ensure that only GET requests are allowed

	if r.Method != "GET"{
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Hello Response to user
	fmt.Fprint(w, "Hello 👋")
}

func main() {
	// Serve static files from the "static" directory
	fileserver := http.FileServer(http.Dir("./static")) 
	// What it does: This serves static files from the static directory. Any file in this directory can be accessed via the root route (/). For example, index.html and form.html can be accessed by visiting /index.html and /form.html in the browser.

	// Set up routes

	http.Handle("/", fileserver)  // Serve static files at root
	http.HandleFunc("/hello", helloHandler) // hanlde /hello request
	http.HandleFunc("/form", FormHandler) // handle form request

	// Start the server and listen for incoming requests
	fmt.Println("Starting server on port 8080...")
	log.Fatalf("Server failed: %v", http.ListenAndServe(":8000", nil))

	// nil: Tells the server to use the default HTTP request multiplexer (http.DefaultServeMux) to handle incoming requests.
	// The default mux is an instance of http.ServeMux that automatically handles the default routes, such as / and others, based on the registered handlers.
}