package web_server

import (
	"log"
	"time"
)

type (
	UserSession struct {
		Username   string
		Password   string
		Expiration time.Time
		State      State
		Cookie     string
	}
)

func NewUserSession(Username, Password string, Expiration time.Time, State State, Cookie string) *UserSession {
	log.Println("Create new user:", Username, "expiration:", Expiration, "with first state:", State, "cookie:", Cookie)
	return &UserSession{
		Username:   Username,
		Password:   Password,
		Expiration: Expiration,
		State:      State,
		Cookie:     Cookie,
	}
}

func (us *UserSession) SetNewState(NewState State) (valid bool) {
	log.Println("User:", us.Username, "previous state:", us.State)
	switch us.State {
	case StateMainPage:
		switch NewState {
		case StateNotAuthorised:
			// Valid new states
		default:
			log.Println("User:", us.Username, "invalid new state:", NewState)
			return
		}
	}
	us.State = NewState
	log.Println("User:", us.Username, "new state:", us.State)
	return true
}
