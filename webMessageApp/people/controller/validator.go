package controller

type CreatePersonInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdatePersonInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
