# Go Defined.fi API Client

A client library for accessing [Defined.fi's](http://www.defined.fi) realtime graphql api

## Installation

```bash
go get -u github.com/dbeauchamp/go-defined
```

## Usage
You'll need a Defined API Key. See instructions here: [defined docs](https://docs.defined.fi/)

```go
package main

import (
	"fmt"
	"os"

	"github.com/dbeauchamp/go-defined/pkg/wsclient"
)

func main() {
	// Your api key
	ws := wsclient.New(os.Getenv("DEFINED_API_KEY"))

	// Subscribe to events
	msgCh, doneCh, err := ws.SubscribeToPairEvents(
		wsclient.PairEventsSubscriptionArgs{
			Id: "0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640:1", // WETH/USDC
		},
	)
	if err != nil {
		...
	}

	// Do something fun!
	for msg := range *msgCh {
		fmt.Println(*msg.Maker)
	}

	// Unsubscribe
	*doneCh <- struct{}{}

	// Close connection
	ws.Close()
}
```

## License
MIT

## Contributing
Contributions are welcome!
