package wsclient

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/gorilla/websocket"

	"github.com/dbeauchamp/go-defined/pkg/logger"
)

const (
	definedWSEndpoint = "/graphql/realtime"
	definedWSHost     = "realtime.api.defined.fi"
)

type WSClient struct {
	apiKey    string
	c         *websocket.Conn
	publisher *chan WSMsg
}

func base64EncodedAuth(key string) string {
	authObj, _ := json.Marshal(map[string]any{
		"host": definedWSHost,
		"Authorization": key,
	})
	return base64.StdEncoding.EncodeToString(authObj)
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
		logger.Fatal("websocket dailer error", err)
	}

	publisher := make(chan WSMsg)
	client := WSClient{
		apiKey:    apiKey,
		c:         c,
		publisher: &publisher,
	}
	defer client.listen()

	return client
}

func (ws *WSClient) listen() {
	go func() {
		for {
			wsMsg := ws.readMessage()
			*ws.publisher <- wsMsg
		}
	}()
}

func (ws *WSClient) unsubscribe(id string) {
	stop := map[string]any{
		"id": id,
		"type": "stop",
	}
	msg, _ := json.Marshal(stop)
	ws.c.WriteMessage(websocket.TextMessage, msg)
}

func (ws *WSClient) Close() error {
	err := ws.c.Close()
	if err != nil {
		return err
	}
	return nil
}

func (ws *WSClient) buildSubscriptionData(
	data map[string]any,
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
		logger.Error("Error reading ws msg", err)
	}

	var wsMsg WSMsg
	err = json.Unmarshal(msg, &wsMsg)
	if err != nil {
		logger.Error("Error unmarshalling ws msg", err)
	}

	return wsMsg
}
