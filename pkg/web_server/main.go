package web_server

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"math"
	"net/http"
	"time"
)

var (
	activeUsers = make(map[string]*UserSession)
)

func randomBase64String(l int) string {
	buff := make([]byte, int(math.Round(float64(l)/float64(1.33333333333))))
	rand.Read(buff)
	str := base64.RawURLEncoding.EncodeToString(buff)
	return str[:l] // strip the one extra byte we get from half the results.
}

func NewLoginPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// data, err := httputil.DumpRequest(r, true)
		// if err != nil {
		// 	log.Println("Can't dump request")
		// } else {
		// 	log.Println("Request:", string(data))
		// }

		// Check cookie
		activeUser := checkAuth(r)

		/**
		 * If user not authorised
		 */

		if activeUser == nil {
			switch r.Method {
			case http.MethodGet:
				loginPage(w)
				return
			case http.MethodPost:
				activeUser = authUser(w, r)
			default:
				log.Println("Forbidden method:", r.Method)
				w.Write([]byte("Forbidden method: " + r.Method))
				return
			}
		}

		/**
		 * If user authorised
		 */

		// If we got POST first we will try to change active state
		if r.Method == http.MethodPost {
			newStateValue := r.FormValue("new_state")
			if len(newStateValue) > 0 {
				switch newStateValue {
				case "logout":
					activeUser.SetNewState(StateNotAuthorised)
					delete(activeUsers, activeUser.Cookie)
				default:
					log.Println("Unknown new state:", newStateValue)
				}
			}
		}

		// Response with page for current active state
		switch activeUser.State {
		case StateNotAuthorised:
			loginPage(w)
		case StateMainPage:
			mainPage(w)
		default:
			w.Write([]byte("Unknown state"))
		}
	}
}

func authUser(w http.ResponseWriter, r *http.Request) *UserSession {
	username := r.FormValue("uname")
	password := r.FormValue("pwd")

	// expires := time.Now().AddDate(0, 0, 1)
	expires := time.Now().Add(time.Minute)
	cookie := &http.Cookie{
		Name:    "SESSION_ID",
		Path:    "/",
		Expires: expires,
		Value:   randomBase64String(64),
	}
	http.SetCookie(w, cookie)

	userSession := NewUserSession(username, password, expires, StateMainPage, cookie.Value)
	activeUsers[cookie.Value] = userSession

	return userSession
}

func checkAuth(r *http.Request) *UserSession {
	var cookie string
	// Check if we got cookie
	if len(r.Cookies()) == 0 {
		return nil
	}
	// Check if cookie exists
	cookie = r.Cookies()[0].Value
	userSession, ok := activeUsers[cookie]
	if !ok {
		return nil
	}
	// Check if cookie still valid
	if time.Now().After(userSession.Expiration) {
		delete(activeUsers, cookie)
		return nil
	}
	// Return user
	return userSession
}

func NewFaviconHandler() http.HandlerFunc {
	return faviconHandler
}

func NewAvatarHandler() http.HandlerFunc {
	return avatarHandler
}
