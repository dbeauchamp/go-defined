package wsclient

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"

	. "github.com/dbeauchamp/go-defined/pkg/graphql"
)

func (ws *WSClient) buildNFTEventsJSON(opts NFTEventsSubscriptionArgs) ([]byte, error) {
	query := NFTEventsSubscription
	if opts.SubscriptionOptions.Query != "" {
		query = opts.Query
	}

	data := map[string]interface{}{
		"query": query,
		"variables": map[string]interface{}{
			"address":   opts.Address,
			"networkId": opts.NetworkId,
		},
	}

	subscriptionId := fmt.Sprintf("%v:%v", opts.Address, opts.NetworkId)
	if opts.SubscriptionOptions.SubscriptionId != "" {
		subscriptionId = opts.SubscriptionOptions.SubscriptionId
	}

	subJSON, err := ws.buildSubscriptionData(data, subscriptionId)
	if err != nil {
		return nil, err
	}

	return subJSON, nil
}

func (ws *WSClient) SubscribeToNFTEvents(
	args NFTEventsSubscriptionArgs,
) (*chan *NFTEvent, *chan bool, error) {
	config, err := ws.buildNFTEventsJSON(args)
	if err != nil {
		return nil, nil, err
	}
	err = ws.c.WriteMessage(websocket.TextMessage, config)
	if err != nil {
		return nil, nil, err
	}

	msgCh := make(chan *NFTEvent)
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
					var payload OnCreateNFTEventsPayload
					err = json.Unmarshal(wsMsg.Payload, &payload)
					if err != nil {
						log.Printf("Error unmarshalling wsMsg: %v \n", err)
					}
					events := payload.Data.OnCreateNFTEvents.Events

					for _, e := range *events {
						msgCh <- &e
					}
				}
			}
		}
	}()

	return &msgCh, &done, nil
}
