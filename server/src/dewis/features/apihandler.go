package features

import (
	//"fmt"
    "net/http"
    "encoding/json"
    "log"
    //"reflect"
)

type RequestJSON struct {
	Request string
	Action 	string
	Data	map[string]string
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {	
	if r.Method == "POST" {
		var req RequestJSON;
		
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Printf("ApiHandler: Something went wrong when decoding the JSON object.\n%v\n", err)
			http.Error(w, "Bad request", http.StatusBadRequest)
		}
		
		// Call different handles depending on type of request
		switch req.Request {
			case "Timeline":
				res := TimelineHandler(req)
				if err := json.NewEncoder(w).Encode(res); err != nil {
					log.Printf("ApiHandler: Something went wrong when encoding the JSON object.\n%v\n", err)
					http.Error(w, "Oops. Something went wrong.", http.StatusInternalServerError)
				}
			case "Login":
				res := LoginHandler(req)
				if err := json.NewEncoder(w).Encode(res); err != nil {
					log.Printf("ApiHandler: Something went wrong when encoding the JSON object.\n%v\n", err)
					http.Error(w, "Oops. Something went wrong.", http.StatusInternalServerError)
				}
		}
	} else {
		http.Error(w, "This method is not allowed.", 403)
	}
}
