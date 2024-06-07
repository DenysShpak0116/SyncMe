package controllers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/render"
	"github.com/jordan-wright/email"
)

var (
	emailApPassword = os.Getenv("MAIL_CODE")
	yourMail = "sme962304@gmail.com"
	hostAddress = "smtp.gmail.com"
	hostPort = "587"
)

func SendMailFunc(w http.ResponseWriter, r *http.Request) {
	var body struct {
		To string `json:"to"`
		Body string `json:"body"`
	}
	if err := render.Decode(r, &body); err != nil {
		http.Error(w, "Cannot decode: "+err.Error(), http.StatusBadRequest)
		return
	}
	emailInstance := email.NewEmail()
	emailInstance.From = yourMail
	emailInstance.To = []string{body.To}
	emailInstance.Subject = "Payment Bill"
	emailInstance.Text = []byte(body.Body)

	err := emailInstance.Send(fmt.Sprintf("%s:%s", hostAddress, hostPort), smtp.PlainAuth("", yourMail, emailApPassword, hostAddress))
	if err != nil {
			fmt.Println("There was an error sending the mail")
	}
	response := map[string]string{"message": "Email sent successfully"}
	json.NewEncoder(w).Encode(response)
}

func SendVerificationCodeFunc(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Email string `json:"email"`
	}
	if err := render.Decode(r, &body); err != nil {
		http.Error(w, "Cannot decode: "+err.Error(), http.StatusBadRequest)
		return
	}

	code := generateVerificationCode()


	emailInstance := email.NewEmail()
	emailInstance.From = yourMail
	emailInstance.To = []string{body.Email}
	emailInstance.Subject = "Verification Code"
	emailInstance.Text = []byte(code)
	err := emailInstance.Send(fmt.Sprintf("%s:%s", hostAddress, hostPort), smtp.PlainAuth("", yourMail, emailApPassword, hostAddress))
	if err != nil {
		http.Error(w, "Failed to send verification code: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(code)
}


func generateVerificationCode() string {
	rand.Seed(time.Now().UnixNano())
	code := ""
	for i := 0; i < 6; i++ {
		code += strconv.Itoa(rand.Intn(10))
	}
	return code
}