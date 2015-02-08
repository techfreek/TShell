//This is the sterlizing handler, will be used to sterlize the incoming message to avoid any user shenanigans
//Accepts a map from the Twilio input/output handler that contains the submitting number/user [should be occupied],
//      the message from the user to be sterlized, and an empty response storage element

package TwitHandler

import(
    "TwitterShell/main"
)

//Normally blocking sterlization main function
func Sterlhand(toProcess chan<- TwilData, fromTwilio <-chan TwilData) {
    toOperate := <-fromTwilio
    cleanedMessage := toOperate.inMessage
    cleanMessage(&cleanedMessage) //clean the message here
    toOperate.inMessage = cleanedMessage
    toProcess <- toOperate

//process against known unacceptable commands (? gut out appended commands (such as with && and |?)
}

//performs the sterlizing of the message
func cleanMessage(message *string) {
    //oldMessage := *message
    newMessage := "NEW Message :)"
    *message = newMessage
}

