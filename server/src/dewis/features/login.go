package features

import (
	//"fmt"
	//"reflect"
	//"log"
	//"gopkg.in/mgo.v2"
   	//"gopkg.in/mgo.v2/bson"
    //"strconv"
    //"encoding/hex"
)

type LoginResp struct {
	status bool
}

func LoginHandler(dataMap RequestJSON) LoginResp {
	return LoginResp{false}
}
