package chat

type NewChatDto struct {
	Fullname string `query:"fullname"`
	Email string `query:"email"`
}

type MailStrategyDto struct {
	Fullname string `json:"fullname"`
	Email string `json:"email"`
	Message string `json:"message"`
}

type OldChatDto struct {
	ChatID string `param:"id"`
}
