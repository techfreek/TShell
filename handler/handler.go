//This is the sterlizing handler, will be used to sterlize the incoming message to avoid any user shenanigans
//Accepts a map from the Twilio input/output handler that contains the submitting number/user [should be occupied],
//      the message from the user to be sterlized, and an empty response storage element

//Change Log
//  6:13pm, Proof of concept works; can receive, modify, and send command on

package Sterilizer

import(
    "TwitterShell/twilio"
)

const no [11]string{"curl", "wget", "rm", "pwd", "chroot", "apt-get", "dpkg", "rmp", "sudo", "su", "chmod"}

//Normally blocking sterlization main function
func Sterlhand(fromTwilio <-chan Twilio.TwilData, toProcess chan<- Twilio.TwilData) {
    toOperate := <-fromTwilio 
    cleanedMessage := toOperate.InMessage
    cleanMessage(&cleanedMessage) //clean the message here
    toOperate.InMessage = cleanedMessage
    toProcess <- toOperate

//process against known unacceptable commands (? gut out appended commands (such as with && and |?)
}

//performs the sterlizing of the message
func cleanMessage(message *string) {
    //oldMessage := *message
    newMessage := "pwd"
    *message = newMessage
}

