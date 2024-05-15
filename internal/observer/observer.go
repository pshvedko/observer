package observer

import (
	"context"
	"github.com/mdigger/esl"
)

type Watch struct {
	id string
	ch chan<- esl.Event
}

type Observer struct {
	watch chan Watch
}

func (o *Observer) Watch(ch chan<- esl.Event, id string) {
	if ch == nil {
		panic("Watch with nil channel")
	}

	o.watch <- Watch{
		id: id,
		ch: ch,
	}
}

type State struct {
	watches []chan<- esl.Event
	events  []esl.Event
}

// RunContext ...
func (o *Observer) RunContext(ctx context.Context, events chan esl.Event) {
	if ctx == nil {
		panic("Run with nil context")
	}
	if events == nil {
		panic("Run with nil channel")
	}

	changes := map[string][]esl.Event{}
	closure := map[chan<- esl.Event]string{}
	watches := map[string][]chan<- esl.Event{}

	defer func() {
		for _, watch := range watches {
			for _, w := range watch {
				close(w)
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case w := <-o.watch:
			change, ok := changes[w.id]
			if !ok {
				close(w.ch)
				continue
			}

			for _, e := range change {
				w.ch <- e
			}

			watches[w.id] = append(watches[w.id], w.ch)
			closure[w.ch] = w.id
		case e, ok := <-events:
			if !ok {
				return
			}
			id := e.Get("Channel-Call-UUID")
			if id != "" {
				for _, w := range watches[id] {
					w <- e
				}

				changes[id] = append(changes[id], e)
			}
		}
	}
}

func New() *Observer {
	return &Observer{
		watch: make(chan Watch, 1),
	}
}
