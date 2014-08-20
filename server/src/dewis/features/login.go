package features

import (
	//"fmt"
	//"reflect"
	"log"
	"gopkg.in/mgo.v2"
   	"gopkg.in/mgo.v2/bson"
    "code.google.com/p/go.crypto/bcrypt"
    //"encoding/hex"
)

type LoginResp struct {
	Status bool
	Message string
}

type UserLogin struct {
	Username string	"email"
	Password string "password"
}

func loginHandler(req RequestJSON) LoginResp {
	var user UserLogin
	
	// Connecting to the database
    session, err := mgo.Dial("localhost");
    if err != nil {
    	log.Printf("Function getRecords: Error when opening connection to database.\n%v\n", err)
    	return LoginResp{false, "Error when opening connection to database"}
    }
    defer session.Close()

    // Querying the database
    conn := session.DB(DATABASE_NAME).C(USERS_COLLECTION)
    if err := conn.Find(bson.M{"email": req.Data["Username"]}).One(&user); err != nil {
    	log.Printf("Function loginHandler: Error when querying database. User not found.\n%v\n", err)
    	return LoginResp{false, "User not found."}
    }
    
    //Checking if password matches
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Data["Password"])); err == nil {
    	return LoginResp{true, ""}
    } else {
		return LoginResp{false, "Wrong password"}
	}
}
