package main

import (
	"container/list"
)

var knownUserIds *list.List

func InitializeBotMemory() {
	knownUserIds = list.New()
	knownUserIds.Init()
}

func AddKnownUser(userId string) {
	if knownUserIds.Len() > 30 {
		knownUserIds.Init()
	}

	if IsNewUser(userId) {
		knownUserIds.PushBack(userId)
	}
}

func IsNewUser(userId string) bool {
	for e := knownUserIds.Front(); e != nil; e = e.Next() {
		if e.Value == userId {
			return false
		}
	}
	return true
}
