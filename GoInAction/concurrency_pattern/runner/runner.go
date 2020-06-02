package runner

import (
	"errors"
	"os"
	"time"
)

type Runner struct {
	interrupt chan os.Signal
	compelete chan error
	timeout   <-chan time.Time
	tasks     []func(int)
}

var ErrTimeout = errors.New("received timeout")
var ErrInterrupt = errors.New("received interrupt")

func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		compelete: make(chan error),
		timeout:   time.After(d),
	}
}
