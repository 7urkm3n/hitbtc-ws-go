### Go Microservice

The app is listens HitBTC api via sockets and caches last update in the memory.

The app interfaces for the rest client only for the two main coins ["BTC", "ETH"].

######Available endpoints:
`/all` and `/{Symbol}`
symbols: [btc, eth]

######To run server:
`go run src/main.go`

######Dependencies:
`go-chi` router and golang `net` package.
