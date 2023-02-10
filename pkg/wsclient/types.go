package wsclient

import "encoding/json"

type WSMsg struct {
	Id      string          `json:"id"`
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
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

type SubscriptionOptions struct {
	SubscriptionId string
	Query          string
}

type PairEventsSubscriptionArgs struct {
	Id string
	SubscriptionOptions
}

type AggregateSubscriptionArgs struct {
	Id string
	SubscriptionOptions
}

type PriceUpdateSubscriptionArgs struct {
	Address   string
	NetworkId int
	SubscriptionOptions
}

type NFTEventsSubscriptionArgs struct {
	Address   string
	NetworkId int
	SubscriptionOptions
}
