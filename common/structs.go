package common

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Comment struct {
	Username  string `json:"username"`
	CommentId string `json:"comment_id"`
	Comment   string `json:"comment"`
}
