package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

var topics = map[string]string{
	"BTC": "Bitcoin",
	"ETH": "Ethereum",
	"XRP": "Ripple",
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func main() {
	broker := NewBroker()

	// Create a subscriber and subscribe it to the BTC and ETH topics
	_, s1 := NewSubscriber()
	broker.Subscribe(s1, topics["BTC"])
	broker.Subscribe(s1, topics["ETH"])

	// Create a subscriber and subscribe it to the ETH and XRP topics
	_, s2 := NewSubscriber()
	broker.Subscribe(s2, topics["ETH"])
	broker.Subscribe(s2, topics["XRP"])

	// Start the subscribers processing (listening)
	go s1.Process()
	go s2.Process()

	// Publish messages to random topics
	go (func(broker *Broker) {
		values := make([]string, 0, len(topics))
		for _, v := range topics {
			values = append(values, v)
		}

		for {
			randomValue := values[rand.Intn(len(values))]

			message := fmt.Sprintf("%f", rand.Float64())
			go broker.Publish(randomValue, message)

			time.Sleep(time.Duration(rand.Intn(4)) * time.Second)
		}
	})(broker)

	// After 3 seconds broadcast a message to all subscribers
	go (func(broker *Broker) {
		time.Sleep(3 * time.Second)
		broker.Broadcast([]string{topics["BTC"], topics["ETH"], topics["XRP"]}, "Broadcast message")
	})(broker)

	fmt.Scanln()
	fmt.Println("Done!")

}
