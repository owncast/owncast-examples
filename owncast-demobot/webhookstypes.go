package main

import (
	"time"
)

type EventType = string

const (
	MessageSent EventType = "CHAT"
	UserJoined EventType = "USER_JOINED"
	UserNameChanged EventType = "NAME_CHANGE"
	StreamStarted EventType = "STREAM_STARTED"
	StreamStopped EventType = "STREAM_STOPPED"
)

type User struct {
	Id				string				`json:"id"`
	AccessToken		string				`json:"-"`
	DisplayName		string				`json:"displayName"`
	DisplayColor	int					`json:"displayColor"`
	CreatedAt		time.Time			`json:"createdAt"`
	DisabledAt		time.Time			`json:"disabledAt,omitempty"`
	PreviousNames	[]string			`json:"previousNames"`
	NameChangedAt	time.Time			`json:"nameChangedAt,omitempty"`
}


type ChatMessage struct {
	User			User 				`json:"user,omitempty"`
	Body			string				`json:"body,omitempty"`
	RawBody			string				`json:"rawBody,omitempty"`
	ID				string				`json:"id,omitempty"`
	Visible			bool				`json:"visible"`
	Timestamp		time.Time			`json:"timestamp,omitempty"`
}

type StreamEvent struct {
	Summary			string				`json:"summary"`
	Name			string				`json:"name"`
	StreamTitle		string				`json:"streamTitle"`
}

type WebhookEvent struct {
	Type      		EventType			`json:"type"`
	EventData 		interface{}			`json:"eventData,omitempty"`
}

type WebhookChatEvent struct {
	Type      		EventType			`json:"type"`
	EventData 		ChatMessage			`json:"eventData,omitempty"`
}

type WebhookStreamStartStopEvent struct {
	Type			EventType			`json:"type"`
	EventData 		StreamEvent			`json:"eventData,omitempty"`
}

type NameChangeEvent struct {
	Type			string				`json:"type"`
	Id				string				`json:"id"`
	Timestamp		time.Time			`json:"timestamp"`
	User			User				`json:"user"`
}

type NameChangeWebhookEvent struct {
	Type			EventType			`json:"type"`
	EventData 		NameChangeEvent		`json:"eventData,omitempty"`
}