package observer

import (
	"context"
	"fmt"
	"github.com/mdigger/esl"
)

func ExampleNew() {
	ctx := context.Background()

	events := make(chan esl.Event, 1)

	o := New()

	go o.Run(ctx, events)

	w := make(chan struct{})

	go func() {
		for _, e := range Events {
			events <- esl.NewEvent(e.Name, e.Headers, []byte{})
			if e.Done {
				w <- struct{}{}
				<-w
			}
		}
		close(events)
	}()

	<-w

	watch := make(chan []esl.Event, 1)

	o.Watch(watch, "baf54d6c-ecba-4183-b2aa-0f40b0df308f")

	close(w)

	for ee := range watch {
		for _, e := range ee {
			fmt.Println(e.Name(),
				e.Get("Channel-Call-State"),
				e.Get("Channel-Call-UUID"),
				e.Get("Channel-State"))
		}
	}

	fmt.Println("DONE")

	// Output:
	// DONE
}

type Event struct {
	Name    string
	Headers map[string]string
	Done    bool
}

var Events = []Event{{
	Headers: map[string]string{
		"Channel-Call-State": "DOWN",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_NEW",
	},
	Name: "CHANNEL_STATE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "DOWN",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_INIT",
	},
	Name: "CHANNEL_STATE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "DOWN",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_INIT",
	},
	Name: "CHANNEL_CREATE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_ROUTING",
	},
	Name: "CHANNEL_CALLSTATE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_ROUTING",
	},
	Name: "CHANNEL_STATE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_STATE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_EXECUTE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_EXECUTE_COMPLETE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_EXECUTE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_EXECUTE_COMPLETE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_EXECUTE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_EXECUTE_COMPLETE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_EXECUTE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_EXECUTE_COMPLETE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_EXECUTE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_EXECUTE_COMPLETE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_EXECUTE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_EXECUTE_COMPLETE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_EXECUTE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_EXECUTE_COMPLETE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_EXECUTE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_EXECUTE_COMPLETE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_EXECUTE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_EXECUTE_COMPLETE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_EXECUTE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_EXECUTE_COMPLETE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_EXECUTE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_EXECUTE_COMPLETE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_EXECUTE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_EXECUTE_COMPLETE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_EXECUTE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_EXECUTE_COMPLETE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_EXECUTE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_EXECUTE_COMPLETE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_EXECUTE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_EXECUTE_COMPLETE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_EXECUTE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_EXECUTE_COMPLETE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_EXECUTE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_EXECUTE_COMPLETE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_EXECUTE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "DOWN",
		"Channel-Call-UUID":  "baf54d6c-ecba-4183-b2aa-0f40b0df308f",
		"Channel-State":      "CS_NONE",
	},
	Name: "CHANNEL_OUTGOING",
	Done: true,
}, {
	Headers: map[string]string{
		"Channel-Call-State": "DOWN",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_INIT",
	},
	Name: "CHANNEL_STATE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "DOWN",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_INIT",
	},
	Name: "CHANNEL_CREATE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "DOWN",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_INIT",
	},
	Name: "CHANNEL_ORIGINATE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "DOWN",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_ROUTING",
	},
	Name: "CHANNEL_STATE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "DOWN",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_CONSUME_MEDIA",
	},
	Name: "CHANNEL_STATE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "DOWN",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_CONSUME_MEDIA",
	},
	Name: "CHANNEL_PROGRESS",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_CONSUME_MEDIA",
	},
	Name: "CHANNEL_CALLSTATE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_PROGRESS_MEDIA",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "EARLY",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_CALLSTATE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "RINGING",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_CONSUME_MEDIA",
	},
	Name: "CHANNEL_ANSWER",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "ACTIVE",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_CONSUME_MEDIA",
	},
	Name: "CHANNEL_CALLSTATE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "EARLY",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_ANSWER",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "ACTIVE",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_CALLSTATE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "ACTIVE",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_CONSUME_MEDIA",
	},
	Name: "CHANNEL_OUTGOING",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "ACTIVE",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_BRIDGE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "ACTIVE",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXCHANGE_MEDIA",
	},
	Name: "CHANNEL_STATE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "ACTIVE",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_HANGUP",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "ACTIVE",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXCHANGE_MEDIA",
	},
	Name: "CHANNEL_HANGUP",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "ACTIVE",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_HANGUP",
	},
	Name: "CHANNEL_STATE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "HANGUP",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_HANGUP",
	},
	Name: "CHANNEL_CALLSTATE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "ACTIVE",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_UNBRIDGE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "ACTIVE",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_EXECUTE",
	},
	Name: "CHANNEL_EXECUTE_COMPLETE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "ACTIVE",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_HANGUP",
	},
	Name: "CHANNEL_STATE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "HANGUP",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_HANGUP",
	},
	Name: "CHANNEL_CALLSTATE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "HANGUP",
		"Channel-Call-UUID":  "baf54d6c-ecba-4183-b2aa-0f40b0df308f",
		"Channel-State":      "CS_REPORTING",
	},
	Name: "CHANNEL_STATE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "HANGUP",
		"Channel-Call-UUID":  "baf54d6c-ecba-4183-b2aa-0f40b0df308f",
		"Channel-State":      "CS_REPORTING",
	},
	Name: "CHANNEL_HANGUP_COMPLETE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "HANGUP",
		"Channel-Call-UUID":  "baf54d6c-ecba-4183-b2aa-0f40b0df308f",
		"Channel-State":      "CS_REPORTING",
	},
	Name: "CHANNEL_DESTROY",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "HANGUP",
		"Channel-Call-UUID":  "baf54d6c-ecba-4183-b2aa-0f40b0df308f",
		"Channel-State":      "CS_DESTROY",
	},
	Name: "CHANNEL_STATE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "HANGUP",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_REPORTING",
	},
	Name: "CHANNEL_STATE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "HANGUP",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_REPORTING",
	},
	Name: "CHANNEL_HANGUP_COMPLETE",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "HANGUP",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_REPORTING",
	},
	Name: "CHANNEL_DESTROY",
}, {
	Headers: map[string]string{
		"Channel-Call-State": "HANGUP",
		"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
		"Channel-State":      "CS_DESTROY",
	},
	Name: "CHANNEL_STATE",
}}
