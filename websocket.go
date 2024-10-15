package hyperliquid

import "golang.org/x/net/websocket"

// WSClient websocket client for hyperliquid
type WSClient struct {
	ws *websocket.Conn
}

// NewWSClient return new instance of WSClient
func NewWSClient() *WSClient {
	return &WSClient{}
}
