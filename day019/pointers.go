package main

import (
	"fmt"
	"strings"
)

type Message struct {
	Recipient string
	Text      string
}

func sendMessage(m Message) {
	fmt.Printf("To : %v\n", m.Recipient)
	fmt.Printf("Message : %v\n", m.Text)
}

func removeProfanity(message *string) {
	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "Dang", "****")
	messageVal = strings.ReplaceAll(messageVal, "Shoot", "*****")
	messageVal = strings.ReplaceAll(messageVal, "heck", "****")
	*message = messageVal
	// *message = strings.ReplaceAll(*message, "Dang", "****")
	// *message = strings.ReplaceAll(*message, "Shoot", "*****")
	// *message = strings.ReplaceAll(*message, "heck", "****")
}

func test(messages []string) {
	for _, message := range messages {
		removeProfanity(&message)
		fmt.Println(message)
	}
}

func main() {
	sendMessage(Message{
		Recipient: "Pradeep",
		Text:      "Hello i am a developer",
	})
	test([]string{"Dang", "My name", "is", "Pradeep Shoot", "a backend heck", "developer"})
	test([]string{"Hello", "My name", "is", "Pradeep Devka", "a backend go", "developer"})
	test([]string{"Hey", "My name", "is", "Pradeep Shoot", "a backend heck", "developer"})
}
