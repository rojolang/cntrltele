package bot

type Update struct {
	MessageID int
	ChatID    int64
	UserID    int
	Update    *telebot.Update
}

type User struct {
	User     *telebot.User
	ChatID   int64
	Username string
}

type Message struct {
	// TODO: Define necessary fields
}

type CallbackQuery struct {
	// TODO: Define necessary fields
}

func (u *Update) GetChatID() int64 {
	// TODO: Implement this method
}

func (u *User) GetUsername() string {
	// TODO: Implement this method
}
