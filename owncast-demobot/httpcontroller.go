package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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

	/* sleeps to make the repsonses more "human-like", just don't overdo it, otherwise the webhook will timeout */
	time.Sleep(2 * time.Second)
	SendSystemMessage(GetUserJoinMessage(event.EventData.User.DisplayName))
	time.Sleep(2 * time.Second)
	SendSystemMessage("May I introduce myself? I am a friendly Owncast Bot, programmed to help you around here!")
	time.Sleep(1 * time.Second)
	SendSystemMessage("Owncast has assigned you " + event.EventData.User.DisplayName + " as Username. If you don't like it, or would like to change it you can do so in the upper right corner")
}

func UserNameChange(w http.ResponseWriter, r *http.Request) {
	var event NameChangeWebhookEvent
	body, err := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	err = json.Unmarshal(body, &event)
	if err != nil {
		log.Print(err)
	}

	time.Sleep(2 * time.Second)
	SendSystemMessage(GetNamechangeMessage(event.EventData.User.DisplayName))
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
		SendSystemMessage(GetStreamStartedMessage())
	} else {
		SendSystemMessage(GetStreamStoppedMessage())
	}
}
