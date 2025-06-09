package main

import (
	"errors"
	"fmt"
	"log"
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

	log.Println("✅ Server started on http://localhost:4000")
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}
}

// ==================== REGISTER ====================
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
		log.Println("❌ Error hashing password:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	users[username] = Login{
		HashedPassword: hashedPassword,
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "✅ User registered successfully")
}

// ==================== LOGIN ====================
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
		HttpOnly: false, // CSRF needs to be readable by frontend JS
	})

	user.SessionToken = sessionToken
	user.CSRFToken = csrfToken
	users[username] = user

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "✅ Login successful")
}

// ==================== LOGOUT ====================
func logout(w http.ResponseWriter, r *http.Request) {
	if err := Authorize(r); err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Expire cookies
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HttpOnly: true,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HttpOnly: false,
	})

	username := r.FormValue("username")
	user, ok := users[username]
	if ok {
		user.SessionToken = ""
		user.CSRFToken = ""
		users[username] = user
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "✅ Logged out successfully!")
}

// ==================== PROTECTED ROUTE ====================
func protected(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if err := Authorize(r); err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	username := r.FormValue("username")
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "🛡️ CSRF Validation successful! Welcome, %s", username)
}

// ==================== AUTHORIZATION ====================
var AuthError = errors.New("unauthorized")

func Authorize(r *http.Request) error {
	username := r.FormValue("username")
	user, ok := users[username]
	if !ok {
		return AuthError
	}

	st, err := r.Cookie("session_token")
	if err != nil || st.Value != user.SessionToken {
		return AuthError
	}

	csrf := r.Header.Get("X-CSRF-Token")
	if csrf != user.CSRFToken || csrf == "" {
		return AuthError
	}

	return nil
}
