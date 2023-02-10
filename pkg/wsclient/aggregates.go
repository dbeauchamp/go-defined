package wsclient

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"

	. "github.com/dbeauchamp/go-defined/pkg/graphql"
)

func (ws *WSClient) buildAggregatesJSON(opts AggregateSubscriptionArgs) ([]byte, error) {
	query := SubscribeToAggregates
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

func (ws *WSClient) SubscribeToAggregates(
	args AggregateSubscriptionArgs,
) (*chan *AggregateBatchUpdate, *chan bool, error) {
	config, err := ws.buildAggregatesJSON(args)
	if err != nil {
		return nil, nil, err
	}
	err = ws.c.WriteMessage(websocket.TextMessage, config)
	if err != nil {
		return nil, nil, err
	}

	msgCh := make(chan *AggregateBatchUpdate)
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
					var payload OnUpdateAggregateBatchPayload
					err = json.Unmarshal(wsMsg.Payload, &payload)
					if err != nil {
						log.Printf("Error unmarshalling wsMsg: %v \n", err)
					}

					aggBatch := payload.Data.OnUpdateAggregateBatch
					msgCh <- &aggBatch
				}
			}
		}
	}()

	return &msgCh, &done, nil
}
