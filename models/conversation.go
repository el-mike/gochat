package models

import "github.com/google/uuid"

// CONVERSATION_RESOURCE - Name of Conversation resource.
const CONVERSATION_RESOURCE = "Conversation"

// ConversationModel - Conversation DB model.
type ConversationModel struct {
	BaseModel
	Users    []uuid.UUID     `json:"users"`
	Messages []*MessageModel `json:"messages"`
}

// ResourceName - returns the name of Conversation resource.
func (cm *ConversationModel) ResourceName() string {
	return CONVERSATION_RESOURCE
}
