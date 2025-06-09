package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/vaibhavkothari33/login/utils"
)

type Login struct {
	HashedPassword string
	SessionToken   string
	CSRFToken      string
}

// key is the username
var users = map[string]Login{}

func main() {
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/protected", protected)
	http.ListenAndServe(":4000", nil)

}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if len(username) < 8 || len(password) < 8 {
		http.Error(w, "Username and password must be at least 8 characters", http.StatusNotAcceptable)
		return
	}

	if _, exists := users[username]; exists {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	users[username] = Login{
		HashedPassword: hashedPassword,
	}

	fmt.Fprintln(w, "✅ User registered successfully")
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	user, ok := users[username]
	if !ok || !utils.CheckPasswordHash(password, user.HashedPassword) {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	sessionToken := utils.GenerateToken(32)
	csrfToken := utils.GenerateToken(32)

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    csrfToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: false,
	})

	user.SessionToken = sessionToken
	user.CSRFToken = csrfToken
	users[username] = user

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "✅ Login successful")
}

func logout(w http.ResponseWriter, r *http.Request) {

}
func protected(w http.ResponseWriter, r *http.Request) {

}
