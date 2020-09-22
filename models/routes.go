package models

import (
	"html/template"
	"log"
	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {

	logger := SetLoggerText()
	logger.Infoln("Audition Application Login Page Requested")
	HeaderXframeUtility(w, r)

	t, err := template.ParseFiles("assets/userInterface/login.html") //parse the html file homepage.html
	if err != nil {                                                  // if there is an error
		log.Print("template parsing error: ", err) // log it
	}

	err = t.Execute(w, "") //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {        // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

func HandleRegistry(w http.ResponseWriter, r *http.Request) {

	logger := SetLoggerText()
	logger.Infoln("Audition Application Register Page")
	HeaderXframeUtility(w, r)

	t, err := template.ParseFiles("assets/userInterface/register.html") //parse the html file homepage.html
	if err != nil {                                                     // if there is an error
		log.Print("template parsing error: ", err) // log it
	}

	err = t.Execute(w, "") //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {        // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}
