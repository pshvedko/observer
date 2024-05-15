package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/mdigger/esl"
	"github.com/pshvedko/observer/internal/observer"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	events := make(chan esl.Event, 1)
	conn, err := esl.Connect("10.10.61.92", "ClueCon", esl.WithEvents(events))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	err = conn.Subscribe(
		"CHANNEL_ANSWER",
		"CHANNEL_BRIDGE",
		"CHANNEL_CALLSTATE",
		"CHANNEL_CREATE",
		"CHANNEL_DESTROY",
		"CHANNEL_EXECUTE",
		"CHANNEL_EXECUTE_COMPLETE",
		"CHANNEL_HANGUP",
		"CHANNEL_HANGUP_COMPLETE",
		"CHANNEL_ORIGINATE",
		"CHANNEL_OUTGOING",
		"CHANNEL_PROGRESS",
		"CHANNEL_PROGRESS_MEDIA",
		"CHANNEL_STATE",
		"CHANNEL_UNBRIDGE",
	)
	if err != nil {
		log.Fatal(err)
	}

	o := observer.New()

	go o.Run(ctx, events)

	<-ctx.Done()
}
