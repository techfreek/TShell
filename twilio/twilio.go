
package Twilio

import (
	"os"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"github.com/op/go-logging"
	//"io"
)

var log = logging.MustGetLogger("TShell")
var format = logging.MustStringFormatter(
	"%{color}%{time:15:04:05.000} %{shortfunc} -> %{level:.4s} %{id:03x}%{color:reset} %{message}",
)

type Cred struct {
	Sid string
	Auth string
} 

var twil struct {
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
var processing chan TwilData

func Initialize(proc chan TwilData) error {
	//Set the channel where texts are sent to be processed
	processing = proc;

	//Initialize a struct
	creds := Cred{}

	//Open Credentials file
	credFile, err := os.Open("twilioAPI.json")
	if err != nil {
		fmt.Println("Error opening file")
		return err
	}
	// parse credentials
	jsonParser := json.NewDecoder(credFile)
	if err = jsonParser.Decode(&creds); err != nil {
		fmt.Println("Error parsing file")
		return err
	}

	//Create struct
	twil.Creds = creds;
	twil.HTTP = http.DefaultClient;

	return nil
}

func SendText(toTwilio <-chan TwilData) {
	//fmt.Println("SendText()")
	for data := range toTwilio {	
		//fmt.Println("Sending text")
		values := Valueify(data)
		encodedValues := values.Encode()
		uri := apiURL + "Accounts/" + twil.Creds.Sid + "/Messages.json?" + encodedValues
		//fmt.Println("Uri: " + uri);
		

		req, err := http.NewRequest("POST", uri, strings.NewReader(values.Encode()))

		//fmt.Println("Values: " + values.Encode())

		if err != nil {
			fmt.Println("Error creating request");
			continue
		}

		req.SetBasicAuth(twil.Creds.Sid, twil.Creds.Auth)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		log.Info("[%s] Reply sent", data.PhoneNum)

		//fmt.Printf("Sending text: %v\n", data)
		_, err = twil.HTTP.Do(req)
		//fmt.Printf("Text sent: %s\n", data.PhoneNum)

		if err != nil {
			fmt.Println("Error sending request: ")
			fmt.Println(err)
			continue
		}
	}

}

func Valueify(data TwilData) url.Values {
	form := url.Values{}

	form.Set("From", From)
	form.Set("To", data.PhoneNum)
	form.Set("Body", data.OutMessage) 

	if data.MediaURL != "" {
		form.Set("MediaUrl", data.MediaURL)	
	}
	
	form.Set("ApplicationSid", twil.Creds.Sid)

	//fmt.Printf("Values: %v\n", form)

	return form
}

func GotText(res http.ResponseWriter, req *http.Request) {
	msg := TwilData{
		PhoneNum: req.FormValue("From"),
		InMessage: req.FormValue("Body"),
		OutMessage: "",
		MediaURL: "",
		Error: false,
	}

	log.Info("[%s] Request for: %s", msg.PhoneNum, msg.InMessage)

	//fmt.Printf("processing <- %v", msg)
	processing <- msg
}
