package features

import (
	//"fmt"
	"time"
    "net/http"
    "encoding/json"
    "log"
    "github.com/gorilla/securecookie"
    //"reflect"
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
}

/* TODO
 * Figure out why it is so slow to return the login anwser
 * Modularize some "Login" code into functions
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
				respCookie := checkCookie(r)
				if respCookie.Status == false {
					save(w, &respCookie)
				}

				res := timelineHandler(req)
				save(w, &res)
			case "Auth":
				res := loginHandler(req)

				if res.Status == true {
					var session *Session

					//If there is no session for that user, create one with a unique SessionID, add it to a cookie and send the cookie back 
					sessionID, err := Store.GetSessionID(req.Data["Username"]); 
					if err != nil {
						//Crating session with unique sessionID and storing it
						sessionID = Store.GenerateSessionID(req.Data["Username"])

						//Check if that sessionID already exists (shit happens) and create another if so
						_, err := Store.GetUsername(sessionID);
						for err == nil {
							sessionID = Store.GenerateSessionID(req.Data["Username"])
							_, err = Store.GetUsername(sessionID);
						}

						session = &Session{req.Data["Username"], sessionID, time.Now().Unix()}
						Store.AddSession(*session)
					} else {
						session = Store.GetSession(sessionID)
						session.lastUsed = time.Now().Unix()
					}

					//Store.Print()

					//Creating encoding value
					var s = securecookie.New([]byte("dewis-hashkey-cookie"), []byte("encryption-key-dewis-hash78aw971"))
					encoded, err := s.Encode("session", session);

					//Creating and setting cookie
					if err == nil {
					    cookie := &http.Cookie{
							Name:  "session",
					    	Value: encoded,
					    	Path:  "/",
					    	HttpOnly: true,
							}
					    http.SetCookie(w, cookie)
					} else {
						log.Printf("ApiHandler: Something went wrong when encoding the cookie values.\n%v\n", err)
						http.Error(w, "Oops. Something went wrong.", http.StatusInternalServerError)
					}
				}
				
				save(w, &res)
			case "User":
				// Check the session cookie to see if user is authenticated and session is valid
				respCookie := checkCookie(r)
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