package request

type CommentReqBody struct {
	CommentText string `json:"commentText"`
}

type CommentParam struct {
	StatusId  string
	CommentId string
}
