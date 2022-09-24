package controller

type CreateGroupInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateGroupInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
