package models

import (
	"encoding/json"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func executeTemplate(w http.ResponseWriter, file string, data interface{}) {
	htmlFile := "assets/userInterface/" + file
	t, err := template.ParseFiles(htmlFile) //parse the html file homepage.html
	if err != nil {                         // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, data) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {          // if there is an error
		log.Print("template executing error: ", err) //log it
	}
	return
}

func PostToInterface(Body io.ReadCloser, subject interface{}) error {
	u, _ := ioutil.ReadAll(Body)
	err := json.Unmarshal(u, &subject)
	return err
}

func ReturnJSONAPISuccess(w http.ResponseWriter, extra interface{}) {
	w.Header().Set("Content-Type", "application/json")
	j, _ := json.Marshal(struct {
		Status bool
		Msg    string
		Extra  interface{}
	}{Msg: "success", Status: true, Extra: extra})
	_, err := w.Write(j)
	if err != nil {
		log.Fatal("Error in JSON parse execution", err.Error())
	}
	return
}
