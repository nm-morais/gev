package main

import (
	"github.com/nm-morais/gev"
	"github.com/nm-morais/gev/plugins/websocket"
	"github.com/nm-morais/gev/plugins/websocket/ws"
)

// NewWebSocketServer 创建 WebSocket Server
func NewWebSocketServer(handler websocket.WSHandler, u *ws.Upgrader, opts ...gev.Option) (server *gev.Server, err error) {
	opts = append(opts, gev.Protocol(websocket.New(u)))
	return gev.NewServer(websocket.NewHandlerWrap(u, handler), opts...)
}
