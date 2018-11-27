package colored_logger

import (
	"fmt"
	"strconv"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
// 	"github.com/sirupsen/logrus"
)

// var logImpl = logrus.New()

const (
	ivMessage   = "message"
	ivFlowInfo  = "flowInfo"
	ivAddToFlow = "addToFlow"

	ovMessage = "message"
)

func init() {
	// activityLog.SetLogLevel(logger.InfoLevel)
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
// 	logImpl.Formatter = &logrus.TextFormatter{
// 			DisableColors: false,
// 			FullTimestamp: true,
// 	}

//   logImpl.Infof("\033[1;31m%s\033[0m","ErrorColor")

	//mv := context.GetInput(ivMessage)
	message, _ := context.GetInput(ivMessage).(string)
	flowInfo, _ := toBool(context.GetInput(ivFlowInfo))
	addToFlow, _ := toBool(context.GetInput(ivAddToFlow))

	msg := message

	if flowInfo {
		msg = fmt.Sprintf("'%s'", msg)
	}

	// activityLog.Info(msg)

	if addToFlow {
		context.SetOutput(ovMessage, msg)
	}

	return true, nil
}

func toBool(val interface{}) (bool, error) {

	b, ok := val.(bool)
	if !ok {
		s, ok := val.(string)

		if !ok {
			return false, fmt.Errorf("unable to convert to boolean")
		}

		var err error
		b, err = strconv.ParseBool(s)

		if err != nil {
			return false, err
		}
	}

	return b, nil
}
