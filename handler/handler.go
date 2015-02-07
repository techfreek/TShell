//This is the sterlizing handler, will be used to sterlize the incoming message to avoid any user shenanigans
//Accepts a map from the Twilio input/output handler that contains the submitting number/user [should be occupied],
//      the message from the user to be sterlized, and an empty response storage element

package TwitHandler

//import(
    //"fmt"
//)

func ping(pings chan<- string, msg string) {
    pings <- msg
}

//Normally blocking sterlization main function
func sterlhand (toProcess chan<- string, fromTwilio <-chan string) {
    msg := <-fromTwilio
    toProcess <- msg

//save out message from the incoming struct
//process against known unacceptable commands (? gut out appended commands (such as with && and |?)
//if message is clean/cleaned, remove original message and put new clean/cleaned instruction that was operated on
//send to toProcess

}


