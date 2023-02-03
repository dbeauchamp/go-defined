package main

import (
	"fmt"

	"github.com/dbeauchamp/go-defined/pkg/defined"
)

func main() {
	fmt.Printf("... starting \n")
	ws := defined.NewWSClient("")
	noop := func(msg []byte) {}

	fmt.Printf("... init subscription \n")
	msgCh, _ := ws.SubscribeToPairEvents(
		defined.PairEventsSubscriptionArgs{
			Id: "0x0ed7e52944161450477ee417de9cd3a859b14fd0:56",
		},
		noop,
	)

	for {
		select {
		case msg := <-*msgCh:
			fmt.Printf("recieved \n")
			fmt.Printf(string(msg))
		}

	}
}
