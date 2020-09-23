package models

import (
	"html/template"
	"log"
	"net/http"
)

func executeTemplate(w http.ResponseWriter, file string) {
	htmlFile := "assets/userInterface/" + file
	t, err := template.ParseFiles(htmlFile) //parse the html file homepage.html
	if err != nil {                         // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, "") //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {        // if there is an error
		log.Print("template executing error: ", err) //log it
	}
	return
}
