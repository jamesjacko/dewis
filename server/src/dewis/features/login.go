package features

import (
    //"fmt"
    "time"
	"log"
    "code.google.com/p/go.crypto/bcrypt"
    "github.com/gorilla/securecookie"
    "net/http"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    //"encoding/hex"
)

type AuthResp struct {
    Status      bool
    Message     string
    CurrentUser UserSession
}

type UserLogin struct {
	Username string	"email"
	Password string "password"
}

type UserSession struct {
    Id          bson.ObjectId   "_id"
    FirstName   string          "first_name"
    LastName    string          "last_name"
    IsAdmin     bool            "is_admin"
    Email       string          "email"
    Avatar      string          "avatar"
}

/* TODO 
 * Check if session is valid
 */
func CheckCookie(r *http.Request) AuthResp {
    // Parsing cookie from request
    reqCookie, errCookie := r.Cookie("session")
    
    //There is no session cookie in the request
    if errCookie != nil {
        log.Printf("CheckCookie: Cookie not found.\n%v\n", errCookie)
        return AuthResp{false, "User not authenticated", UserSession{}}
    }

    var s = securecookie.New([]byte("dewis-hashkey-cookie"), []byte("encryption-key-dewis-hash78aw971"))
    session := Session{}

    if err := s.Decode("session", reqCookie.Value, &session); err != nil {
        log.Printf("CheckCookie: Something went wrong when decoding the cookie values.\n%v\n", err)
        return AuthResp{false, "Error when decrypting cookie value", UserSession{}}
    }

    //Returning the current user
    // Connecting to the database
    current := UserSession{}

    dbSession, err := mgo.Dial("localhost");
    if err != nil {
        log.Printf("Function CheckCookie: Error when opening connection to database.\n%v\n", err)
        return AuthResp{false, "Error when opening connection to database", UserSession{}}
    }
    defer dbSession.Close()

    // Querying the database
    conn := dbSession.DB(DATABASE_NAME).C(USERS_COLLECTION)
    if err := conn.Find(bson.M{"email": session.User}).One(&current); err != nil {
        log.Printf("Function CheckCookie: Error when querying database. User not found.\n%v\n", err)
        return AuthResp{false, "User not found.", UserSession{}}
    }

    return AuthResp{true, "", current}
}

func loginHandler(req RequestJSON, w http.ResponseWriter, r *http.Request) AuthResp {
    switch req.Action {
        case "Login":
           	var user UserLogin
            var res AuthResp
           	
           	// Connecting to the database
            session, err := mgo.Dial("localhost");
            if err != nil {
              	log.Printf("Function loginHandler: Error when opening connection to database.\n%v\n", err)
               	return AuthResp{false, "Error when opening connection to database", UserSession{}}
            }
            defer session.Close()

            // Querying the database
            conn := session.DB(DATABASE_NAME).C(USERS_COLLECTION)
            if err := conn.Find(bson.M{"email": req.Data["Username"]}).One(&user); err != nil {
              	log.Printf("Function loginHandler: Error when querying database. User not found.\n%v\n", err)
               	return AuthResp{false, "User not found.", UserSession{}}
            }
                
            //Checking if password matches
            if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Data["Password"])); err == nil {
               	res.Status  = true 
                res.Message = ""
            } else {
           		res.Status  = false 
                res.Message = "Wrong password"
          	}

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
                //If there is a session, update the last used time
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

            return res
        case "GetUser":
            res := CheckCookie(r)
            return res
    }
    return AuthResp{}
}
