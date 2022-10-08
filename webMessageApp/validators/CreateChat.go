package validators

type CreateChatInput struct {
	People   []uint `json:"person" binding:"required"`
	ChatName string `json:"chatName" binding:"required"`
}
