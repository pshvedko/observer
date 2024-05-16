package observer

import (
	"context"
	"github.com/mdigger/esl"
	"sync"
)

type Observer struct {
	mu       sync.Mutex
	closures map[chan<- []esl.Event]string
	watchers map[string][]chan<- []esl.Event
	backlogs map[string][]esl.Event
	ready    bool
}

func (o *Observer) Watch(ch chan<- []esl.Event, id string) {
	if ch == nil {
		panic("Watch with nil channel")
	}

	o.mu.Lock()
	defer o.mu.Unlock()

	if !o.ready {
		close(ch)
		return
	}

	o.closures[ch] = id
	o.watchers[id] = append(o.watchers[id], ch)

	ch <- o.backlogs[id]
}

func (o *Observer) Close(ch chan<- []esl.Event) {
	if ch == nil {
		panic("Close with nil channel")
	}

	o.mu.Lock()
	defer o.mu.Unlock()

	id, ok := o.closures[ch]
	if !ok {
		return
	}

	o.close(ch, id)
}

func (o *Observer) close(ch chan<- []esl.Event, id string) {
	watchers, ok := o.watchers[id]
	if ok {
		for i, w := range watchers {
			if w == ch {
				n := len(watchers)
				n--
				if n > 0 {
					if n > i {
						watchers[i] = watchers[n]
					}
					o.watchers[id] = watchers[:n]
					break
				}
				delete(o.watchers, id)
				break
			}
		}
	}

	delete(o.closures, ch)
	close(ch)
}

// Run ...
//
//	CHANNEL_ANSWER
//	CHANNEL_BRIDGE
//	CHANNEL_CALLSTATE
//	CHANNEL_CREATE
//	CHANNEL_DESTROY
//	CHANNEL_EXECUTE
//	CHANNEL_EXECUTE_COMPLETE
//	CHANNEL_HANGUP
//	CHANNEL_HANGUP_COMPLETE
//	CHANNEL_ORIGINATE
//	CHANNEL_OUTGOING
//	CHANNEL_PROGRESS
//	CHANNEL_PROGRESS_MEDIA
//	CHANNEL_STATE
//	CHANNEL_UNBRIDGE
func (o *Observer) Run(ctx context.Context, events <-chan esl.Event) {
	if ctx == nil {
		panic("Run with nil context")
	}
	if events == nil {
		panic("Run with nil channel")
	}

	defer func() {
		o.mu.Lock()
		o.ready = false

		for _, watchers := range o.watchers {
			for _, ch := range watchers {
				close(ch)
			}
		}

		o.mu.Unlock()
	}()

	o.mu.Lock()
	o.ready = true
	o.mu.Unlock()

	for {
		select {
		case <-ctx.Done():
			return

		case e, ok := <-events:
			if !ok {
				return
			}

			id := e.Get("Unique-ID")
			if id != "" {
				o.send(e, id)
			}
		}
	}
}

func (o *Observer) send(e esl.Event, id string) {
	o.mu.Lock()
	defer o.mu.Unlock()

	var end bool
	switch e.Get("Channel-State") {
	case "CS_DESTROY":
		end = true
		delete(o.backlogs, id)
	default:
		o.backlogs[id] = append(o.backlogs[id], e)
	}

	for _, ch := range o.watchers[id] {
		ch <- []esl.Event{e}
		if end {
			o.close(ch, id)
		}
	}
}

func New() *Observer {
	return &Observer{
		mu:       sync.Mutex{},
		closures: map[chan<- []esl.Event]string{},
		watchers: map[string][]chan<- []esl.Event{},
		backlogs: map[string][]esl.Event{},
		ready:    false,
	}
}
