package main

import (
	"os"
    "fmt"
    "net/http"
    //"encoding/json"
    "dewis/features"
)

func apiHandler(w http.ResponseWriter, r *http.Request) {	
	if r.Method == "POST" {
		r.ParseForm()
		//request := (r.Form["request"])

		switch request {
			case "timeline":
				res := timelineHandler(&r.Form)
		}

		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.Println(err)
			http.Error(w, "Oops. Something went wrong.", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "This method is not allowed.", 403)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Not enough arguments. Please read documentation.")
		os.Exit(1)
	}
 
	if len(args) == 1 && args[0] == "run" {
		fmt.Println("Starting server")
		http.HandleFunc("/api", apiHandler)
		http.Handle("/", http.FileServer(http.Dir("/var/develop/dewis/client")))
		fmt.Println("Server running...")
    	http.ListenAndServe(":8080", nil)
	} else {
		fmt.Println("Wrong arguments. Please read documentation.")
	}
	fmt.Println("Server closed")
}
