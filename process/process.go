
package CmdProcessor

import(
    //"io/ioutil"
    "os/exec"
    "TwitterShell/twilio" //Twilio, Twilio.TwilData
)

//This package handles running an arbitrary command and then capturing the output to return to the requested user. At this point, the command is assumed to have been thoroughly cleaned/vetted to prevent security issues.

//The easiest approach to this appears to use the command format: "bash", "-c", "command to run".

//Fetch request from the appropriate channel, pull command out of struct, execute command and capture output, add output to the struct, and pass struct to the final return channel.

func RunProcess(FromSterilizer <-chan Twilio.TwilData, ToTwilio chan<- Twilio.TwilData) {
    toOperate := <- FromSterilizer
    if len([]rune(toOperate.MediaURL)) == 0 {
    //if no shark/image, then the command is clean
        commandString := toOperate.InMessage
        cmd := exec.Command("bash", "-c", commandString)
        out, err := cmd.Output()
        if err != nil {
            panic(err)
        }
        toOperate.OutMessage = (string(out)) //assumes the output is good, should crash before this otherwise
    }

    ToTwilio <- toOperate
}
