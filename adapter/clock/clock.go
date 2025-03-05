package clock

import (
	"chat_server/usecase/output_port"
	"time"
)

type Clock struct {}

func New() output_port.Clock {
	return &Clock{}
}

func (c *Clock) Now() time.Time {
	return time.Now()
}