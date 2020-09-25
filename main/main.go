package main

//elkarto91@Author : Karthik
//Main file for setting up the golang server and routing for the APIs
//
import (
	"github.com/elkarto91/audition/databases"
	"github.com/elkarto91/audition/models"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {

	//Start Logger
	logger := models.SetLoggerText()
	logger.Infoln("Audition Application Has Started")

	//Establish Working Directory
	dir, err := os.Getwd()
	if err != nil {
		logger.Errorln("OS Native Check Error", err.Error())
	}
	logger.Infoln("Working Directory ", dir)

	//Initialization of Mongo DB
	err = databases.InitMongo()
	if err != nil {
		logger.Errorln("Mongo DB Initialization Error", err.Error())
	}

	//Serve Routes
	logger.Infoln("Opening API Routes ")

	//Routes for UI
	r := mux.NewRouter()
	r.HandleFunc("/", models.HandleHome).Methods("GET")
	r.HandleFunc("/login", models.HandleHome).Methods("GET")
	r.HandleFunc("/login", models.HandleDashboard).Methods("POST")
	r.HandleFunc("/register", models.HandleRegistry).Methods("GET")
	r.HandleFunc("/register", models.RegisterUser).Methods("POST")

	//Routes for Action
	r.HandleFunc("/submitComment", models.SubmitComment).Methods("POST")
	r.HandleFunc("/deleteComment", models.DeleteComment).Methods("POST")
	r.HandleFunc("/checkComment", models.CheckComment).Methods("POST")
	r.HandleFunc("/viewAllComment", models.ViewAllComments).Methods("POST")

	//Routes for REST API Calls

	r.HandleFunc("/api/register", models.BasicAuthMiddleware(http.HandlerFunc(models.RegisterUserAPI))).Methods("POST")
	r.HandleFunc("/api/submitComment", models.BasicAuthMiddleware(http.HandlerFunc(models.SubmitCommentAPI))).Methods("POST")
	r.HandleFunc("/api/deleteComment", models.BasicAuthMiddleware(http.HandlerFunc(models.DeleteCommentAPI))).Methods("POST")
	r.HandleFunc("/api/checkComment", models.BasicAuthMiddleware(http.HandlerFunc(models.CheckCommentAPI))).Methods("POST")
	r.HandleFunc("/api/viewAllComment", models.BasicAuthMiddleware(http.HandlerFunc(models.GetAllCommentAPI))).Methods("GET")

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		logger.Fatalln("Server failed")
	}
}
