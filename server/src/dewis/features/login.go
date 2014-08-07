package features

import (
	//"fmt"
	//"reflect"
	"log"
	"gopkg.in/mgo.v2"
   	"gopkg.in/mgo.v2/bson"
    //"strconv"
    //"encoding/hex"
)

type LoginResp struct {
	Status bool
	Message string
}

type User struct {
	Username string	"email"
	Password string "password"
}

func loginHandler(dataMap RequestJSON) LoginResp {
	var user User
	
	// Connecting to the database
    session, err := mgo.Dial("localhost");
    if err != nil {
    	log.Printf("Function getRecords: Error when opening connection to database.\n%v\n", err)
    	return LoginResp{false, "Error when opening connection to database"}
    }
    defer session.Close()
    
    // Querying the database
    conn := session.DB("dewis").C("Users")
    if err := conn.Find(bson.M{"email": dataMap.Data["User"]}).One(&user); err != nil {
    	log.Printf("Function loginHandler: Error when querying database. User not found.\n%v\n", err)
    	return LoginResp{false, "User not found."}
    }
    
    //Checking if password matches
    if dataMap.Data["password"] == user.Password {
    	return LoginResp{true, ""}
    } else {
		return LoginResp{false, "Wrong password"}
	}
}
