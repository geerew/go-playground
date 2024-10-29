package main

import (
	"crypto/rand"
	"fmt"
	"sync"
)

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Subscriber represents a subscriber
type Subscriber struct {
	id       string
	messages chan *Message
	topics   map[string]bool
	active   bool
	mu       sync.RWMutex
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewSubscriber creates a new subscriber
func NewSubscriber() (string, *Subscriber) {

	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	id := fmt.Sprintf("subscriber-%x", b)

	return id, &Subscriber{
		id:       id,
		messages: make(chan *Message),
		topics:   make(map[string]bool),
		active:   true,
	}
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// AddTopic adds a topic to the subscriber's list of topics
func (s *Subscriber) AddTopic(topic string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.topics[topic] = true
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// RemoveTopic removes a topic from the subscriber's list of topics
func (s *Subscriber) RemoveTopic(topic string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.topics, topic)
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Signal pushes a message to the subscriber's message channel
func (s *Subscriber) Signal(message *Message) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.active {
		s.messages <- message
	}
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Close closes the subscriber's message channel
func (s *Subscriber) Close() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.active = false
	close(s.messages)
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Process processes messages sent to the subscriber
func (s *Subscriber) Process() {
	for {
		if message, ok := <-s.messages; ok {
			fmt.Printf("Subscriber %s received message: %s for topic: %s\n", s.id, message.payload, message.topic)
		}
	}
}
