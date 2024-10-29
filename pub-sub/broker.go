package main

import (
	"fmt"
	"sync"
)

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

type Subscribers map[string]*Subscriber

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Broker represents a broker
type Broker struct {
	topics map[string]Subscribers
	mu     sync.RWMutex
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewBroker creates a new broker
func NewBroker() *Broker {
	return &Broker{
		topics: make(map[string]Subscribers),
	}
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Subscribe adds a subscriber to the broker's list of subscribers for a given topic. It then adds
// the topic to the subscriber's list of topics
func (b *Broker) Subscribe(s *Subscriber, topic string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.topics[topic] == nil {
		b.topics[topic] = make(Subscribers)
	}

	s.AddTopic(topic)
	b.topics[topic][s.id] = s

	fmt.Printf("Subscriber %s subscribed to topic %s\n", s.id, topic)
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Unsubscribe removes a subscriber from the broker's list of subscribers for a given topic. It
// then removes the topic from the subscriber's list of topics
func (b *Broker) Unsubscribe(s *Subscriber, topic string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.topics[topic] == nil {
		return
	}

	delete(b.topics[topic], s.id)
	s.RemoveTopic(topic)
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Publish sends a message to all subscribers of a given topic
func (b *Broker) Publish(topic string, message string) {
	b.mu.RLock()
	bTopics := b.topics[topic]
	b.mu.RUnlock()

	for _, s := range bTopics {
		m := NewMessage(message, topic)
		if !s.active {
			return
		}

		go (func(s *Subscriber) {
			s.Signal(m)
		})(s)
	}
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Broadcast sends a message to all subscribers of selected topics
func (b *Broker) Broadcast(topics []string, message string) {
	for _, topic := range topics {
		b.Publish(topic, message)
	}
}
