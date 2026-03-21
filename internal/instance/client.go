package instance

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"ninja_v1/internal/game"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

const WriteWait = 10 * time.Second
const PongWait = 60 * time.Second
const PingPeriod = (PongWait * 9) / 10
const MaxMessageSize = 512

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // TODO
	},
}

type Client struct {
	ID       uuid.UUID
	conn     *websocket.Conn
	ctx      context.Context
	cancel   context.CancelFunc
	instance *Instance

	nextState   chan game.Game
	nextClients chan []*Client
}

func NewClient(instance *Instance) *Client {
	ctx, cancel := context.WithCancel(instance.ctx)
	return &Client{
		ID:       uuid.New(),
		ctx:      ctx,
		cancel:   cancel,
		instance: instance,

		nextState:   make(chan game.Game, 5),
		nextClients: make(chan []*Client, 5),
	}
}

func (c *Client) Connect(w http.ResponseWriter, r *http.Request) error {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}

	c.conn = conn
	c.instance.Register <- c
	return nil
}

func (c *Client) WriteState(state game.Game) error {
	json, err := json.Marshal(NewStateMessage(state))
	if err != nil {
		return err
	}
	if err = c.conn.WriteMessage(websocket.TextMessage, json); err != nil {
		return err
	}

	return nil
}

func (c *Client) WriteClients(clients []*Client) error {
	json, err := json.Marshal(NewClientsMessage(clients))
	if err != nil {
		return err
	}

	if err = c.conn.WriteMessage(websocket.TextMessage, json); err != nil {
		return err
	}

	return nil
}

func (c *Client) listenForRequest(request *Request) error {
	_, raw, err := c.conn.ReadMessage()
	if err != nil {
		return err
	}
	if err := json.Unmarshal(raw, request); err != nil {
		return err
	}
	return nil
}

func (c *Client) listenIn() {
	defer func() {
		c.instance.Unregister <- c
		c.conn.Close()
		c.cancel()
	}()

	pongHandler := func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(PongWait))
		return nil
	}
	c.conn.SetReadLimit(MaxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(PongWait))
	c.conn.SetPongHandler(pongHandler)

	for {
		var request Request
		if err := c.listenForRequest(&request); err != nil {
			// if this error is an expected close error
			// or a message format error,
			//    then we can close the client
			break
		}

		select {
		case c.instance.ReadRequest <- request:
		case <-c.ctx.Done():
			return
		}
	}
}

func (c *Client) listenOut() {
	clock := time.NewTicker(PingPeriod)
	defer func() {
		clock.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case state := <-c.nextState:
			c.conn.SetWriteDeadline(time.Now().Add(WriteWait))
			if err := c.WriteState(state); err != nil {
				return
			}
		case clients := <-c.nextClients:
			c.conn.SetWriteDeadline(time.Now().Add(WriteWait))
			if err := c.WriteClients(clients); err != nil {
				return
			}
		// this block ensures that the client doesnt' get disconnected
		// automatically when the connection is idle
		case <-clock.C:
			c.conn.SetWriteDeadline(time.Now().Add(WriteWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		case <-c.ctx.Done():
			return
		}
	}
}

func (c *Client) Run() {
	go c.listenIn()
	go c.listenOut()
}
