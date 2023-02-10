package wsclient

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

type WSClient struct {
	apiKey string
	c      *websocket.Conn
}

func base64EncodedAuth(key string) string {
	authObj := fmt.Sprintf(
		"{\"host\": \"%v\", \"Authorization\": \"%v\" }",
		definedWSHost,
		key,
	)
	return base64.StdEncoding.EncodeToString([]byte(authObj))
}

func New(apiKey string) WSClient {
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

func (ws *WSClient) buildSubscriptionData(
	data map[string]interface{},
	id string,
) ([]byte, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	config := SubscriptionConfig{
		Type: "start",
		Id:   id,
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

	subData, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	return subData, nil
}

func (ws *WSClient) readMessage() WSMsg {
	_, msg, err := ws.c.ReadMessage()
	if err != nil {
		log.Println("Error reading ws msg: %v \n", err)
	}

	var wsMsg WSMsg
	err = json.Unmarshal(msg, &wsMsg)
	if err != nil {
		log.Println("Error unmarshalling ws msg: %v \n", err)
	}

	return wsMsg
}

