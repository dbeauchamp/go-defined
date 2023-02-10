package graphql

const PairEventsSubscription = `
    subscription CreateEvents($id: String) {
      onCreateEvents(id: $id) {
        events {
          address
          baseTokenPrice
          blockHash
          blockNumber
          eventDisplayType
          eventType
          logIndex
          id
          liquidityToken
          maker
          networkId
          timestamp
          token0SwapValueUsd
          token0ValueBase
          token1SwapValueUsd
          token1ValueBase
          transactionHash
          transactionIndex
          data {
            ... on MintEventData {
              amount0
              amount1
              amount0Shifted
              amount1Shifted
              tickLower
              tickUpper
              type
            }
            ... on BurnEventData {
              amount0
              amount1
              amount0Shifted
              amount1Shifted
              tickLower
              tickUpper
              type
            }
            ... on SwapEventData {
              amount0
              amount0In
              amount0Out
              amount1
              amount1In
              amount1Out
              amountNonLiquidityToken
              priceBaseToken
              priceBaseTokenTotal
              priceUsd
              priceUsdTotal
              tick
              type
            }
          }
        }
        address
        id
        networkId
      }
    }
`

const SubscribeToAggregates = `
	subscription UpdateAggregateBatch($pairId: String) {
	  onUpdateAggregateBatch(pairId: $pairId) {
		eventSortKey
		networkId
		pairAddress
		pairId
		timestamp
		aggregates {
		r1 {
		  t
		  usd {
			t
			o
			h
			l
			c
			volume
		  }
		  token {
			t
			o
			h
			l
			c
			volume
		  }
		}
		r5 {
		  t
		  usd {
			t
			o
			h
			l
			c
			volume
		  }
		  token {
			t
			o
			h
			l
			c
			volume
		  }
		}
		r15 {
		  t
		  usd {
			t
			o
			h
			l
			c
			volume
		  }
		  token {
			t
			o
			h
			l
			c
			volume
		  }
		}
		r30 {
		  t
		  usd {
			t
			o
			h
			l
			c
			volume
		  }
		  token {
			t
			o
			h
			l
			c
			volume
		  }
		}
		r60 {
		  t
		  usd {
			t
			o
			h
			l
			c
			volume
		  }
		  token {
			t
			o
			h
			l
			c
			volume
		  }
		}
		r240 {
		  t
		  usd {
			t
			o
			h
			l
			c
			volume
		  }
		  token {
			t
			o
			h
			l
			c
			volume
		  }
		}
		r720 {
		  t
		  usd {
			t
			o
			h
			l
			c
			volume
		  }
		  token {
			t
			o
			h
			l
			c
			volume
		  }
		}
		r1D {
		  t
		  usd {
			t
			o
			h
			l
			c
			volume
		  }
		  token {
			t
			o
			h
			l
			c
			volume
		  }
		}
		r7D {
		  t
		  usd {
			t
			o
			h
			l
			c
			volume
		  }
		  token {
			t
			o
			h
			l
			c
			volume
		  }
		}
	  }

`

const PriceUpdateSubscription = `
	subscription UpdatePrice($address: String, $networkId: Int) {
	  onUpdatePrice(address: $address, networkId: $networkId) {
		address
		networkId
		priceUsd
		timestamp
	  }
	}
`
const NFTEventsSubscription = `
	subscription CreateNftEvents($address: String, $networkId: Int) {
	  onCreateNftEvents(address: $address, networkId: $networkId) {
		address
		id
		networkId
		events {
		  blockNumber
		  contractAddress
		  data {
			buyHash
			metadata
			price
			maker
			taker
			type
			sellHash
		  }
		  eventType
		  exchangeAddress
		  fillSource
		  id
		  individualPriceNetworkBaseToken
		  individualPriceUsd
		  individualTradePrice
		  logIndex
		  maker
		  networkId
		  numberOfTokens
		  paymentTokenAddress
		  poolAddress
		  priceError
		  sortKey
		  taker
		  timestamp
		  tokenId
		  totalPriceNetworkBaseToken
		  totalPriceUsd
		  totalTradePrice
		  transactionHash
		  transactionIndex
		}
	  }
	}
`
