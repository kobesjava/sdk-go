package websocket

type WsClient interface {
	SubscribeKline(handler func(string)) (err error)
}
