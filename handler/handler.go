//This is the sterlizing handler, will be used to sterlize the incoming message to avoid any user shenanigans
//Accepts a map from the Twilio input/output handler that contains the submitting number/user [should be occupied],
//      the message from the user to be sterlized, and an empty response storage element

//Change Log
//  6:13pm, Proof of concept works; can receive, modify, and send command on

package Sterilizer

import(
    "fmt"
    "TwitterShell/twilio"
    "strings"
    "math/rand"
)

var no = [12]string {"curl", "wget", "rm", "passwd", "chroot", "apt-get", "dpkg", "rmp", "sudo", "su", "chmod", "hexdump"}
var leftsharks = [6]string {
	"http://pbs.twimg.com/media/B80Q0_3CIAAWy90.jpg",
	"http://rack.0.mshcdn.com/media/ZgkyMDE1LzAyLzAzLzk4L2xlZnRzaGFya2luLjdhMWJjLmpwZwpwCXRodW1iCTk1MHg1MzQjCmUJanBn/50c64a3c/04e/leftsharkindiegogo.jpg",
	"http://cdn.baeblemusic.com/images/bblog/sharkone.jpg",
	"http://3dprint.com/wp-content/uploads/2015/02/leftsharkfeatured.jpg",
	"https://cdn1.lockerdome.com/uploads/3543c4041c5326700a535f1257bacd8b29d516f0bebcd7e2bfedbe4097b5b3dd_large",
	"http://media.giphy.com/media/yoJC2ESKKQZ5swxztS/giphy.gif"}

//Normally blocking sterlization main function
func Sterlhand(fromTwilio <-chan Twilio.TwilData, toProcess chan<- Twilio.TwilData, toTwilio chan<-Twilio.TwilData) {
    for toOperate :=  range fromTwilio {
        fmt.Printf("toOperate: %v\n", toOperate)
	    cleanedMessage := toOperate.InMessage
	    if !cleanMessage(cleanedMessage) || checkSpecial(cleanedMessage) {
	        if strings.Contains("Go Cougs!", cleanedMessage) {
	            toOperate.MediaURL = "http://hackathon.eecs.wsu.edu/images/cougar_logo.png" //sets hackathon image
	        } else {
	        	i := rand.Int() % len(leftsharks)
	        	toOperate.MediaURL = leftsharks[i]
	        }
	        fmt.Printf("toTwilio <- %v \n", toOperate)
	        toTwilio <- toOperate //unsafe message, return to send back, no chance of running command
	    } else {
		    toOperate.InMessage = cleanedMessage
		    toProcess <- toOperate
	    
	    }
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
    if strings.Contains("Go Cougs!", message) || strings.Contains("LShark", message) {
        return true
    }
    return false
}

