package features

import (
	//"fmt"
    "net/http"
    "encoding/json"
    "log"
)

type RequestJSON struct {
	Request string
	Action 	string
	Data	map[string]string
}

// Constants for database names
const DATABASE_NAME			= "dewis"
const USERS_COLLECTION		= "Users"
const RECORDS_COLLECTION	= "timelineRecords"

func save(w http.ResponseWriter, resp interface{}) {
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("ApiHandler: Something went wrong when encoding the JSON object.\n%v\n", err)
		http.Error(w, "Oops. Something went wrong.", http.StatusInternalServerError)
	}
	//fmt.Println(encoder)
}

/* TODO
 * Figure out why it is so slow to return the login anwser
 * Modularize some "Login" code into functions
 * Add Login and Logout actions in Auth
*/
func ApiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Parsing JSON from request
		var req RequestJSON;
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Printf("ApiHandler: Something went wrong when decoding the JSON object.\n%v\n", err)
			http.Error(w, "Bad request", http.StatusBadRequest)
		}
		
		// Call different handlers depending on type of request
		switch req.Request {
			case "Timeline":
				// Check the session cookie to see if user is authenticated and session is valid
				respCookie := CheckCookie(r)
				if respCookie.Status == false {
					save(w, &respCookie)
				}

				res := timelineHandler(req)
				save(w, &res)
			case "Auth":
				res := loginHandler(req, w, r)
				save(w, &res)
			case "User":
				// Check the session cookie to see if user is authenticated and session is valid
				respCookie := CheckCookie(r)
				if respCookie.Status == false {
					save(w, &respCookie)
				}

				res := userHandler(req)
				save(w, &res)
		}
	} else {
		http.Error(w, "This method is not allowed.", 403)
	}
}