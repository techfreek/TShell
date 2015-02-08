package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"bytes"
)

type twilText struct {
	MessageSid string 		`json:"MessageSid"`
	SmsSid string 			`json:"SmsSid"`
	AccountSid string 		`json:"AccountSid"`
	From string 			`json:"From"`
	To string 				`json:"To"`
	Body string 			`json:"Body"`
	NumMedia string 		`json:"NumMedia"`
}

func main() {
	url := "http://localhost:8000/command"

	text := twilText{
		asid: "1", 
		version: "2",
		body: "LeftShark",
		errcode: "3", 
		errmsg: "4", 
		segments: "5",
		medias: "6", 
		created: "7",
		sent: "8",
		updated: "9",
		direction: "10",
		from: "+1234567890",
		price: "11",
		sid: "12",
		status: "13",
		to: "+14254175393",
		uri: "14",
	}

	parsedText, _ := json.Marshal(text)

	var str = []byte(string(parsedText))

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(str))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Error sending message")
	}

	defer res.Body.Close()

}
