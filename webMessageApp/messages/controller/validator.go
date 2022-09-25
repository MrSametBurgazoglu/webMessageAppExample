package controller

type CreateMessageInput struct {
	PersonID uint   `json:"person" binding:"required"`
	Context  string `json:"context" binding:"required"`
	ChatID   uint   `json:"chat" binding:"required"`
}

type UpdateMessageInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type CreateChatPersonToPersonInput struct {
	Person1 uint `json:"person1" binding:"required"`
	Person2 uint `json:"person2" binding:"required"`
}
