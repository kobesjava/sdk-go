package iwebsocket

type WsClient interface {
	SubscribeKline(handler func(string)) (err error)
}
