package colored_logger

import (
	"fmt"
	 "io"
	 "io/ioutil"
	 "log"
	 "os"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"

)

var (
    Trace   *log.Logger
    Info    *log.Logger
    Warning *log.Logger
    Error   *log.Logger
)

const (
	ivMessage   = "message"
	ivFlowInfo  = "flowInfo"
	ivAddToFlow = "addToFlow"

	ovMessage = "message"
)


func Init(
    traceHandle io.Writer,
    infoHandle io.Writer,
    warningHandle io.Writer,
    errorHandle io.Writer) {

    Trace = log.New(traceHandle,
        "TRACE: ",
        log.Ldate|log.Ltime|log.Lshortfile)

    Info = log.New(infoHandle,
        "INFO: ",
        log.Ldate|log.Ltime|log.Lshortfile)

    Warning = log.New(warningHandle,
        "WARNING: ",
        log.Ldate|log.Ltime|log.Lshortfile)

    Error = log.New(errorHandle,
        "ERROR: ",
        log.Ldate|log.Ltime|log.Lshortfile)
}
func init() {

}

// LogActivity is an Activity that is used to log a message to the console
// inputs : {message, flowInfo}
// outputs: none
type CLogActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &CLogActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *CLogActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *CLogActivity) Eval(context activity.Context) (done bool, err error) {

	Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	 Trace.Println("yay!")
	 Info.Println("yay!")
	 Warning.Println("yay!")
	 Error.Println("yay!")

	//mv := context.GetInput(ivMessage)
	message, _ := context.GetInput(ivMessage).(string)

	msg := message


	msg = fmt.Sprintf("'%s' - FlowInstanceID [%s], Flow [%s], Task [%s]", msg,
			context.ActivityHost().ID(), context.ActivityHost().Name(), context.Name())



	context.SetOutput(ovMessage, msg)


	return true, nil
}
