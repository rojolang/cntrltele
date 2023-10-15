package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Update struct represents an update in the MongoDB database
type Update struct {
	ID        primitive.ObjectID `bson:"_id"`
	UserID    int                `bson:"user_id"`
	Message   string             `bson:"message"`
	Timestamp time.Time          `bson:"timestamp"`
}

// User struct represents a user in the MongoDB database
type User struct {
	ID         primitive.ObjectID `bson:"_id"`
	UserID     int                `bson:"user_id"`
	Username   string             `bson:"username"`
	AccessHash string             `bson:"access_hash"`
	FirstName  string             `bson:"first_name"`
	LastName   string             `bson:"last_name"`
	Messages   []Message          `bson:"messages"` // Array of messages
	PhotoURL   string             `bson:"photo_url"`
}

// Message struct represents a message from a user
type Message struct {
	ID       primitive.ObjectID `bson:"_id"`
	Text     string             `bson:"text"`
	Date     time.Time          `bson:"date"`
	SenderID int                `bson:"sender_id"`
}

// Admin struct represents an admin in the MongoDB database
type Admin struct {
	ID               primitive.ObjectID `bson:"_id"`
	SDKey            string             `bson:"sd_key"`
	AWSKey           string             `bson:"aws_key"`
	AWSRegion        string             `bson:"aws_region"`
	AWSSecret        string             `bson:"aws_secret"`
	SerperAPIKey     string             `bson:"serper_api_key"`
	OpenAIKey        string             `bson:"openai_api_key"`
	TelegramAPIID    string             `bson:"telegram_api_id"`
	TelegramAPIHash  string             `bson:"telegram_api_hash"`
	TelegramBotToken string             `bson:"telegram_bot_token"`
	TestBotToken     string             `bson:"test_bot_token"`
}

// Chat struct represents a chat in the MongoDB database
type Chat struct {
	ID           primitive.ObjectID `bson:"_id"`
	ChatID       int64              `bson:"chat_id"`
	Participants []User             `bson:"participants"` // Array of users
	Title        string             `bson:"title"`
	Messages     []Message          `bson:"messages"` // Array of messages
}

// Error struct represents an error in the MongoDB database
type Error struct {
	ID        primitive.ObjectID `bson:"_id"`
	Timestamp time.Time          `bson:"timestamp"`
	Message   string             `bson:"message"`
}

// Image struct represents an image in the MongoDB database
type Image struct {
	ID  primitive.ObjectID `bson:"_id"`
	URL string             `bson:"url"`
}
