package main

import (
    "fmt"
    "TwitterShell/handler" //this would be used as TwitHandler
)

type Datacontainer struct {
    phoneNum string
    inMessage string
    outMessage string
}


func main() {

    //initializes the twilio send/receive function
    //initializes the handler sterlizer function
    //initializes the process function
    fmt.Printf("Hello world\n")

    fakeData := Datacontainer{phoneNum: "555-555-5555", inMessage: "I like trains!"}
    
    hand := make(chan Datacontainer)
    demo := make(chan Datacontainer)

    TwitHandler.Sterlhand(demo, hand)
    
    fmt.Println(<-demo)

}

