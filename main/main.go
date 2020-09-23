package main

import (
	"github.com/elkarto91/audition/databases"
	"github.com/elkarto91/audition/models"

	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {

	logger := models.SetLoggerText()
	logger.Infoln("Audition Application Has Started")

	dir, err := os.Getwd()
	if err != nil {
		logger.Errorln("Error", err.Error())
	}
	logger.Infoln("Directory ", dir)

	err = databases.InitMongo()
	if err != nil {
		logger.Errorln("Error", err.Error())
	}
	r := mux.NewRouter()
	r.HandleFunc("/login", models.HandleHome).Methods("GET")
	r.HandleFunc("/login", models.HandleDashboard).Methods("POST")

	r.HandleFunc("/register", models.HandleRegistry).Methods("GET")
	r.HandleFunc("/register", models.RegisterUser).Methods("POST")
	r.HandleFunc("/submitComment", models.SubmitComment).Methods("POST")

	//r.HandleFunc("/dashboard", models.RegisterUser).Methods("GET")

	/*	staticFileDirectory := http.FileServer(http.Dir("/assets/static"))
		r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", staticFileDirectory))*/
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		logger.Fatalln("Server failed")
	}

}
