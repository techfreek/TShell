package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"bytes"
)

type twilText struct {
	asid string				`json:"account_sid"`
	version string 			`json:"api_version"`
	body string 			`json:"body"`
	errcode string 			`json:"error_code""`
	errmsg string 			`json:"error_message"`
	segments string 		`json:"num_segments"`
	medias string 			`json:"num_media"`
	created string 			`json:"date_created"`
	sent string 			`json:"date_sent"`
	updated string 			`json:"date_updated"`
	direction string 		`json:"direction"`
	from string 			`json:"from"`
	price string 			`json:"string"`
	sid string 				`json:"sid"`
	status string 			`json:"status"`
	to string 				`json:"to"`
	uri string 				`json:"uri"`
}

func main() {
	url := "http://localhost:8000/command"

	text := twilText{
		asid: "", 
		version: "",
		body: "LeftShark",
		errcode: "", 
		errmsg: "", 
		segments: "",
		medias: "", 
		created: "",
		sent: "",
		updated: "",
		direction: "",
		from: "+1234567890",
		price: "",
		sid: "",
		status: "",
		to: "+14254175393",
		uri: "",
	}

	parsedText, _ := json.Marshal(text)

	var str = []byte(string(parsedText))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(str))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Error sending message")
	}

	defer res.Body.Close()

}