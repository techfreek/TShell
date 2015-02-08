
package Twilio

import (
	"os"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"io"
)

type Cred struct {
	Sid string
	Auth string
} 

type Twil struct {
	Creds Cred
	HTTP *http.Client
}

type TwilData struct {
	PhoneNum string
	InMessage string
	OutMessage string
	MediaURL string
	Error bool
}

var LeftShark = "http://pbs.twimg.com/media/B80Q0_3CIAAWy90.jpg"
var From = "+15012297152"
var apiURL = "https://api.twilio.com/2010-04-01/"
var processing = make(chan TwilData)

//start server so we can get texts
var mux = http.NewServeMux()
mux.HandleFunc("/", gotText)
http.ListenAndServe(":8000", mux)

func Initialize(proc chan TwilData) (error, *Twil) {
	//Set the channel where texts are sent to be processed
	processing = proc;

	//Initialize a struct
	creds := Cred{}

	//Open Credentials file
	credFile, err := os.Open("twilioAPI.json")
	if err != nil {
		fmt.Println("Error opening file")
		return err, nil
	}
	// parse credentials
	jsonParser := json.NewDecoder(credFile)
	if err = jsonParser.Decode(&creds); err != nil {
		fmt.Println("Error parsing file")
		return err, nil
	}

	//Create struct
	twil := Twil{creds, http.DefaultClient}

	return nil, &twil
}

func (twil *Twil) SendText(data TwilData) {
	fmt.Println("Sending text");
	values := Valueify(data, twil)
	url := apiURL + "Accounts/" + twil.Creds.Sid + "/Messages.json"
	fmt.Println("Url: " + url);
	

	req, err := http.NewRequest("POST", url, strings.NewReader(values.Encode()))

	fmt.Println("Values: " + values.Encode())

	if err != nil {
		fmt.Println("Error creating request");
		return
	}

	req.SetBasicAuth(twil.Creds.Sid, twil.Creds.Auth)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	fmt.Println("Sending text")
	_, err = twil.HTTP.Do(req)
	fmt.Println("Text sent")

	if err != nil {
		fmt.Println("Error sending request: ")
		fmt.Println(err)
		return
	}

	return

}

func Valueify(data TwilData, twil *Twil) url.Values {
	form := url.Values{}

	form.Set("From", From)
	form.Set("To", data.PhoneNum)
	form.Set("Body", data.OutMessage)
	form.Set("MediaUrl", data.MediaURL)
	form.Set("ApplicationSid", twil.Creds.Sid)

	fmt.Println("Values: ")
	fmt.Println(form)

	return form
}

func GotText(res http.ResponseWriter, req *http.Request) {
	defer io.WriteString(res, "ACK")

	msg := TwilData{
		PhoneNum: req.FormValue("From"),
		InMessage: req.FormValue("Body"),
		OutMessage: "",
		MediaURL: "",
		Error: false,
	}

	fmt.Println(msg)

	processing <- msg
}
