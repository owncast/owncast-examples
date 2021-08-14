package main

import (
	"fmt"
	"math/rand"
)

func GetNamechangeMessage(newUsername string) string {
	messages := []string{
		"%s, I like your new name a lot.",
		"'%s'... What a cool name.",
		"%s. This is how heroes are called",
		"Great %s, you found the name-change function",
	}
	return fmt.Sprintf(messages[rand.Intn(len(messages))], newUsername)
}

func GetBotIntroductionMessage() string {
	return "May I introduce myself? I am a friendly Owncast Bot, programmed to help you around here! I'm not very smart yet."
}

func GetNameChangeHint(username string) string {
	return "Owncast has assigned you " + username + " as Username. If you don't like it, or want to change it, you can do so by clicking the '" + username + "' label in the upper right corner."
}

func GetFurtherResourcesMessage() string {
	return `# Here are some links for you:
- find us on our **[Website](https://owncast.online/)**
- join and chat with us on **[RocketChat](https://owncast.rocket.chat)**
- read our source code on **[Github](https://github.com/owncast)**
- check out our **[documentation](https://owncast.online/docs/)**
- see **[how you can get your personal Owncast up and running](https://owncast.online/quickstart/)**`
}

func GetBotHelpText() string {
	return `Beep bop, I'm a bot. [Here](https://github.com/unclearParadigm/owncast-examples/) you can find my source code.
Here are the commands that I understand:
- !bot - the message you are reading right now
- !links - send links to various Owncast resources
`
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
