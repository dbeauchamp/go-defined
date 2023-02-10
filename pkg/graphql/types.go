package graphql

type OnCreateEventsPayload struct {
	Data OnCreateEventsData `json:"data"`
}

type OnCreateEventsData struct {
	OnCreateEvents CreateEvents `json:"onCreateEvents"`
}

type CreateEvents struct {
	Address   *string  `json:"address,omitempty"`
	Id        *string  `json:"id,omitempty"`
	NetworkId *int     `json:"networkId,omitempty"`
	Events    *[]Event `json:"events,omitempty"`
}

type Event struct {
	Address            *string    `json:"address,omitempty"`
	BaseTokenPrice     *string    `json:"baseTokenPrice,omitempty"`
	BlockHash          *string    `json:"blockHash,omitempty"`
	BlockNumber        *int       `json:"blockNumber,omitempty"`
	Data               *EventData `json:"data,omitempty"`
	EventType          *string    `json:"eventType,omitempty"` // TODO: enum
	Id                 *string    `json:"id,omitempty"`
	LiquidityToken     *string    `json:"liquidityToken,omitempty"`
	LogIndex           *int       `json:"logIndex,omitempty"`
	Maker              *string    `json:"maker,omitempty"`
	NetworkId          *int       `json:"networkId,omitempty"`
	Timestamp          *int       `json:"timestamp,omitempty"`
	Token0SwapValueUsd *string    `json:"token0SwapValueUsd,omitempty"`
	Token1SwapValueUsd *string    `json:"token1SwapValueUsd,omitempty"`
	Token0ValueBase    *string    `json:"token0ValueBase,omitempty"`
	Token1ValueBase    *string    `json:"token1ValueBase,omitempty"`
	TransactionHash    *string    `json:"transactionHash,omitempty"`
	TransactionIndex   *int       `json:"transactionIndex,omitempty"`
	EventDisplayType   *string    `json:"eventDisplayType,omitempty"` // TODO: enum
}

type EventData struct {
	Amount0                 *string `json:"amount0,omitempty"`
	Amount1                 *string `json:"amount1,omitempty"`
	TickLower               *string `json:"tickLower,omitempty"`
	TickUpper               *string `json:"tickUpper,omitempty"`
	Type                    *string `json:"type,omitempty"` // TODO: enum
	Amount0In               *string `json:"amount0In,omitempty"`
	Amount1In               *string `json:"amount1In,omitempty"`
	Amount0Out              *string `json:"amount0Out,omitempty"`
	Amount1Out              *string `json:"amount1Out,omitempty"`
	AmountNonLiquidityToken *string `json:"amountNonLiquidityToken,omitempty"`
	PriceBaseToken          *string `json:"priceBaseToken,omitempty"`
	PriceUsd                *string `json:"priceUsd,omitempty"`
	PriceUsdTotal           *string `json:"priceUsdTotal,omitempty"`
	Tick                    *string `json:"tick,omitempty"`
}

type OnUpdateAggregateBatchPayload struct {
	Data OnUpdateAggregateBatchData `json:"data"`
}

type OnUpdateAggregateBatchData struct {
	OnUpdateAggregateBatch AggregateBatchUpdate `json:"onCreateEvents"`
}

type AggregateBatchUpdate struct {
	PairAddress  *string              `json:"pairAddress,omitempty"`
	NetworkId    *int                 `json:"networkId,omitempty"`
	Timestamp    *int                 `json:"timestamp,omitempty"`
	PairId       *string              `json:"pairId,omitempty"`
	EventSortKey *string              `json:"eventSortKey,omitempty"`
	Aggregates   *[]ResolutionBarData `json:"aggregates,omitempty"`
}

type ResolutionBarData struct {
	R1   *CurrencyBarData `json:"r1,omitempty"`
	R5   *CurrencyBarData `json:"r5,omitempty"`
	R15  *CurrencyBarData `json:"r15,omitempty"`
	R30  *CurrencyBarData `json:"r30,omitempty"`
	R60  *CurrencyBarData `json:"r60,omitempty"`
	R240 *CurrencyBarData `json:"r240,omitempty"`
	R720 *CurrencyBarData `json:"r720,omitempty"`
	R1D  *CurrencyBarData `json:"r1D,omitempty"`
	R7D  *CurrencyBarData `json:"r7D,omitempty"`
}

type CurrencyBarData struct {
	T     *int               `json:"t,omitempty"`
	Usd   *IndividualBarData `json:"usd,omitempty"`
	Token *IndividualBarData `json:"token,omitempty"`
}

type IndividualBarData struct {
	O *[]float64 `json:"o,omitempty"`
	H *[]float64 `json:"h,omitempty"`
	L *[]float64 `json:"l,omitempty"`
	C *[]float64 `json:"c,omitempty"`
	V *[]int     `json:"v,omitempty"`
	T *[]int     `json:"t,omitempty"`
	S *string    `json:"s,omitempty"`
}

type OnUpdatePricePayload struct {
	Data OnUpdatePriceData `json:"data"`
}

type OnUpdatePriceData struct {
	OnUpdatePrice PriceUpdate `json:"onUpdatePrice"`
}

type PriceUpdate struct {
	Address   *string  `json:"address,omitempty"`
	NetworkId *int     `json:"networkId,omitempty"`
	Timestamp *int     `json:"timestamp,omitempty"`
	PriceUsd  *float64 `json:"priceUsd,omitempty"`
}

type OnCreateNFTEventsPayload struct {
	Data OnCreateNFTEventsData `json:"data"`
}

type OnCreateNFTEventsData struct {
	OnCreateNFTEvents CreateNFTEvents `json:"onCreateNFTEvents"`
}

type CreateNFTEvents struct {
	Address   *string     `json:"address,omitempty"`
	Id        *string     `json:"id,omitempty"`
	NetworkId *int        `json:"networkId,omitempty"`
	Events    *[]NFTEvent `json:"events,omitempty"`
}

type NFTEvent struct {
	Id                         *string       `json:"id:omitempty"`
	ContractAddress            *string       `json:"contractAddress:omitempty"`
	NetworkId                  *int          `json:"networkId:omitempty"`
	TokenId                    *string       `json:"tokenId:omitempty"`
	AggregatorAddress          *string       `json:"aggregatorAddress :omitempty"`
	FillSource                 *string       `json:"fillSource:omitempty"`
	Maker                      *string       `json:"maker:omitempty"`
	Taker                      *string       `json:"taker:omitempty"`
	TotalTradePrice            *string       `json:"totalTradePrice:omitempty"`
	TotalPriceUsd              *string       `json:"totalPriceUsd:omitempty"`
	TotalPriceNetworkBaseToken *string       `json:"totalPriceNetworkBaseToken:omitempty"`
	PaymentTokenAddress        *string       `json:"paymentTokenAddress:omitempty"`
	EventType                  *string       `json:"eventType:omitempty"`
	Data                       *NFTEventData `json:"data:omitempty"`
	ExchangeAddress            *string       `json:"exchangeAddress:omitempty"`
	PoolAddress                *string       `json:"poolAddress:omitempty"`
	SortKey                    *string       `json:"sortKey:omitempty"`
	BlockNumber                *int          `json:"blockNumber,omitempty"`
	TransactionHash            *string       `json:"transactionHash,omitempty"`
	TransactionIndex           *int          `json:"transactionIndex,omitempty"`
	LogIndex                   *int          `json:"logIndex,omitempty"`
	Timestamp                  *int          `json:"timestamp,omitempty"`
	NumberOfTokens             *string       `json:"numberOfTokens,omitempty"`
	PriceError                 *string       `json:"priceError,omitempty"`
}

type NFTEventData struct {
	BuyHash  *string `json:"buyHash,omitempty"`
	Metadata *string `json:"metadata,omitempty"`
	Price    *string `json:"price,omitempty"`
	Maker    *string `json:"maker,omitempty"`
	Taker    *string `json:"taker,omitempty"`
	Type     *string `json:"type,omitempty"`
	SellHash *string `json:"sellHash,omitempty"`
}
