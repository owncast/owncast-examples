package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func getApiUrlFor(endpoint string) string {
	return strings.TrimRight(DemoBotConfiguration.OwncastAddress, "/") + endpoint
}

func SendSystemMessage(message string, sendDelay int) {
	time.Sleep(time.Duration(sendDelay) * time.Second)
	postBody, _ := json.Marshal(map[string]string{
		"body": message,
	})
	responseBody := bytes.NewBuffer(postBody)
	req, _ := http.NewRequest("POST", getApiUrlFor("/api/integrations/chat/system"), responseBody)
	req.Header.Add("Authorization", "Bearer "+DemoBotConfiguration.AccessToken)
	req.Header.Add("ContentType", "application/json")

	client := &http.Client{}
	_, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
}

func SendSystemMessageToClient(clientId uint, message string, sendDelay int) {
	time.Sleep(time.Duration(sendDelay) * time.Second)
	var clientIdString string = strconv.FormatUint(uint64(clientId), 10)

	postBody, _ := json.Marshal(map[string]string{
		"body": message,
	})
	responseBody := bytes.NewBuffer(postBody)
	req, _ := http.NewRequest("POST", getApiUrlFor("/api/integrations/chat/system/client/"+clientIdString), responseBody)
	req.Header.Add("Authorization", "Bearer "+DemoBotConfiguration.AccessToken)
	req.Header.Add("ContentType", "application/json")

	client := &http.Client{}
	_, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
}
