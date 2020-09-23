package models

import (
	"errors"
	"github.com/elkarto91/audition/common"
	"github.com/elkarto91/audition/databases"
	"net/http"
)

var (
	ErrIncorrectAdminCredentials = errors.New("admin credentials dont match")
)

func HandleHome(w http.ResponseWriter, r *http.Request) {

	logger := SetLoggerText()
	logger.Infoln("Audition Application Login Page Requested")
	HeaderXframeUtility(w, r)
	executeTemplate(w, "login.html")
}
func HandleDashboard(w http.ResponseWriter, r *http.Request) {

	logger := SetLoggerText()
	logger.Infoln("Audition Application Dashboard Page")
	HeaderXframeUtility(w, r)
	executeTemplate(w, "dashboard.html")
}

func HandleRegistry(w http.ResponseWriter, r *http.Request) {

	logger := SetLoggerText()
	logger.Infoln("Audition Application Register Page")
	HeaderXframeUtility(w, r)
	executeTemplate(w, "register.html")

}
func RegisterUser(w http.ResponseWriter, r *http.Request) {

	logger := SetLoggerText()
	logger.Infoln("Audition Application Register User")
	HeaderXframeUtility(w, r)

	_ = r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	admin := r.PostForm.Get("admin")
	adminpass := r.PostForm.Get("adminpass")

	user := &common.User{
		Username: username,
		Password: password,
	}
	logger.Infoln("Register User Credentials ", user)

	if admin != "admin" {
		logger.Error("Admin Credential Mismatch ", ErrIncorrectAdminCredentials)
		executeTemplate(w, "login.html")
	}
	if adminpass != "password" {
		logger.Error("Admin Credential Mismatch ", ErrIncorrectAdminCredentials)
		executeTemplate(w, "login.html")
	}

	err := databases.RegisterUser(user)
	if err != nil {
		logger.Errorln("Database updation error: ", err)
	}
	logger.Infoln("Database updated for user: ", user.Username)
	executeTemplate(w, "login.html")
}
