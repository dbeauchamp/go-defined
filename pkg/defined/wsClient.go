package defined

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

const (
	definedWSEndpoint = "/graphql/realtime"
	definedWSHost     = "realtime.api.defined.fi"
)

type Handler func(message []byte)

type WSClient struct {
	apiKey string
	c      *websocket.Conn
}

type SubscriptionAuth struct {
	Host          string `json:"host"`
	Authorization string `json:"Authorization"`
}

type SubscriptionPayloadExtensions struct {
	Authorization SubscriptionAuth `json:"authorization"`
}

type SubscriptionPayload struct {
	Data       string                        `json:"data"`
	Extensions SubscriptionPayloadExtensions `json:"extensions"`
}

type SubscriptionConfig struct {
	Id      string              `json:"id"`
	Payload SubscriptionPayload `json:"payload"`
	Type    string              `json:"type"`
}

type PairEventsSubscriptionArgs struct {
	Id string
}

func base64EncodedAuth(key string) string {
	authObj := fmt.Sprintf(
		"{\"host\": \"%v\", \"Authorization\": \"%v\" }",
		definedWSHost,
		key,
	)
	return base64.StdEncoding.EncodeToString([]byte(authObj))
}

func NewWSClient(apiKey string) WSClient {
    encodedAuthObj := base64EncodedAuth(apiKey)
	wsUrl := url.URL{
		Scheme:   "wss",
		Host:     definedWSHost,
		Path:     definedWSEndpoint,
		RawQuery: fmt.Sprintf("header=%v&payload=e30=", encodedAuthObj),
	}

	dialer := websocket.DefaultDialer
	dialer.Subprotocols = []string{"graphql-ws"}
	c, _, err := dialer.Dial(wsUrl.String(), nil)
	if err != nil {
		log.Fatalf("websocket dailer error: %v", err)
	}

	return WSClient{
		apiKey: apiKey,
		c:      c,
	}
}

func (ws *WSClient) getPairEventsJSON(id string) []byte {
	data := map[string]interface{}{
		"query": PairEventsSubscription,
		"variables": map[string]interface{}{
			"id": id,
		},
	}
	bytes, err := json.Marshal(data)
	if err != nil {
		log.Fatal("Could not encode pair events subscription data")
	}

	config := SubscriptionConfig{
		Type: "start",
		Id:   "testing123",
		Payload: SubscriptionPayload{
			Extensions: SubscriptionPayloadExtensions{
				Authorization: SubscriptionAuth{
					Host:          definedWSHost,
					Authorization: ws.apiKey,
				},
			},
			Data: string(bytes),
		},
	}

	subJSON, _ := json.Marshal(config)

	return subJSON
}

func (ws *WSClient) SubscribeToPairEvents(
	args PairEventsSubscriptionArgs,
	handler Handler,
) (*chan []byte, *chan bool) {
	config := ws.getPairEventsJSON(args.Id)
	err := ws.c.WriteMessage(websocket.TextMessage, config)
	if err != nil {
        // TODO the error
		fmt.Printf("Could not subscribe to pair events: %v", err)
        return nil, nil
	}

	msgCh := make(chan []byte)
    done := make(chan bool)

	go func() {
		for {
            select {
            case <-done:
                ws.c.Close()
                return
            default:
                _, msg, err := ws.c.ReadMessage()
                if err != nil {
                    fmt.Printf("Error reading ws msg: %v \n", err)
                }
                msgCh <- msg
            }
		}
	}()

	return &msgCh, &done
}
