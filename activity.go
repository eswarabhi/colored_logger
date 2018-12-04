package colored_logger

import (
	"fmt"
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/fatih/color"

)

type Level struct {
	Trace, Debug, Info, Print, Warn, Error, Fatal string
}

const (
	ivMessage   = "message"
	ivFlowInfo  = "flowInfo"
	ivAddToFlow = "addToFlow"
	ivLevel = "level"

	ovMessage = "message"
)

var S_level Level = Level{
	"Trace", "Debug", "Info", "Print", "Warn", "Error", "Fatal",
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

	var CLog *color.Color

	t := time.Now()

	//mv := context.GetInput(ivMessage)
	message, _ := context.GetInput(ivMessage).(string)
	addToFlow, _ := toBool(context.GetInput(ivAddToFlow))
	level, _ := context.GetInput(ivLevel).(string)

	msg := message

	msg = fmt.Sprintf("%s %-6s [%s] - '%s'", t.Format("2006-01-02 15:04:05.000"), level, context.Name(), msg)

	switch level {
	case S_level.Trace:
		CLog = color.New(color.FgCyan)
	case S_level.Debug:
		CLog = color.New(color.FgBlue)
	case S_level.Info:
		CLog = color.New(color.FgGreen)
	case S_level.Print:
		CLog = color.New(color.FgWhite)
	case S_level.Warn:
		CLog = color.New(color.FgYellow)
	case S_level.Error:
		CLog = color.New(color.FgRed)
	case S_level.Fatal:
		CLog = color.New(color.FgMagenta)
	}

	CLog.Println(msg)

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
