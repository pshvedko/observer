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
			events <- esl.NewEvent(e["Event-Name"], e, []byte{})
			if _, ok := e["Begin"]; ok {
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
				e.Get("Unique-ID"),
				e.Get("Channel-State"))
		}
	}

	fmt.Println("DONE")

	// Output:
	// CHANNEL_OUTGOING DOWN baf54d6c-ecba-4183-b2aa-0f40b0df308f CS_NONE
	// CHANNEL_STATE DOWN baf54d6c-ecba-4183-b2aa-0f40b0df308f CS_INIT
	// CHANNEL_CREATE DOWN baf54d6c-ecba-4183-b2aa-0f40b0df308f CS_INIT
	// CHANNEL_ORIGINATE DOWN baf54d6c-ecba-4183-b2aa-0f40b0df308f CS_INIT
	// CHANNEL_STATE DOWN baf54d6c-ecba-4183-b2aa-0f40b0df308f CS_ROUTING
	// CHANNEL_STATE DOWN baf54d6c-ecba-4183-b2aa-0f40b0df308f CS_CONSUME_MEDIA
	// CHANNEL_PROGRESS DOWN baf54d6c-ecba-4183-b2aa-0f40b0df308f CS_CONSUME_MEDIA
	// CHANNEL_CALLSTATE RINGING baf54d6c-ecba-4183-b2aa-0f40b0df308f CS_CONSUME_MEDIA
	// CHANNEL_ANSWER RINGING baf54d6c-ecba-4183-b2aa-0f40b0df308f CS_CONSUME_MEDIA
	// CHANNEL_CALLSTATE ACTIVE baf54d6c-ecba-4183-b2aa-0f40b0df308f CS_CONSUME_MEDIA
	// CHANNEL_OUTGOING ACTIVE baf54d6c-ecba-4183-b2aa-0f40b0df308f CS_CONSUME_MEDIA
	// CHANNEL_STATE ACTIVE baf54d6c-ecba-4183-b2aa-0f40b0df308f CS_EXCHANGE_MEDIA
	// CHANNEL_HANGUP ACTIVE baf54d6c-ecba-4183-b2aa-0f40b0df308f CS_EXCHANGE_MEDIA
	// CHANNEL_STATE ACTIVE baf54d6c-ecba-4183-b2aa-0f40b0df308f CS_HANGUP
	// CHANNEL_CALLSTATE HANGUP baf54d6c-ecba-4183-b2aa-0f40b0df308f CS_HANGUP
	// CHANNEL_STATE HANGUP baf54d6c-ecba-4183-b2aa-0f40b0df308f CS_REPORTING
	// CHANNEL_HANGUP_COMPLETE HANGUP baf54d6c-ecba-4183-b2aa-0f40b0df308f CS_REPORTING
	// CHANNEL_DESTROY HANGUP baf54d6c-ecba-4183-b2aa-0f40b0df308f CS_REPORTING
	// CHANNEL_STATE HANGUP baf54d6c-ecba-4183-b2aa-0f40b0df308f CS_DESTROY
	// DONE
}

type Event []map[string]string

var Events = Event{{
	"Channel-Call-State": "DOWN",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_NEW",
	"Event-Name":         "CHANNEL_STATE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "DOWN",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_INIT",
	"Event-Name":         "CHANNEL_STATE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "DOWN",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_INIT",
	"Event-Name":         "CHANNEL_CREATE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_ROUTING",
	"Event-Name":         "CHANNEL_CALLSTATE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_ROUTING",
	"Event-Name":         "CHANNEL_STATE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_STATE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_EXECUTE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_EXECUTE_COMPLETE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_EXECUTE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_EXECUTE_COMPLETE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_EXECUTE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_EXECUTE_COMPLETE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_EXECUTE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_EXECUTE_COMPLETE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_EXECUTE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_EXECUTE_COMPLETE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_EXECUTE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_EXECUTE_COMPLETE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_EXECUTE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_EXECUTE_COMPLETE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_EXECUTE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_EXECUTE_COMPLETE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_EXECUTE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_EXECUTE_COMPLETE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_EXECUTE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_EXECUTE_COMPLETE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_EXECUTE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_EXECUTE_COMPLETE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_EXECUTE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_EXECUTE_COMPLETE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_EXECUTE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_EXECUTE_COMPLETE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_EXECUTE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_EXECUTE_COMPLETE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_EXECUTE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_EXECUTE_COMPLETE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_EXECUTE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_EXECUTE_COMPLETE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_EXECUTE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_EXECUTE_COMPLETE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_EXECUTE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "DOWN",
	"Channel-Call-UUID":  "baf54d6c-ecba-4183-b2aa-0f40b0df308f",
	"Channel-State":      "CS_NONE",
	"Event-Name":         "CHANNEL_OUTGOING",
	"Unique-ID":          "baf54d6c-ecba-4183-b2aa-0f40b0df308f",
	"Begin":              "",
}, {
	"Channel-Call-State": "DOWN",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_INIT",
	"Event-Name":         "CHANNEL_STATE",
	"Unique-ID":          "baf54d6c-ecba-4183-b2aa-0f40b0df308f",
}, {
	"Channel-Call-State": "DOWN",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_INIT",
	"Event-Name":         "CHANNEL_CREATE",
	"Unique-ID":          "baf54d6c-ecba-4183-b2aa-0f40b0df308f",
}, {
	"Channel-Call-State": "DOWN",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_INIT",
	"Event-Name":         "CHANNEL_ORIGINATE",
	"Unique-ID":          "baf54d6c-ecba-4183-b2aa-0f40b0df308f",
}, {
	"Channel-Call-State": "DOWN",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_ROUTING",
	"Event-Name":         "CHANNEL_STATE",
	"Unique-ID":          "baf54d6c-ecba-4183-b2aa-0f40b0df308f",
}, {
	"Channel-Call-State": "DOWN",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_CONSUME_MEDIA",
	"Event-Name":         "CHANNEL_STATE",
	"Unique-ID":          "baf54d6c-ecba-4183-b2aa-0f40b0df308f",
}, {
	"Channel-Call-State": "DOWN",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_CONSUME_MEDIA",
	"Event-Name":         "CHANNEL_PROGRESS",
	"Unique-ID":          "baf54d6c-ecba-4183-b2aa-0f40b0df308f",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_CONSUME_MEDIA",
	"Event-Name":         "CHANNEL_CALLSTATE",
	"Unique-ID":          "baf54d6c-ecba-4183-b2aa-0f40b0df308f",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_PROGRESS_MEDIA",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "EARLY",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_CALLSTATE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "RINGING",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_CONSUME_MEDIA",
	"Event-Name":         "CHANNEL_ANSWER",
	"Unique-ID":          "baf54d6c-ecba-4183-b2aa-0f40b0df308f",
}, {
	"Channel-Call-State": "ACTIVE",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_CONSUME_MEDIA",
	"Event-Name":         "CHANNEL_CALLSTATE",
	"Unique-ID":          "baf54d6c-ecba-4183-b2aa-0f40b0df308f",
}, {
	"Channel-Call-State": "EARLY",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_ANSWER",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "ACTIVE",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_CALLSTATE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "ACTIVE",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_CONSUME_MEDIA",
	"Event-Name":         "CHANNEL_OUTGOING",
	"Unique-ID":          "baf54d6c-ecba-4183-b2aa-0f40b0df308f",
}, {
	"Channel-Call-State": "ACTIVE",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_BRIDGE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "ACTIVE",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXCHANGE_MEDIA",
	"Event-Name":         "CHANNEL_STATE",
	"Unique-ID":          "baf54d6c-ecba-4183-b2aa-0f40b0df308f",
}, {
	"Channel-Call-State": "ACTIVE",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_HANGUP",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "ACTIVE",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXCHANGE_MEDIA",
	"Event-Name":         "CHANNEL_HANGUP",
	"Unique-ID":          "baf54d6c-ecba-4183-b2aa-0f40b0df308f",
}, {
	"Channel-Call-State": "ACTIVE",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_HANGUP",
	"Event-Name":         "CHANNEL_STATE",
	"Unique-ID":          "baf54d6c-ecba-4183-b2aa-0f40b0df308f",
}, {
	"Channel-Call-State": "HANGUP",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_HANGUP",
	"Event-Name":         "CHANNEL_CALLSTATE",
	"Unique-ID":          "baf54d6c-ecba-4183-b2aa-0f40b0df308f",
}, {
	"Channel-Call-State": "ACTIVE",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_UNBRIDGE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "ACTIVE",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_EXECUTE",
	"Event-Name":         "CHANNEL_EXECUTE_COMPLETE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "ACTIVE",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_HANGUP",
	"Event-Name":         "CHANNEL_STATE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "HANGUP",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_HANGUP",
	"Event-Name":         "CHANNEL_CALLSTATE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "HANGUP",
	"Channel-Call-UUID":  "baf54d6c-ecba-4183-b2aa-0f40b0df308f",
	"Channel-State":      "CS_REPORTING",
	"Event-Name":         "CHANNEL_STATE",
	"Unique-ID":          "baf54d6c-ecba-4183-b2aa-0f40b0df308f",
}, {
	"Channel-Call-State": "HANGUP",
	"Channel-Call-UUID":  "baf54d6c-ecba-4183-b2aa-0f40b0df308f",
	"Channel-State":      "CS_REPORTING",
	"Event-Name":         "CHANNEL_HANGUP_COMPLETE",
	"Unique-ID":          "baf54d6c-ecba-4183-b2aa-0f40b0df308f",
}, {
	"Channel-Call-State": "HANGUP",
	"Channel-Call-UUID":  "baf54d6c-ecba-4183-b2aa-0f40b0df308f",
	"Channel-State":      "CS_REPORTING",
	"Event-Name":         "CHANNEL_DESTROY",
	"Unique-ID":          "baf54d6c-ecba-4183-b2aa-0f40b0df308f",
}, {
	"Channel-Call-State": "HANGUP",
	"Channel-Call-UUID":  "baf54d6c-ecba-4183-b2aa-0f40b0df308f",
	"Channel-State":      "CS_DESTROY",
	"Event-Name":         "CHANNEL_STATE",
	"Unique-ID":          "baf54d6c-ecba-4183-b2aa-0f40b0df308f",
}, {
	"Channel-Call-State": "HANGUP",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_REPORTING",
	"Event-Name":         "CHANNEL_STATE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "HANGUP",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_REPORTING",
	"Event-Name":         "CHANNEL_HANGUP_COMPLETE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "HANGUP",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_REPORTING",
	"Event-Name":         "CHANNEL_DESTROY",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}, {
	"Channel-Call-State": "HANGUP",
	"Channel-Call-UUID":  "5a4f1c56-e680-4335-90dc-d2f3639513ca",
	"Channel-State":      "CS_DESTROY",
	"Event-Name":         "CHANNEL_STATE",
	"Unique-ID":          "5a4f1c56-e680-4335-90dc-d2f3639513ca",
}}
