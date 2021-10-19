package main

import (
	"fmt"
	"net/http"
)



func main() {
	//health check route 
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Healthy!")
	})

	//Return a simple hello world message. This can be used to make sure the paths work.
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hello world!")
    })
	
	fmt.Printf("Starting server at port 8080\n")
	http.ListenAndServe(":8080", nil)

}
