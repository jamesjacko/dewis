package features

import (
	"fmt"
	"errors"
	"time"
	"math/rand"
	"strconv"
)

/* This file implements dewis version of sessions. It uses the gorilla's session package
 * and builds a wrap in it for our purpose
 */

/* The session struct contains also the name of the user whose this session belongs.
 * This way, its easier to handle a session by username and we can use randoms SessionID,
 * making the hijack of a session much more complicated.
 */

/* Maximum time in seconds that a session is valid*/
const maxTime int64 = 60 * 60 * 24 * 5 //60s in 60min in 24 hours in x days

type Session struct {
	User 		string
	SessionID 	string
	lastUsed	int64
}

/* Alias so we can use object oriented paradigms */
type SessionStorage []Session

/* Slice containing all active sessions in runtime */
var store SessionStorage

func (s *SessionStorage) AddSession(session Session) {
	for i, _ := range *s {
		if (*s)[i].User == session.User {
			(*s)[i].SessionID = session.SessionID
			return
		}
	}

	//If there is no session for that username, create one.
	*s = append((*s), session)
}

func (s *SessionStorage) RemoveSession(username string) {
	for i, _ := range *s {
		if (*s)[i].User == username {
			(*s) = append((*s)[:i], (*s)[i+1:]...)
			return
		}
	}
}

/* Checks to see if a session is still valid.
 * A session is valid if it has been used in less than the maximum expiry date;
 */

 // TODO = Fix the index. This is currently wrong because we are removing one element from the array without changing the initial programmed lenght for the loop.
 // This will lead to a segmentation fault.
func (s *SessionStorage) CheckSessions() {
	for i, _ := range *s {
		dif := time.Now().Unix() - (*s)[i].lastUsed
		if dif >= maxTime {
			s.RemoveSession((*s)[i].User)
		}
	}
}

/* Generate a (hopefully) unique sessionID */
func (s *SessionStorage) GenerateSessionID(username string) string {
	rand.Seed(time.Now().UnixNano())
	str, _ := strconv.Atoi(username)
	return strconv.Itoa(rand.Intn(9999999999999999)) + strconv.Itoa(str) + strconv.Itoa(rand.Intn(9999999999999999))
}

func (s *SessionStorage) GetSessionID(username string) (string, error) {
	for _, element := range *s {
		if element.User == username {
			return element.SessionID, nil
		}
	}
	return "", errors.New("Sessions: session not found")
}

func (s *SessionStorage) GetUsername(sessionID string) (string, error) {
	for _, element := range *s {
		if element.SessionID == sessionID {
			return element.User, nil
		}
	}
	return "", errors.New("Sessions: session not found")	
}

func (s *SessionStorage) GetSession(SessionID string) *Session {
	for i, _ := range *s {
		if (*s)[i].SessionID == SessionID {
			return 	&(*s)[i]
		}
	}
	return &Session{}
}

func (s *SessionStorage) Print() {
	fmt.Printf("{ ")
	for _, element := range *s {
		fmt.Printf("[%s, %s, %d] ", element.User, element.SessionID, element.lastUsed)
	}
	fmt.Printf("}\n")
}