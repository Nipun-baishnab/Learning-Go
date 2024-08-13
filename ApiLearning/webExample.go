package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	html := `
    <!DOCTYPE html>
    <html>
    <head>
        <title>Hello, World</title>
    </head>
    <body>
        <h1>Hello, World!</h1>
        <p>Welcome to my webpage.</p> <!-- This line should be present -->
    </body>
    </html>`
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
