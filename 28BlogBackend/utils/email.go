package utils

import (
	"fmt"
	"os"
	"strconv"
	"math/rand"
	"time"

	"gopkg.in/gomail.v2"
)

func GenerateVerificationCode() string {
    rand.Seed(time.Now().UnixNano())
    code := rand.Intn(9000) + 1000 // Generates number between 1000-9999
    return strconv.Itoa(code)
}

func SendVerificationEmail(email, code string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_EMAIL"))
	m.SetHeader("TO", email)
	m.SetHeader("Subject", "Email Verification ")

	body := fmt.Sprintf(`
<h1>Email Verification</h1>
<p>Your verification code is <strong>%s</strong></p>
<p>If you didn't request this, please ignore this email.</p>
`, code)

	m.SetBody("text/html", body)

	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	d := gomail.NewDialer(os.Getenv("SMTP_HOST"), port, os.Getenv("SMTP_EMAIL"), os.Getenv("SMTP_PASSWORD"))

	return d.DialAndSend(m)
}

func SendPasswordResetEmail(email, resetToken string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_EMAIL"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Password Reset - Blog App")

	resetLink := fmt.Sprintf("http://localhost:3000/reset-password?token=%s", resetToken)

	body := fmt.Sprintf(`
        <h2>Password Reset</h2>
        <p>Click the link below to reset your password:</p>
        <p><a href="%s">Reset Password</a></p>
        <p>This link will expire in 1 hour.</p>
        <p>If you didn't request this, please ignore this email.</p>
    `, resetLink)

	m.SetBody("text/html", body)

	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	d := gomail.NewDialer(os.Getenv("SMTP_HOST"), port, os.Getenv("SMTP_EMAIL"), os.Getenv("SMTP_PASSWORD"))

	return d.DialAndSend(m)
}
