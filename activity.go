package colored_logger

import (
	"fmt"
	 "io"
	 "io/ioutil"
	 "log"
	 "os"
	
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
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

// activityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-colored_logger")

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
	activityLog.SetLogLevel(logger.InfoLevel)

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

	 Trace.Println("I have something standard to say")
	 Info.Println("Special Information")
	 Warning.Println("There is something you need to know about")
	 Error.Println("Something has failed")

	//mv := context.GetInput(ivMessage)
	message, _ := context.GetInput(ivMessage).(string)

	msg := message


	msg = fmt.Sprintf("'%s' - FlowInstanceID [%s], Flow [%s], Task [%s]", msg,
			context.ActivityHost().ID(), context.ActivityHost().Name(), context.Name())

	activityLog.Info(msg)


	context.SetOutput(ovMessage, msg)


	return true, nil
}
