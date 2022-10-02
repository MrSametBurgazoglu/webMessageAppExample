package validators

type CreatePersonInput struct {
	Nickname string `json:"title" binding:"required"`
	Author   string `json:"author" binding:"required"`
}
