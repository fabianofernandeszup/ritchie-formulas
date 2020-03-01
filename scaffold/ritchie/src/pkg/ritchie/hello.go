package ritchie

import (
	"github.com/fatih/color"
)
type Msg struct {
	Value string
}

func Send(msg Msg)  {
	color.Yellow(msg.Value)
}
