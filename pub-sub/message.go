package main

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Message represents a message
type Message struct {
	topic   string
	payload string
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewMessage creates a new message
func NewMessage(payload, topic string) *Message {
	return &Message{
		topic:   topic,
		payload: payload,
	}
}
