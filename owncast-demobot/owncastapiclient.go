package main

import (
	"log"
	"bytes"
	"strings"
	"net/http"
	"encoding/json"
 )

 func getApiUrlFor(endpoint string) string {
	 return strings.TrimRight(DemoBotConfiguration.OwncastAddress, "/") + endpoint
 }
 
 func SendSystemMessage(message string) {
	postBody, _ := json.Marshal(map[string]string{
	   "body":  message,
	})
	responseBody := bytes.NewBuffer(postBody)
	var bearer = "Bearer " + DemoBotConfiguration.AccessToken
    req, _ := http.NewRequest("POST", getApiUrlFor("/api/integrations/chat/system"), responseBody)
    req.Header.Add("Authorization", bearer)
	req.Header.Add("ContentType", "application/json")

    client := &http.Client{}
    _, err := client.Do(req)
    if err != nil {
        log.Println("Error on response.\n[ERROR] -", err)
    }
 }