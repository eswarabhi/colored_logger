package colored_logger

import(
  "fmt"
  "strconv"

  "github.com/TIBCOSoftware/flogo-lib/core/activity"
  "github.com/TIBCOSoftware/flogo-lib/logger"
)

var activityLog = logger.GetLogger("abhi-colored-logger")

func init() {
  activityLog.SetLogLevel(logger.InfoLevel)
}

type CLogActivity struct {
  metadata *activity.Metadata
}

func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &CLogActivity{metadata: metadata}
}

func (a *CLogActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *CLogActivity) Eval(context activity.Context) (done bool, err error) {

	//mv := context.GetInput(ivMessage)
	message, _ := context.GetInput(ivMessage).(string)

	flowInfo, _ := toBool(context.GetInput(ivFlowInfo))
	addToFlow, _ := toBool(context.GetInput(ivAddToFlow))

	msg := message

	if flowInfo {

		msg = fmt.Sprintf("'%s' - FlowInstanceID [%s], Flow [%s], Task [%s]", msg,
			context.ActivityHost().ID(), context.ActivityHost().Name(), context.Name())
	}

	activityLog.Info(msg)

	if addToFlow {
		context.SetOutput(ovMessage, msg)
	}

	return true, nil
}
