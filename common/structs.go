package common

//elkarto91@Author : Karthik
//User and Comment Data Structure

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Comment struct {
	Username  string `json:"username"`
	CommentId string `json:"comment_id"`
	Comment   string `json:"comment"`
}
