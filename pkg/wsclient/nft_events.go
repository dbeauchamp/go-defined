package wsclient

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"

	. "github.com/dbeauchamp/go-defined/pkg/graphql"
)

func (ws *WSClient) buildNFTEventsJSON(opts NFTEventsSubscriptionArgs) ([]byte, error) {
	query := opts.Query
	if len(opts.Query) == 0 {
		query = NFTEventsSubscription
	}

	data := map[string]any{
		"query": query,
		"variables": map[string]any{
			"address":   opts.Address,
			"networkId": opts.NetworkId,
		},
	}

	subJSON, err := ws.buildSubscriptionData(data, opts.SubscriptionId)
	if err != nil {
		return nil, err
	}

	return subJSON, nil
}

func (ws *WSClient) SubscribeToNFTEvents(
	args NFTEventsSubscriptionArgs,
) (*chan *NFTEvent, *chan struct{}, error) {
	subscriptionId := args.SubscriptionId
	if len(args.SubscriptionId) == 0 {
		subscriptionId = fmt.Sprintf("%v:%v", args.Address, args.NetworkId)
		args.SubscriptionId = subscriptionId
	}

	config, err := ws.buildNFTEventsJSON(args)
	if err != nil {
		return nil, nil, err
	}
	err = ws.c.WriteMessage(websocket.TextMessage, config)
	if err != nil {
		return nil, nil, err
	}

	msgCh := make(chan *NFTEvent)
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
					var payload OnCreateNFTEventsPayload
					err = json.Unmarshal(msg.Payload, &payload)
					if err != nil {
						log.Printf("Error unmarshalling wsMsg: %v \n", err)
					}
					events := payload.Data.OnCreateNFTEvents.Events

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
