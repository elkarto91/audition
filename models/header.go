package models

//elkarto91@Author : Karthik
//Preventing Cross Side Scripting - Standard

import "net/http"

func HeaderXframeUtility(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("X-Frame-Options", "sameorigin")
	w.Header().Set("X-XSS-Protection", "1")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	return

}
