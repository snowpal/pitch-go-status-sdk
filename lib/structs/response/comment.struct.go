package response

type Comments struct {
	Comments []Comment `json:"comments"`
}

type Comment struct {
	ID          string `json:"id"`
	CommentText string `json:"commentText"`
}
