//This is the sterlizing handler, will be used to sterlize the incoming message to avoid any user shenanigans
//Accepts a map from the Twilio input/output handler that contains the submitting number/user [should be occupied],
//      the message from the user to be sterlized, and an empty response storage element

//Change Log
//  6:13pm, Proof of concept works; can receive, modify, and send command on

package Sterilizer

import(
    //"fmt"
    "TwitterShell/twilio"
    "strings"
)

var no = [11]string {"curl", "wget", "rm", "pwd", "chroot", "apt-get", "dpkg", "rmp", "sudo", "su", "chmod"}

//Normally blocking sterlization main function
func Sterlhand(fromTwilio <-chan Twilio.TwilData, toProcess chan<- Twilio.TwilData, toTwilio chan<-Twilio.TwilData) {
    toOperate := <-fromTwilio 
    cleanedMessage := toOperate.InMessage
    if !cleanMessage(cleanedMessage) || checkSpecial(cleanedMessage) {
        if strings.Contains("Go Cougs!", cleanedMessage) {
            toOperate.MediaURL = "http://pbs.twimg.com/media/B80Q0_3CIAAWy90.jpg" //sets shark
        } else {
            toOperate.MediaURL = "http://pbs.twimg.com/media/B80Q0_3CIAAWy90.jpg"
        }
        toTwilio <- toOperate //unsafe message, return to send back, no chance of running command
    } else {
    toOperate.InMessage = cleanedMessage
    toProcess <- toOperate
    }

//process against known unacceptable commands (? gut out appended commands (such as with && and |?)
}

//performs the sterlizing of the message
func cleanMessage(message string) bool{
    for _, check := range no {
        if strings.Contains(check, message) {
            return false
        }
    }
    return true
}

func checkSpecial(message string) bool {
    //check for "Go Cougs!" and "Left Shark"
    if strings.Contains("Go Cougs!", message) || strings.Contains("Left Shark", message) {
        return true
    }
    return false
}

