
package Twilio

import (
	"github.com/sfreiberg/gotwilio"
	"os"
	"encoding/json"
	"fmt"
)

var creds struct {
	Sid string
	Auth string
}

var LeftShark = "http://pbs.twimg.com/media/B80Q0_3CIAAWy90.jpg"
var From = "+15012297152"

func Initialize() {
	credFile, err := os.Open("twilioAPI.json")
	if err != nil {
		fmt.Println("Error opening file")
	}

	jsonParser := json.NewDecoder(credFile)
	if err = jsonParser.Decode(&creds); err != nil {
		fmt.Println("Error parsing file")
	}
	fmt.Printf("%+v", creds);

	twilio := gotwilio.NewTwilioClient(creds.Sid, creds.Auth)

	to := "+14254175393"

	message := "Hello world!"

	twilio.SendMMS(From, to, message, LeftShark, "", "")

	//fmt.Printf("%s %s" credentials.Sid, credentials.Auth)
}
