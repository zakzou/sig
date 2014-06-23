package sig

import (
	"fmt"
	"os"
)

type SignalHandler func(s os.Signal, arg interface{})

type SignalSet struct {
	m       map[os.Signal]SignalHandler
	Signals []os.Signal
}

func NewSignalSet() *SignalSet {
	ss := new(SignalSet)
	ss.m = make(map[os.Signal]SignalHandler)
	return ss
}

func (set *SignalSet) Register(sig os.Signal, handler SignalHandler) {
	if _, ok := set.m[sig]; !ok {
		set.m[sig] = handler
		set.Signals = append(set.Signals, sig)
	}
}

func (set *SignalSet) Handle(sig os.Signal, arg interface{}) error {
	if _, ok := set.m[sig]; ok {
		set.m[sig](sig, arg)
	} else {
		return fmt.Errorf("no handler available for signal %v", sig)
	}
	return nil
}
