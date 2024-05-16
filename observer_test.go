package observer

import (
	"context"
	"fmt"
	"github.com/mdigger/esl"
	"sync"
)

func ExampleNew() {
	ctx := context.Background()

	events := make(chan esl.Event, 1)

	o := New()

	go o.Run(ctx, events)

	var wg1, wg2 sync.WaitGroup

	wg1.Add(2)
	wg2.Add(1)

	go func() {
		for i, e := range Events {
			events <- esl.NewEvent(e.Name, e.Headers, []byte{})
			if i < 2 {
				wg1.Done()
			} else if i == 2 {
				wg2.Wait()
			}
		}
		close(events)
	}()

	wg1.Wait()

	watch := make(chan []esl.Event, 1)

	o.Watch(watch, "5a4f1c56-e680-4335-90dc-d2f3639513ca")

	wg2.Done()

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
	// CHANNEL_STATE DOWN 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_NEW
	// CHANNEL_STATE DOWN 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_INIT
	// CHANNEL_CREATE DOWN 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_INIT
	// CHANNEL_CALLSTATE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_ROUTING
	// CHANNEL_STATE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_ROUTING
	// CHANNEL_STATE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_EXECUTE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_EXECUTE_COMPLETE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_EXECUTE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_EXECUTE_COMPLETE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_EXECUTE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_EXECUTE_COMPLETE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_EXECUTE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_EXECUTE_COMPLETE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_EXECUTE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_EXECUTE_COMPLETE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_EXECUTE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_EXECUTE_COMPLETE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_EXECUTE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_EXECUTE_COMPLETE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_EXECUTE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_EXECUTE_COMPLETE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_EXECUTE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_EXECUTE_COMPLETE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_EXECUTE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_EXECUTE_COMPLETE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_EXECUTE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_EXECUTE_COMPLETE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_EXECUTE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_EXECUTE_COMPLETE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_EXECUTE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_EXECUTE_COMPLETE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_EXECUTE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_EXECUTE_COMPLETE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_EXECUTE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_EXECUTE_COMPLETE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_EXECUTE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_EXECUTE_COMPLETE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_EXECUTE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_EXECUTE_COMPLETE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_EXECUTE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_STATE DOWN 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_INIT
	// CHANNEL_CREATE DOWN 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_INIT
	// CHANNEL_ORIGINATE DOWN 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_INIT
	// CHANNEL_STATE DOWN 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_ROUTING
	// CHANNEL_STATE DOWN 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_CONSUME_MEDIA
	// CHANNEL_PROGRESS DOWN 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_CONSUME_MEDIA
	// CHANNEL_CALLSTATE RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_CONSUME_MEDIA
	// CHANNEL_PROGRESS_MEDIA RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_CALLSTATE EARLY 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_ANSWER RINGING 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_CONSUME_MEDIA
	// CHANNEL_CALLSTATE ACTIVE 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_CONSUME_MEDIA
	// CHANNEL_ANSWER EARLY 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_CALLSTATE ACTIVE 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_OUTGOING ACTIVE 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_CONSUME_MEDIA
	// CHANNEL_BRIDGE ACTIVE 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_STATE ACTIVE 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXCHANGE_MEDIA
	// CHANNEL_HANGUP ACTIVE 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_HANGUP ACTIVE 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXCHANGE_MEDIA
	// CHANNEL_STATE ACTIVE 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_HANGUP
	// CHANNEL_CALLSTATE HANGUP 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_HANGUP
	// CHANNEL_UNBRIDGE ACTIVE 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_EXECUTE_COMPLETE ACTIVE 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_EXECUTE
	// CHANNEL_STATE ACTIVE 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_HANGUP
	// CHANNEL_CALLSTATE HANGUP 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_HANGUP
	// CHANNEL_STATE HANGUP 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_REPORTING
	// CHANNEL_HANGUP_COMPLETE HANGUP 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_REPORTING
	// CHANNEL_DESTROY HANGUP 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_REPORTING
	// CHANNEL_STATE HANGUP 5a4f1c56-e680-4335-90dc-d2f3639513ca CS_DESTROY
	// DONE
}

type Event struct {
	Name    string
	Headers map[string]string
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
