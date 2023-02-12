package wsclient

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"

	. "github.com/dbeauchamp/go-defined/pkg/graphql"
)

func (ws *WSClient) buildPairEventsJSON(opts PairEventsSubscriptionArgs) ([]byte, error) {
	query := opts.Query
	if len(opts.Query) == 0 {
		query = PairEventsSubscription
	}

	data := map[string]any{
		"query": query,
		"variables": map[string]any{
			"id": opts.Id,
		},
	}

	subJSON, err := ws.buildSubscriptionData(data, opts.SubscriptionId)
	if err != nil {
		return nil, err
	}

	return subJSON, nil
}

func (ws *WSClient) SubscribeToPairEvents(
	args PairEventsSubscriptionArgs,
) (*chan *Event, *chan struct{}, error) {
	subscriptionId := args.SubscriptionId
	if len(subscriptionId) == 0 {
		subscriptionId = args.Id
		args.SubscriptionId = subscriptionId
	}

	config, err := ws.buildPairEventsJSON(args)
	if err != nil {
		return nil, nil, err
	}
	err = ws.c.WriteMessage(websocket.TextMessage, config)
	if err != nil {
		return nil, nil, err
	}

	msgCh := make(chan *Event)
	done := make(chan struct{})

	go func() {
		for {
			var msg WSMsg
			select {
			case <-done:
				ws.unsubscribe(subscriptionId)
				return
			case msg = <-*ws.publisher:
				if msg.Type == "data" && msg.Id == subscriptionId {
					var payload OnCreateEventsPayload
					err = json.Unmarshal(msg.Payload, &payload)
					if err != nil {
						log.Printf("Error unmarshalling wsMsg: %v \n", err)
					}
					events := payload.Data.OnCreateEvents.Events

					for _, e := range *events {
						msgCh <- &e
					}
				}
			default:
			}
		}
	}()

	return &msgCh, &done, nil
}
