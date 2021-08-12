package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func AutoRoute(w http.ResponseWriter, r *http.Request) {
	var event WebhookEvent
	body, err := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	err = json.Unmarshal(body, &event)
	if err != nil {
		log.Print(err)
	}

	switch event.Type {
	case MessageSent:
		UserMessage(w, r)
		return
	case StreamStarted, StreamStopped:
		StreamStartStop(w, r)
		return
	case UserNameChanged:
		var nameChangeEvent NameChangeWebhookEvent
		nameChangeEvent.EventData.Type = UserNameChanged
		body, err := ioutil.ReadAll(r.Body)
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		err = json.Unmarshal(body, &nameChangeEvent)
		if err != nil {
			log.Print(err)
		}

		switch nameChangeEvent.EventData.Type {
		// see https://github.com/owncast/owncast/issues/1302
		case "":
			UserJoin(w, r)
			return
		case UserNameChanged:
			UserNameChange(w, r)
			return
		}
		return

	default:
		log.Printf("Warning Unknown EventType")
		return
	}
}

func UserJoin(w http.ResponseWriter, r *http.Request) {
	var event NameChangeWebhookEvent
	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &event)
	if err != nil {
		log.Print(err)
	}

	go SendSystemMessage(GetUserJoinMessage(event.EventData.User.DisplayName), 1)
	if IsNewUser(event.EventData.User.Id) {
		go SendSystemMessage(GetBotIntroductionMessage(), 3)
		go SendSystemMessage(GetNameChangeHint(event.EventData.User.DisplayName), 5)
	}

	AddKnownUser(event.EventData.User.Id)
}

func UserNameChange(w http.ResponseWriter, r *http.Request) {
	var event NameChangeWebhookEvent
	body, err := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	err = json.Unmarshal(body, &event)
	if err != nil {
		log.Print(err)
	}

	if len(event.EventData.User.PreviousNames) == 1 {
		go SendSystemMessage(GetNamechangeMessage(event.EventData.User.DisplayName), 1)
	}
}

func UserMessage(w http.ResponseWriter, r *http.Request) {
	var event WebhookChatEvent
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &event)
	if err != nil {
		log.Print(err)
	}

	SendSystemMessage(`# Thank you for joining us here.
    I'm just a chat bot

    You can find ...

    <ul>
    <li>our source code on [Github](https://github.com/owncast/owncast)</li>
    <li>our docs and our website on [Owncast](https://owncast.online)</li>
    <li>us on [Rocket.Chat](https://owncast.rocket.chat/) - hang out with us</li>
    </ul>
    `)
}

func StreamStartStop(w http.ResponseWriter, r *http.Request) {
	var event WebhookStreamStartStopEvent
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &event)
	if err != nil {
		log.Print(err)
	}

	if event.Type == StreamStarted {
		go SendSystemMessage(GetStreamStartedMessage(), 5)
	} else {
		go SendSystemMessage(GetStreamStoppedMessage(), 5)
	}
}
