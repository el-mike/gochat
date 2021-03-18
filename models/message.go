package models

// MESSAGE_RESOURCE - name of Message resource.
const MESSAGE_RESOURCE = "Message"

// MessageModel - Message DB model.
type MessageModel struct {
	BaseModel
}

// ResourceName - returns the name of Message resource.
func (mr *MessageModel) ResourceName() string {
	return MESSAGE_RESOURCE
}
