package wsclient

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"

	. "github.com/dbeauchamp/go-defined/pkg/graphql"
	"github.com/dbeauchamp/go-defined/pkg/logger"
)

func (ws *WSClient) buildPriceUpdateJSON(opts PriceUpdateSubscriptionArgs) ([]byte, error) {
	query := SubscribeToAggregates
	if opts.SubscriptionOptions.Query != "" {
		query = opts.Query
	}
	data := map[string]any{
		"query": query,
		"variables": map[string]any{
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
) (*chan *PriceUpdate, *chan struct{}, error) {
	subscriptionId := args.SubscriptionId
	if len(args.SubscriptionId) == 0 {
		subscriptionId = fmt.Sprintf("%v:%v", args.Address, args.NetworkId)
		args.SubscriptionId = subscriptionId
	}

	config, err := ws.buildPriceUpdateJSON(args)
	if err != nil {
		return nil, nil, err
	}
	err = ws.c.WriteMessage(websocket.TextMessage, config)
	if err != nil {
		return nil, nil, err
	}

	msgCh := make(chan *PriceUpdate)
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
					var payload OnUpdatePricePayload
					err = json.Unmarshal(msg.Payload, &payload)
					if err != nil {
						logger.Error("Error unmarshalling wsMsg", err)
					}

					priceUpdate := payload.Data.OnUpdatePrice
					msgCh <- &priceUpdate
				}
			default:
			}
		}
	}()

	return &msgCh, &done, nil
}
