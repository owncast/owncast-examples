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
	case UserJoined:
		UserJoin(w, r)
		return
	case UserNameChanged:
		UserNameChange(w, r)
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

	switch event.EventData.Body { /* bot commands*/
	case "!bot":
		go SendSystemMessage(GetBotHelpText(), 0)
	case "!links":
		go SendSystemMessage(GetFurtherResourcesMessage(), 0)
	}

	if strings.Contains(event.EventData.Body, "?") || strings.Contains(event.EventData.Body, "is ") { // User-Question
		go SendSystemMessage("Good question. I just can't answer it properly yet. I'm a bot, remember? Here's the best I came up with:", 1)
		go SendSystemMessage(GetFurtherResourcesMessage(), 2)
	}
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
