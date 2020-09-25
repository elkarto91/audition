package models

import (
	"github.com/elkarto91/audition/common"
	"github.com/elkarto91/audition/databases"
	"log"
	"net/http"
)

func RegisterUserAPI(w http.ResponseWriter, r *http.Request) {

	logger := SetLoggerText()
	logger.Infoln("Audition Application Register User REST Call")

	adminName, ok := r.URL.Query()["adminun"]
	if !ok || len(adminName[0]) < 1 {
		logger.Errorln("Url Param 'Admin User Name' is missing")
		ReturnJSONAPIErrorWithMessage(w, ErrKeyMissing)

	}
	adminPass, ok := r.URL.Query()["adminpass"]
	if !ok || len(adminPass[0]) < 1 {
		logger.Errorln("Url Param 'Admin Password' is missing")
		ReturnJSONAPIErrorWithMessage(w, ErrKeyMissing)
	}

	log.Println("The Admin Data Received Are : ", adminName, adminPass)

	userName, ok := r.URL.Query()["username"]
	if !ok || len(userName[0]) < 1 {
		logger.Errorln("Url Param 'User Name' is missing")
		ReturnJSONAPIErrorWithMessage(w, ErrKeyMissing)
	}

	userPass, ok := r.URL.Query()["userpass"]
	if !ok || len(userPass[0]) < 1 {
		logger.Errorln("Url Param 'User Password' is missing")
		ReturnJSONAPIErrorWithMessage(w, ErrKeyMissing)
	}
	log.Println("The Admin Data Received Are : ", userName, userPass)

	if adminName[0] != "admin" {
		logger.Error("Admin Credential Mismatch ", ErrIncorrectAdminCredentials)
		ReturnJSONAPIErrorWithMessage(w, ErrIncorrectAdminCredentials)

	}
	if adminPass[0] != "password" {
		logger.Error("Admin Credential Mismatch ", ErrIncorrectAdminCredentials)
		ReturnJSONAPIErrorWithMessage(w, ErrIncorrectAdminCredentials)

	}
	user := &common.User{
		Username: userName[0],
		Password: userPass[0],
	}
	err := databases.RegisterUser(user)
	if err != nil {
		logger.Errorln("Database updation error: ", err)
		ReturnJSONAPIErrorWithMessage(w, err.Error())
	}

	ReturnJSONAPISuccess(w, struct{ Success bool }{true})
}

func SubmitCommentAPI(w http.ResponseWriter, r *http.Request) {

	logger := SetLoggerText()
	logger.Infoln("Adding Comment")

	comments, ok := r.URL.Query()["comment"]
	if !ok || len(comments[0]) < 1 {
		logger.Errorln("Url Param 'comment' is missing")
		ReturnJSONAPIErrorWithMessage(w, ErrKeyMissing)

	}

	commentID, ok := r.URL.Query()["commentid"]
	if !ok || len(commentID[0]) < 1 {
		logger.Errorln("Url Param 'commentid' is missing")
		ReturnJSONAPIErrorWithMessage(w, ErrKeyMissing)
	}

	userid, ok := r.URL.Query()["userid"]
	if !ok || len(userid[0]) < 1 {
		logger.Errorln("Url Param 'userid' is missing")
		ReturnJSONAPIErrorWithMessage(w, ErrKeyMissing)
	}

	req := common.Comment{
		userid[0],
		commentID[0],
		comments[0],
	}

	err := databases.AddComment(&req)
	if err != nil {
		logger.Errorln("Database updation error: ", err)
		ReturnJSONAPIErrorWithMessage(w, err.Error())
	}

	logger.Infoln("Database updated for user: ", req.Username)
	ReturnJSONAPISuccess(w, struct{ Success bool }{true})
}

func GetAllCommentAPI(w http.ResponseWriter, r *http.Request) {

	logger := SetLoggerText()
	logger.Infoln("Get All Comments")

	comments, err := databases.ListAlComments()
	if err != nil {
		logger.Errorln("Database Comment error: ", err)
		ReturnJSONAPIErrorWithMessage(w, err.Error())
	}
	ReturnJSONAPISuccess(w, comments)
}
func DeleteCommentAPI(w http.ResponseWriter, r *http.Request) {

	logger := SetLoggerText()
	logger.Infoln("Adding Comment")

	commentID, ok := r.URL.Query()["commentid"]
	if !ok || len(commentID[0]) < 1 {
		logger.Errorln("Url Param 'commentid' is missing")
		ReturnJSONAPIErrorWithMessage(w, ErrKeyMissing)
	}
	status, err := databases.DeleteCommentExist(commentID[0])
	if err != nil {
		logger.Errorln("Database Deletion error: ", err)
		ReturnJSONAPIErrorWithMessage(w, err.Error())
	}
	logger.Infoln("Comment Deleted with Status", status)
	ReturnJSONAPISuccess(w, status)
}

func CheckCommentAPI(w http.ResponseWriter, r *http.Request) {

	logger := SetLoggerText()
	logger.Infoln("Adding Comment")

	commentID, ok := r.URL.Query()["commentid"]
	if !ok || len(commentID[0]) < 1 {
		logger.Errorln("Url Param 'commentid' is missing")
		ReturnJSONAPIErrorWithMessage(w, ErrKeyMissing)
	}
	com, err := databases.GetCommentByCommentId(commentID[0])
	if err != nil {
		logger.Errorln("Comment is missing in DB ", err.Error())
		ReturnJSONAPIErrorWithMessage(w, err.Error())
	}
	if com != nil {

		flag := CheckPaliendrome(com.Comment)
		if flag == true {
			logger.Infoln("Comment : ", com.Comment, " is a palindrome :", flag)
			data := struct {
				Success bool
				Msg     string
			}{true, "Palindrome"}
			ReturnJSONAPISuccess(w, data)
		} else {
			logger.Infoln("Comment : ", com.Comment, " is not a palindrome :", flag)
			data := struct {
				Success bool
				Msg     string
			}{true, "Not a Palindrome"}
			ReturnJSONAPISuccess(w, data)
		}
	}
	logger.Errorln("Comment is missing in DB ", ErrCommentMissing)
	ReturnJSONAPIErrorWithMessage(w, ErrCommentMissing)
}
