package model

import (
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/datatypes"
)

type Message interface {
	GetID() string
	Create(messageID, messageType, subject string, content map[string]any)
}

type MessageOne struct {
	ID        uint           `gorm:"primarykey"`
	MessageID string         `json:"message_id,omitempty" gorm:"unique_index"`
	Type      string         `json:"type,omitempty"`
	Subject   string         `json:"subject,omitempty" gorm:"unique_index"`
	Content   datatypes.JSON `json:"content,omitempty"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
}

func (m *MessageOne) Create(messageID, messageType, subject string, content map[string]any) {
	c := datatypes.JSON{}
	result, _ := json.Marshal(content)
	_ = c.Scan(result)
	m.MessageID = messageID
	m.Type = messageType
	m.Subject = subject
	m.Content = c
	m.CreatedAt = time.Now()
	return
}

func (m *MessageOne) GetID() string {
	return m.MessageID
}

type MessageTwo struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	MessageID string             `json:"message_id,omitempty" bson:"message_id,omitempty"`
	Type      string             `json:"type,omitempty" bson:"type,omitempty"`
	Subject   string             `json:"subject,omitempty" bson:"subject,omitempty"`
	Content   map[string]any     `json:"content,omitempty" bson:"content,omitempty"`
	CreatedAt string             `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

func (m *MessageTwo) Create(messageID, messageType, subject string, content map[string]any) {
	m.MessageID = messageID
	m.Type = messageType
	m.Subject = subject
	m.Content = content
	m.CreatedAt = time.Now().String()
	return
}

func (m *MessageTwo) GetID() string {
	return m.MessageID
}
