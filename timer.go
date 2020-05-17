package chatbot

import "time"

// FuncOnTimer - function
type FuncOnTimer func()

// Timer - timer
type Timer struct {
	OnTimer FuncOnTimer
	ticket  *time.Ticker
}

// SetTimer - set timer
func SetTimer(t time.Duration, onTimer FuncOnTimer) *Timer {
	obj := &Timer{OnTimer: onTimer}

	obj.ticket = time.NewTicker(t)

	go func() {
		for {
			<-obj.ticket.C

			obj.OnTimer()
		}
	}()

	return obj
}

// Close - close timer
func (t *Timer) Close() {
	if t.ticket != nil {
		t.ticket.Stop()

		t.ticket = nil
	}
}
