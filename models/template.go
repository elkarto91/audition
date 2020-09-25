package models

//elkarto91@Author : Karthik
import (
	"encoding/json"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

//Handle HTML Pages
func executeTemplate(w http.ResponseWriter, file string, data interface{}) {

	logger := SetLoggerText()
	logger.Infoln("Executing Template")

	htmlFile := "assets/userInterface/" + file

	//parse the html file
	t, err := template.ParseFiles(htmlFile)
	if err != nil {
		logger.Errorln("Template Parsing Error: ", err.Error())
	}

	//execute the template with data interface to jquery
	err = t.Execute(w, data)
	if err != nil {
		log.Print("Template Executing Error: ", err.Error())
	}
	return
}

//Utility method to read html body and json unmarshall into the structure
func PostToInterface(Body io.ReadCloser, subject interface{}) error {
	u, _ := ioutil.ReadAll(Body)
	err := json.Unmarshal(u, &subject)
	return err
}

//Successful processing of JSON API
func ReturnJSONAPISuccess(w http.ResponseWriter, extra interface{}) {
	logger := SetLoggerText()

	//Respond with JSON return value
	w.Header().Set("Content-Type", "application/json")
	j, _ := json.Marshal(struct {
		Status bool
		Msg    string
		Extra  interface{}
	}{Msg: "success", Status: true, Extra: extra})

	//Write content into http Responsewritter
	_, err := w.Write(j)
	if err != nil {
		logger.Errorln("Error in JSON parse execution", err.Error())
	}
	return
}

//Return Failure of Operation for API
func ReturnJSONAPIErrorWithMessage(w http.ResponseWriter, extra interface{}) {

	logger := SetLoggerText()

	w.Header().Set("Content-Type", "application/json")
	j, _ := json.Marshal(struct {
		Status bool
		Msg    string
		Extra  interface{}
	}{Msg: "false", Status: false, Extra: extra})
	_, err := w.Write(j)
	if err != nil {
		logger.Errorln("Error in JSON parse execution", err.Error())
	}
	return
}
