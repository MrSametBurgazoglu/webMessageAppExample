package controller

type CreateMessageInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateMessageInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
