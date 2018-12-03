package colored_logger

import (
	"fmt"
	"strconv"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

const (
	ivMessage = "message"

	ovMessage = "message"
)

type CLogActivity struct {
	metadata *activity.Metadata
}

func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &CLogActivity{metadata: metadata}
}

func (a *CLogActivity) Metadata() *activity.Metadata {
	return a.metadata
}

func (a *LogActivity) Eval(context activity.Context) (done bool, err error) {

	//mv := context.GetInput(ivMessage)
	message, _ := context.GetInput(ivMessage).(string)

	fmt.Printf("received message: %s", message)

	return true, nil
}
