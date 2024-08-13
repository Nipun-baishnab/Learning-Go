package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("helloHandler called")
	html := ` 
    <!Doctype html>
    <html>
    <head> 
        <title> Arnab </title>
    </head>
    
    <body>
        <h1>Arnab's Planet!</h1>
        <p>Welcome to my brain.</p>
    </body>
    </html>
    `
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
