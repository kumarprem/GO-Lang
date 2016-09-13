package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		fmt.Fprintf(w, "Webservice is called")
		fmt.Println("Webservice is called");
		
		
    })

    log.Fatal(http.ListenAndServe(":8082", nil))

}