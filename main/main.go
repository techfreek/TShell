package main

import (
    //"fmt"
   // "TwitterShell/handler" //Sterilizer
    "TwitterShell/twilio" //Twilio
    "TwitterShell/process" //CmdProcessor
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
    log.Info("Hello world")

    fakeData := Twilio.TwilData{PhoneNum: "555-555-5555", InMessage: "I like trains!"}

    //fmt.Println(fakeData)
    
    hand := make(chan Twilio.TwilData, 5)
    demo := make(chan Twilio.TwilData, 5)
    final := make(chan Twilio.TwilData, 5)

    demo <- fakeData

    go Sterilizer.Sterlhand(demo, hand)
    go CmdProcessor.RunProcess(hand, final)
    
    fmt.Println("Stuff here")
   
    fmt.Println(<-final)

    anotherFake := Twilio.TwilData{"+14254175393", "Go Cougs!", "Hackathon FTW", "http://pbs.twimg.com/media/B80Q0_3CIAAWy90.jpg"}

    _, twil := Twilio.Initialize()
    twil.SendText(anotherFake);

}

