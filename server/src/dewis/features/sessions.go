package features

import (
	"fmt"
	"github.com/gorilla/sessions"
)

/* This file implements dewis version of sessions. It uses the gorilla's session package
 * and builds a wrap in it for our purpose
 */

/* The session struct contains also the name of the user whose this session belongs.
 * This way, its easier to handle a session by username and we can use randoms SessionID,
 * making the hijack of a session much more complicated.
 */
type sessionStruct struct {
	user 	string
	session sessions.Session
}

/* Alias so we can use object oriented paradigms */
type sessionsSlice []sessionStruct

/* Slice containing all active sessions in runtime */
//var sessionStorage sessionsSlice

func (s sessionsSlice) addSession(username string, session sessions.Session) {
	for i, _ := range s {
		if s[i].user == username {
			s[i].session = session
			return
		}
	}

	//If there is no session for that username, create one.
	s = append(s, sessionStruct{username, session})
}

func (s sessionsSlice) print() {
	fmt.Printf("{ ")
	for _, element := range s {
		fmt.Printf("[ %s, %v ] ", element.user, element.session)
	}
	fmt.Printf("}\n")
}