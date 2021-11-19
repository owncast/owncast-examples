package main

import (
	"time"

	"github.com/muesli/cache2go"
)

var knownUserIds = cache2go.Cache("visitors")

func AddKnownUser(userId string) {
	if IsNewUser(userId) {
		knownUserIds.Add(userId, time.Hour*24*15, userId)
	}
}

func IsNewUser(userId string) bool {
	exists, _ := knownUserIds.Value(userId)
	return exists == nil
}
