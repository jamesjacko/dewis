package features

import (
	"fmt"
	"gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

/* res stands for response.
 *
 */

type Message struct {
	user 	string
	content string
	time 	int
}

func getMessages(res *struct, quantity int) {
	
}

func timelineHandler(dataMap *url.Value) struct {
	switch dataMap["action"] {
		case "getMessages":
			res := { status int, messages []Message }
			getMessages(&res, dataMap["qntMessages"])
			return res
	}
}
