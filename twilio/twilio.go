
package Twilio

import (
	"os"
	"encoding/json"
	"fmt"
    //"net/http"
)

type Creds struct {
	Sid string
	Auth string
} 

type Twil struct {
	Sid string
	Auth string
	BaseURL string
}

type TwilData struct {
	PhoneNum string
	InMessage string
	OutMessage string
}

var LeftShark = "http://pbs.twimg.com/media/B80Q0_3CIAAWy90.jpg"
var From = "+15012297152"
var apiURL = "https://api.twilio.com/2010-04-01"

func Initialize() {
	creds := Creds{};

	credFile, err := os.Open("twilioAPI.json")
	if err != nil {
		fmt.Println("Error opening file")
	}

	jsonParser := json.NewDecoder(credFile)
	if err = jsonParser.Decode(&creds); err != nil {
		fmt.Println("Error parsing file")
	}
	fmt.Printf("%+v", creds);

	//twilio := gotwilio.NewTwilioClient(creds.Sid, creds.Auth)

	//to := "+14254175393"

	//message := "Hello world!"

	//twilio.SendMMS(From, to, message, LeftShark, "", "")
}
