
package Twilio

import (
	"os"
	"encoding/json"
	"fmt"
    "net/http"
    "io/ioutil"
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
	PhoneNum, InMessage, OutMessage, MediaURL string
}

var LeftShark = "http://pbs.twimg.com/media/B80Q0_3CIAAWy90.jpg"
var From = "+15012297152"
var apiURL = "https://api.twilio.com/2010-04-01/"

func Initialize() (error, *Twil) {
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

func (twil *Twil) GetTexts() (error, []TwilData) {
	resp, err := twil.HTTP.Get(apiURL + "Accounts/" + twil.Creds.Sid + "/Messages.json")
	
	if err != nil {
		fmt.Println("Error requesting texts")
		return err, nil
	} else {
		defer resp.Body.Close();
		var contents string
		jsonParser := json.NewDecoder(resp.Body);
		jsonParser.Decode(&contents);
		fmt.Println(contents);
		contents, err := ioutil.ReadAll(resp.Body);
		if err != nil {
			fmt.Println("Error reading body")
			return err, nil
		}

		fmt.Println("Response: ");
		
		return err, nil	
	}
}

func (twil *Twil) SendText(data TwilData) {
	return
}
