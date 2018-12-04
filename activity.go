package colored_logger

import (
	"fmt"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/fatih/color"

)

const (
	ivMessage   = "message"
	ivFlowInfo  = "flowInfo"
	ivAddToFlow = "addToFlow"
	ivLevel = "level"

	ovMessage = "message"
)

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


	//mv := context.GetInput(ivMessage)
	message, _ := context.GetInput(ivMessage).(string)
	level, _ := context.GetInput(ivLevel).(string)

	msg := message


	msg = fmt.Sprintf("%s [%s] - '%s'", level, context.Name(), msg)

	color.Cyan(msg)



	context.SetOutput(ovMessage, msg)

	return true, nil
}
