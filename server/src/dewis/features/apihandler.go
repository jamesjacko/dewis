package features

import (
	"fmt"
    "net/http"
    "encoding/json"
    "log"
    "github.com/gorilla/sessions"
    //"reflect"
)

type RequestJSON struct {
	Request string
	Action 	string
	Data	map[string]string
}

// Constants for database names
const databaseName 			= "dewis"
const usersCol				= "Users"
const timelineRecordsCol	= "timelineRecords"

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	store := sessions.NewCookieStore()
	session, _ := store.Get(r, "thiago")
	session.Values["lol"] = "hehe"
	session.ID = "123456789"

	fmt.Println(session)

	session1, _ := store.Get(r, "thiago")
	fmt.Println(session1)

	session.Values["lol"] = "hehe1"

	fmt.Println(session)
	fmt.Println(session1)

	if r.Method == "POST" {
		var req RequestJSON;
		
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Printf("ApiHandler: Something went wrong when decoding the JSON object.\n%v\n", err)
			http.Error(w, "Bad request", http.StatusBadRequest)
		}
		
		// Call different handles depending on type of request
		switch req.Request {
			case "Timeline":
				res := timelineHandler(req)
				if err := json.NewEncoder(w).Encode(res); err != nil {
					log.Printf("ApiHandler: Something went wrong when encoding the JSON object.\n%v\n", err)
					http.Error(w, "Oops. Something went wrong.", http.StatusInternalServerError)
				}
			case "Login":
				res := loginHandler(req)
				if err := json.NewEncoder(w).Encode(res); err != nil {
					log.Printf("ApiHandler: Something went wrong when encoding the JSON object.\n%v\n", err)
					http.Error(w, "Oops. Something went wrong.", http.StatusInternalServerError)
				}
			case "User":
				res := userHandler(req)
				if err := json.NewEncoder(w).Encode(res); err != nil {
					log.Printf("ApiHandler: Something went wrong when encoding the JSON object.\n%v\n", err)
					http.Error(w, "Oops. Something went wrong.", http.StatusInternalServerError)
				}
		}
	} else {
		http.Error(w, "This method is not allowed.", 403)
	}
}
