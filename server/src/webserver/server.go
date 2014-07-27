package main

import (
	"os"
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to AstralDynamics!")
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Not enough arguments. Please read documentation.")
		os.Exit(1)
	}
 
	if len(args) == 1 && args[0] == "run" {
		fmt.Println("Starting server")
    	http.HandleFunc("/", handler)
		fmt.Println("Server running...")
    	http.ListenAndServe(":8080", nil)
	}

	fmt.Println("Wrong arguments. Please read documentation.")
}
