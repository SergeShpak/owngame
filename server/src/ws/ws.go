package ws

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

const wsReadBufferSize = 1024
const wsWriteBufferSize = 1024

type Client struct {
	conn *websocket.Conn
}

func UpgradeConnection(w http.ResponseWriter, req *http.Request, responseHeader http.Header) (*Client, error) {
	conn, err := websocket.Upgrade(w, req, responseHeader, wsReadBufferSize, wsWriteBufferSize)
	if err != nil {
		return nil, fmt.Errorf("an error occurred during connection upgrade: %v", err)
	}
	client := &Client{
		conn: conn,
	}
	return client, nil
}

func (c *Client) WriteMsg(msg []byte) error {
	return c.conn.WriteMessage(websocket.TextMessage, msg)
}

func (c *Client) Close() error {
	return c.conn.Close()
}
