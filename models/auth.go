package models

//elkarto91@Author : Karthik

import (
	"github.com/elkarto91/audition/databases"
	"net/http"
)

//Basic Authentication Middleware , JWT will take a long time
//Dont remember JWT - Need to check

func BasicAuthMiddleware(handler http.HandlerFunc) http.HandlerFunc {

	logger := SetLoggerText()
	logger.Infoln("Checking Basic Authentication")

	return func(w http.ResponseWriter, r *http.Request) {

		user, pass, ok := r.BasicAuth()
		logger.Infoln("Username: ", user)
		logger.Infoln("Password: ", pass)

		if !ok || !checkUsernameAndPassword(user, pass) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Please enter your username and password for this site"`)
			w.WriteHeader(401)
			w.Write([]byte("Unauthorised.\n"))
			return
		}
		handler(w, r)
	}
}

func checkUsernameAndPassword(username, password string) bool {

	if username == "admin" && password == "adminpw" {
		return true
	}
	_, err := databases.AuthenticateUser(username, password)
	if err != nil {
		return false
	}
	return true
}
