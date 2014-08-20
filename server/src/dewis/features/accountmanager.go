package features

import (
	//"fmt"
	//"reflect"
	"log"
	"gopkg.in/mgo.v2"
   	//"gopkg.in/mgo.v2/bson"
    "code.google.com/p/go.crypto/bcrypt"
)

type AccMngResponse struct {
	Status  bool
	Message string
}

type User struct {
	//Id			bson.ObjectId	`bson: "-"` 	// Ignore this field when inserting into database, as the database will insert it itself
	FirstName	string			"first_name"
	LastName	string			"last_name"
	IsAdmin		bool			"is_admin"
	Email		string			"email"
	Password	string			"password"		// Already encrypted password
	Avatar		string			"avatar"
}

func (user *User) newUser(dataMap map[string]string, res *AccMngResponse){
	// Hashing password, if hash fails, abort user
	hash, err := bcrypt.GenerateFromPassword([]byte(dataMap["Password"]), 12)
	if err != nil {
		log.Printf("Function addUser: Error when encrypting the password.\n%v\b", err)
		res.Status = false
		res.Message = "Error when encrypting the password"
		return
	}

	// Converting the IsAdmin flag from string to bool
	var flag bool
	switch dataMap["Isadmin"] {
		case "true":
			flag = true
		default:
			flag = false
		// default:
		// 	log.Printf("Function addUser: Error when converting isadmin to boolean.\n")
		// 	res.Status = false
		// 	res.Message = "Error when converting isadmin to boolean"
		// 	return
	}

	user.FirstName	= dataMap["Firstname"]
	user.LastName	= dataMap["Lastname"]
	user.IsAdmin	= flag
	user.Email		= dataMap["Email"]
	user.Password	= string(hash)
	user.Avatar		= dataMap["Avatar"]

	res.Status = true
	res.Message = ""
}

// TODO - Finish everything
func addUser(dataMap map[string]string, res *AccMngResponse) {
	// Connecting to the database
	session, err := mgo.Dial("localhost")
	if err != nil {
    	log.Printf("Function addUser: Error when opening connection to database.\n%v\n", err)
    	res.Status = false
		res.Message = "Error when opening connection to database."
    	return
    }
    defer session.Close()

    // Marshalling user data
    user := User{}
    user.newUser(dataMap, res)
    if res.Status == false {
    	return
    }

    // Openning collection
    conn := session.DB(DATABASE_NAME).C(USERS_COLLECTION)
  
    // Inserting an user in the database.
    if err := conn.Insert(&user); err != nil {
    	log.Printf("Function addUser: Error when adding entries to database.\n%v\n", err)
    	res.Status = false
		res.Message = "Error when adding entries to database"
    	return 
    }
}

func userHandler(req RequestJSON) AccMngResponse {
	res := AccMngResponse{}

	switch req.Action {
		case "addUser":
			addUser(req.Data, &res)
			return res
	}

	return res
}