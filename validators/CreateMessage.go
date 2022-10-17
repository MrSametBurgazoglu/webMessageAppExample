package validators

type CreateMessageInput struct {
	Person  uint   `json:"person" binding:"required"`
	Context string `json:"context" binding:"required"`
	Chat    uint   `json:"chat" binding:"required"`
}
