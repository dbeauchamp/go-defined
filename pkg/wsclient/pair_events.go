package wsclient

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"

	. "github.com/dbeauchamp/go-defined/pkg/graphql"
)

func (ws *WSClient) buildPairEventsJSON(opts PairEventsSubscriptionArgs) ([]byte, error) {
	query := PairEventsSubscription
	if opts.SubscriptionOptions.Query != "" {
		query = opts.Query
	}

	data := map[string]interface{}{
		"query": query,
		"variables": map[string]interface{}{
			"id": opts.Id,
		},
	}

	subscriptionId := opts.Id
	if opts.SubscriptionOptions.SubscriptionId != "" {
		subscriptionId = opts.SubscriptionOptions.SubscriptionId
	}

	subJSON, err := ws.buildSubscriptionData(data, subscriptionId)
	if err != nil {
		return nil, err
	}

	return subJSON, nil
}

func (ws *WSClient) SubscribeToPairEvents(
	args PairEventsSubscriptionArgs,
) (*chan *Event, *chan bool, error) {
	config, err := ws.buildPairEventsJSON(args)
	if err != nil {
		return nil, nil, err
	}
	err = ws.c.WriteMessage(websocket.TextMessage, config)
	if err != nil {
		return nil, nil, err
	}

	msgCh := make(chan *Event)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				ws.c.Close()
				return
			default:
				wsMsg := ws.readMessage()
				if wsMsg.Type == "data" {
					var payload OnCreateEventsPayload
					err = json.Unmarshal(wsMsg.Payload, &payload)
					if err != nil {
						log.Printf("Error unmarshalling wsMsg: %v \n", err)
					}
					events := payload.Data.OnCreateEvents.Events

					for _, e := range *events {
						msgCh <- &e
					}
				}
			}
		}
	}()

	return &msgCh, &done, nil
}
