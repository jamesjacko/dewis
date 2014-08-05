package main

import (
	"os"
    "fmt"
    "net/http"
    "dewis/features"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Not enough arguments. Please read documentation.")
		os.Exit(1)
	}
 
	if len(args) == 1 && args[0] == "run" {
		fmt.Println("Starting server test")
		http.HandleFunc("/api", features.ApiHandler)
		http.Handle("/", http.FileServer(http.Dir("/var/develop/dewis/static/")))
		fmt.Println("Server running...")
    	http.ListenAndServe(":8080", nil)
	} else {
		fmt.Println("Wrong arguments. Please read documentation.")
	}
	fmt.Println("Server closed")
}
