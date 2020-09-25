package models

//elkarto91@Author : Karthik
//GET & POST Handlers
import (
	"errors"
	"github.com/elkarto91/audition/common"
	"github.com/elkarto91/audition/databases"
	"net/http"
)

//Error Type
var (
	ErrIncorrectAdminCredentials = errors.New("admin credentials dont match")
	ErrKeyMissing                = errors.New("key missing in url parameters")
	ErrCommentMissing            = errors.New("comment missing in datbase")
)

//Handle Home Page - Login Page ; Since no cookie is being used this will be fine
func HandleHome(w http.ResponseWriter, r *http.Request) {

	logger := SetLoggerText()
	logger.Infoln("Audition Application Login Page Requested")
	HeaderXframeUtility(w, r)
	executeTemplate(w, "login.html", nil)

}

//Handling the dashboard for someone who logs in with account credentials
func HandleDashboard(w http.ResponseWriter, r *http.Request) {

	logger := SetLoggerText()
	logger.Infoln("Audition Application Dashboard Page")
	HeaderXframeUtility(w, r)

	_ = r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	user, err := databases.AuthenticateUser(username, password)
	if err != nil {
		logger.Errorln("Database updation error: ", err)
	}
	logger.Infoln("User Authenticated : ", user)

	comments, err := databases.ListAlComments()
	if err != nil {
		logger.Errorln("Database Comment error: ", err)
	}
	data := struct {
		Updates []*common.Comment
		User    string
	}{
		User:    user.Username,
		Updates: comments,
	}
	executeTemplate(w, "dashboard.html", data)
}

func HandleRegistry(w http.ResponseWriter, r *http.Request) {

	logger := SetLoggerText()
	logger.Infoln("Audition Application Register Page")
	HeaderXframeUtility(w, r)
	executeTemplate(w, "register.html", nil)

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
		executeTemplate(w, "login.html", nil)
	}
	if adminpass != "password" {
		logger.Error("Admin Credential Mismatch ", ErrIncorrectAdminCredentials)
		executeTemplate(w, "login.html", nil)
	}

	err := databases.RegisterUser(user)
	if err != nil {
		logger.Errorln("Database updation error: ", err)
	}
	logger.Infoln("Database updated for user: ", user.Username)
	executeTemplate(w, "login.html", nil)
}
func SubmitComment(w http.ResponseWriter, r *http.Request) {

	logger := SetLoggerText()
	logger.Infoln("Adding Comment")
	HeaderXframeUtility(w, r)

	var req common.Comment
	err := PostToInterface(r.Body, &req)
	if err != nil {
		logger.Errorln("ERROR -------> ", err.Error())
		ReturnJSONAPIErrorWithMessage(w, err.Error())
	}

	err = databases.AddComment(&req)
	if err != nil {
		logger.Errorln("Database updation error: ", err)
		ReturnJSONAPIErrorWithMessage(w, err.Error())
	}
	logger.Infoln("Database updated for user: ", req.Username)
	ReturnJSONAPISuccess(w, struct{ Success bool }{true})
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {

	logger := SetLoggerText()
	logger.Infoln("Deleting Comment")
	HeaderXframeUtility(w, r)

	var req common.Comment
	err := PostToInterface(r.Body, &req)
	if err != nil {
		logger.Errorln("ERROR -------> ", err.Error())
		ReturnJSONAPIErrorWithMessage(w, err.Error())
	}

	status, err := databases.DeleteCommentExist(req.CommentId)
	if err != nil {
		logger.Errorln("Database Deletion error: ", err)
		ReturnJSONAPIErrorWithMessage(w, err.Error())
	}
	logger.Infoln("Database updated for user: ", req.Username, " for comment id", req.CommentId, " with status", status)
	ReturnJSONAPISuccess(w, struct{ Success bool }{true})
}
func CheckComment(w http.ResponseWriter, r *http.Request) {

	logger := SetLoggerText()
	logger.Infoln("Checking Comment")
	HeaderXframeUtility(w, r)

	var req common.Comment
	err := PostToInterface(r.Body, &req)
	if err != nil {
		logger.Errorln("ERROR -------> ", err.Error())
		ReturnJSONAPIErrorWithMessage(w, err.Error())
	}
	flag := CheckPaliendrome(req.Comment)
	if flag == true {
		logger.Infoln("Comment : ", req.Comment, " is a palindrome :", flag)
		data := struct {
			Success bool
			Msg     string
		}{true, "Palindrome"}
		ReturnJSONAPISuccess(w, data)
	} else {
		logger.Infoln("Comment : ", req.Comment, " is not a paliendrome :", flag)
		data := struct {
			Success bool
			Msg     string
		}{true, "Not a Paliendrome"}
		ReturnJSONAPISuccess(w, data)
	}
}

func CheckPaliendrome(comment string) bool {

	startPointer := 0
	lengthOfString := len(comment)
	endPointer := lengthOfString - 1
	paliendromeFlag := true
	commentRune := []rune(comment)

	for startPointer <= endPointer {
		if commentRune[startPointer] == commentRune[endPointer] {
			startPointer++
			endPointer--
			continue
		} else {
			paliendromeFlag = false
			return paliendromeFlag
		}
	}
	return paliendromeFlag
}
func ViewAllComments(w http.ResponseWriter, r *http.Request) {

	logger := SetLoggerText()
	logger.Infoln("Audition Application Dashboard Page")
	HeaderXframeUtility(w, r)

	_ = r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	user, err := databases.AuthenticateUser(username, password)
	if err != nil {
		logger.Errorln("Database updation error: ", err)
		ReturnJSONAPIErrorWithMessage(w, err.Error())
	}
	logger.Infoln("User Authenticated : ", user)

	comments, err := databases.ListAlComments()
	if err != nil {
		logger.Errorln("Database Comment error: ", err)
		ReturnJSONAPIErrorWithMessage(w, err.Error())
	}
	ReturnJSONAPISuccess(w, comments)
}
