package features

import (
	//"fmt"
	//"reflect"
	"log"
	"gopkg.in/mgo.v2"
   	"gopkg.in/mgo.v2/bson"
    "strconv"
    "encoding/hex"
)

/* Default quantity of records when no quantity is passed */
const defaultQnt = 5;

type QueryStruct struct {
	User 		bson.ObjectId	"user"
	Content 	string			"content"
	Time 		int				"time"
}

type User struct {
	Id 			bson.ObjectId	"_id"
	Name 		string			"first_name"
	Avatar 		string			"avatar"
}

/* Structure defining a record */
type Record struct {
	UserId		bson.ObjectId	"user"
	Content 	string			"content"
	Time 		int				"time"
	UserData	User
}

/* Response struct containing the list of records and the status of the query */
type RecordsResp struct {
	Status	bool
	Records	[]Record
}

/* Get records from the database */
func getRecords(res *RecordsResp, qntString string) {
	//Setting the default value of the query status to false.
	//If the query succeeds, at the end, we cange this status to true.
	res.Status = false

	qnt, err := strconv.Atoi(qntString)
	if err != nil {
		log.Printf("Function getRecords: Something went wrong when converting the quantity of records from string to int.\n %v\n", err)
		return
	}
	
	if qnt == 0 {
		qnt = defaultQnt;
	}
	
	// Connecting to the database
    session, err := mgo.Dial("localhost");
    if err != nil {
    	log.Printf("Function getRecords: Error when opening connection to database.\n %v\n", err)
    	return
    }
    defer session.Close()
    
    // Querying the database
    conn := session.DB("dewis").C("timelineRecords")
    if err := conn.Find(nil).All(&res.Records); err != nil {
    	log.Printf("Function getRecords: Error when querying database.\n %v\n", err)
    	return
    }
    
    // Getting the User Data
    conn = session.DB("dewis").C("Users")
    for i, _ := range res.Records {
    	if err := conn.FindId(res.Records[i].UserId).One(&res.Records[i].UserData); err != nil {
    		log.Printf("Function getRecords: Error when getting user data\n %v\n", err)
    		return
    	}
    }
    
    //Query succeeded
    res.Status = true
}

/* Add a record to the database */
func addRecords(dataMap map[string]string) bool {
	// Connecting to the database
	session, err := mgo.Dial("localhost")
	if err != nil {
    	log.Printf("Function addRecords: Error when opening connection to database.\n %#v\n", err)
    	return false
    }
    defer session.Close()
    
    // Inserting values in the database
    conn := session.DB("dewis").C("timelineRecords")
    
    // Inserting a QueryStruct in the database. First parameter is a Database Reference
    // with and unique ID from the user
   	hex, _ := hex.DecodeString(dataMap["User"])
    if err := conn.Insert(&QueryStruct{bson.ObjectId(hex), dataMap["Content"], 0}); err != nil {
    	log.Printf("Function addRecords: Error when adding entries to database.\n %#v\n", err)
    	return false
    }
    
    //Query succeded
    return true
}

func TimelineHandler(dataMap RequestJSON) RecordsResp {
	switch dataMap.Action {
		case "GetRecords":
			var res RecordsResp
			getRecords(&res, dataMap.Data["Quantity"])
			return res
		case "AddRecord":	
			return RecordsResp{addRecords(dataMap.Data), nil}
	}
	return RecordsResp{}
}
