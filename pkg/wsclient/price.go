package wsclient

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"

	. "github.com/dbeauchamp/go-defined/pkg/graphql"
)

func (ws *WSClient) buildPriceUpdateJSON(opts PriceUpdateSubscriptionArgs) ([]byte, error) {
	query := SubscribeToAggregates
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

func (ws *WSClient) SubscribeToPriceUpdates(
	args PriceUpdateSubscriptionArgs,
) (*chan *PriceUpdate, *chan bool, error) {
	config, err := ws.buildPriceUpdateJSON(args)
	if err != nil {
		return nil, nil, err
	}
	err = ws.c.WriteMessage(websocket.TextMessage, config)
	if err != nil {
		return nil, nil, err
	}

	msgCh := make(chan *PriceUpdate)
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
					var payload OnUpdatePricePayload
					err = json.Unmarshal(wsMsg.Payload, &payload)
					if err != nil {
						log.Printf("Error unmarshalling wsMsg: %v \n", err)
					}

					priceUpdate := payload.Data.OnUpdatePrice
					msgCh <- &priceUpdate
				}
			}
		}
	}()

	return &msgCh, &done, nil
}
