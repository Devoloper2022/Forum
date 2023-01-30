package http

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"time"
// )

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != urlUser {
		h.notFound(w)
		return
	}
	if r.Method != "GET" {
		h.clientError(w, 400)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		h.notFound(w)
		return
	}

	ts, err := template.ParseFiles("./ui/templates/user/user.html")
	if err != nil {
		log.Printf("Get Post: Execute:%v", err)
		return
	}

	user, err := h.services.User.Get(int64(id))
	if err != nil {
		h.errorLog.Println(err.Error())
		h.serverError(w, err)
		return
	}

	err = ts.Execute(w, user)
	if err != nil {
		h.serverError(w, err)
		return
	}
}

func (h *Handler) profile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdatePost из page"))
}

// type Credentials struct {
// 	Password string `json:"password"`
// 	Username string `json:"username"`
// }

// func Signin(w http.ResponseWriter, r *http.Request) {
// 	var creds Credentials
// 	// Get the JSON body and decode into credentials
// 	err := json.NewDecoder(r.Body).Decode(&creds)
// 	if err != nil {
// 		// If the structure of the body is wrong, return an HTTP error
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	// Get the expected password from our in memory map
// 	expectedPassword, ok := users[creds.Username]

// 	// If a password exists for the given user
// 	// AND, if it is the same as the password we received, the we can move ahead
// 	// if NOT, then we return an "Unauthorized" status
// 	if !ok || expectedPassword != creds.Password {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}

// 	// Create a new random session token
// 	// we use the "github.com/google/uuid" library to generate UUIDs
// 	sessionToken := uuid.NewString()
// 	expiresAt := time.Now().Add(120 * time.Second)

// 	// Set the token in the session map, along with the session information
// 	sessions[sessionToken] = session{
// 		username: creds.Username,
// 		expiry:   expiresAt,
// 	}

// 	// Finally, we set the client cookie for "session_token" as the session token we just generated
// 	// we also set an expiry time of 120 seconds
// 	http.SetCookie(w, &http.Cookie{
// 		Name:    "session_token",
// 		Value:   sessionToken,
// 		Expires: expiresAt,
// 	})
// }

// func Welcome(w http.ResponseWriter, r *http.Request) {
// 	// We can obtain the session token from the requests cookies, which come with every request
// 	c, err := r.Cookie("session_token")
// 	if err != nil {
// 		if err == http.ErrNoCookie {
// 			// If the cookie is not set, return an unauthorized status
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}
// 		// For any other type of error, return a bad request status
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	sessionToken := c.Value

// 	// We then get the session from our session map
// 	userSession, exists := sessions[sessionToken]
// 	if !exists {
// 		// If the session token is not present in session map, return an unauthorized error
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}
// 	// If the session is present, but has expired, we can delete the session, and return
// 	// an unauthorized status
// 	if userSession.isExpired() {
// 		delete(sessions, sessionToken)
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}

// 	// If the session is valid, return the welcome message to the user
// 	w.Write([]byte(fmt.Sprintf("Welcome %s!", userSession.username)))
// }

// func Refresh(w http.ResponseWriter, r *http.Request) {
// 	// (BEGIN) The code from this point is the same as the first part of the `Welcome` route
// 	c, err := r.Cookie("session_token")
// 	if err != nil {
// 		if err == http.ErrNoCookie {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	sessionToken := c.Value

// 	userSession, exists := sessions[sessionToken]
// 	if !exists {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}
// 	if userSession.isExpired() {
// 		delete(sessions, sessionToken)
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}
// 	// (END) The code until this point is the same as the first part of the `Welcome` route

// 	// If the previous session is valid, create a new session token for the current user
// 	newSessionToken := uuid.NewString()
// 	expiresAt := time.Now().Add(120 * time.Second)

// 	// Set the token in the session map, along with the user whom it represents
// 	sessions[newSessionToken] = session{
// 		username: userSession.username,
// 		expiry:   expiresAt,
// 	}

// 	// Delete the older session token
// 	delete(sessions, sessionToken)

// 	// Set the new token as the users `session_token` cookie
// 	http.SetCookie(w, &http.Cookie{
// 		Name:    "session_token",
// 		Value:   newSessionToken,
// 		Expires: time.Now().Add(120 * time.Second),
// 	})
// }

// 	func Logout(w http.ResponseWriter, r *http.Request) {
// 		c, err := r.Cookie("session_token")
// 		if err != nil {
// 			if err == http.ErrNoCookie {
// 				// If the cookie is not set, return an unauthorized status
// 				w.WriteHeader(http.StatusUnauthorized)
// 				return
// 			}
// 			// For any other type of error, return a bad request status
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}
// 		sessionToken := c.Value

// 		// remove the users session from the session map
// 		delete(sessions, sessionToken)

// 		// We need to let the client know that the cookie is expired
// 		// In the response, we set the session token to an empty
// 		// value and set its expiry as the current time
// 		http.SetCookie(w, &http.Cookie{
// 			Name:    "session_token",
// 			Value:   "",
// 			Expires: time.Now(),
// 		})
// 	}
