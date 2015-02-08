package main

import (
    //"fmt"
    //"TwitterShell/handler" //this would be used as TwitHandler
    "TwitterShell/twilio"
    //"TwitterShell/handler" //this would be used as TwitHandler
    "github.com/op/go-logging"
)

var log = logging.MustGetLogger("TShell")
var format = logging.MustStringFormatter(
	"%{color}%{time:15:04:05.000} %{shortfunc} -> %{level:.4s} %{id:03x}%{color:reset} %{message}",
)


func main() {
    //initializes the twilio send/receive function
    //initializes the handler sterlizer function
    //initializes the process function
    /*fmt.Printf("Hello world\n")
    fmt.Printf("Hello world\n")
    log.Info("Hello world")
    //fakeData := Datacontainer{phoneNum: "555-555-5555", inMessage: "I like trains!"}
    
    //hand := make(chan Datacontainer)
    //demo := make(chan Datacontainer)

    //TwitHandler.Sterlhand(demo, hand)
    
    fmt.Println(<-demo)*/

    _, twil := Twilio.Initialize()
   twil.GetTexts();

}

