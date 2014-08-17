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
const databaseName 			= "dewis"
const usersCol				= "Users"
const timelineRecordsCol	= "timelineRecords"

/* TODO
 * Figure out why it is so slow to return the login anwser
 * Modularize some "Login" code into functions
 * Create "last used"
*/
func ApiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var req RequestJSON;
		
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Printf("ApiHandler: Something went wrong when decoding the JSON object.\n%v\n", err)
			http.Error(w, "Bad request", http.StatusBadRequest)
		}
		
		// Call different handlers depending on type of request
		switch req.Request {
			case "Timeline":
				res := timelineHandler(req)
				if err := json.NewEncoder(w).Encode(res); err != nil {
					log.Printf("ApiHandler: Something went wrong when encoding the JSON object.\n%v\n", err)
					http.Error(w, "Oops. Something went wrong.", http.StatusInternalServerError)
				}
			case "Auth":
				res := loginHandler(req)

				if res.Status == true {
					var session *Session

					//If there is no session for that user, create one with a unique SessionID, add it to a cookie and send the cookie back 
					sessionID, err := store.GetSessionID(req.Data["Username"]); 
					if err != nil {
						//Crating session with unique sessionID and storing it
						sessionID = store.GenerateSessionID(req.Data["Username"])

						//Check if that sessionID already exists (shit happens) and create another if so
						_, err := store.GetUsername(sessionID);
						for err == nil {
							sessionID = store.GenerateSessionID(req.Data["Username"])
							_, err = store.GetUsername(sessionID);
						}

						session = &Session{req.Data["Username"], sessionID, time.Now().Unix()}
						store.AddSession(*session)
					} else {
						session = store.GetSession(sessionID)
						session.lastUsed = time.Now().Unix()
					}

					store.Print()

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