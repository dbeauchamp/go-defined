package defined

import (
	"net/http"
)

type Client struct {
	httpClient *http.Client
}

type ClientConfig struct {
}

func NewClient() Client {
	return Client{}
}
