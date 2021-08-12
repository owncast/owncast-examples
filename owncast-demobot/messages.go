package main

import (
	"fmt"
	"math/rand"
)

func GetNamechangeMessage(newUsername string) string {
	messages := []string{
		"%s, I like your new name a lot",
		"What a cool name.",
		"%s. This is how heroes are called",
	}
	return fmt.Sprintf(messages[rand.Intn(len(messages))], newUsername)
}

func GetStreamStartedMessage() string {
	messages := []string{"Let's get this party started", "There we go.", "Are you ready?", "Wow, we're live", "How exciting, we're ON AIR"}
	return messages[rand.Intn(len(messages))]
}

func GetStreamStoppedMessage() string {
	messages := []string{"Thank you for joining in", "That's it for today", "Good night everyone", "Cya", "Bye Bye"}
	return messages[rand.Intn(len(messages))]
}

func GetUserJoinMessage(username string) string {
	messages := []string{
		"Hello %s. Nice to see you around here",
		"Howdy %s. Glad you reached us",
		"Hey %s. Have an owncast-tastic day",
		"Welcome %s. It's nice to have you here",
		"Greetings %s. Happy to see you",
		"Bonjour %s. What a pleasure to meet you",
		"Hey there %s. How nice that you discovered Owncast",
		"Hi %s. What a concidence to see you here",
	}
	return fmt.Sprintf(messages[rand.Intn(len(messages))], username)
}
