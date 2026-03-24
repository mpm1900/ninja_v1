package instance

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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
	ID       uuid.UUID          `json:"ID"`
	conn     *websocket.Conn    `json:"-"`
	ctx      context.Context    `json:"-"`
	cancel   context.CancelFunc `json:"-"`
	instance *Instance          `json:"-"`
	inbox    chan Response      `json:"-"`
}

func NewClient(instance *Instance) *Client {
	ctx, cancel := context.WithCancel(instance.ctx)
	return &Client{
		ID:       uuid.New(),
		ctx:      ctx,
		cancel:   cancel,
		instance: instance,

		//nextState:   make(chan game.Game, 5),
		//nextClients: make(chan []*Client, 5),
		inbox: make(chan Response, 5),
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

func (c *Client) WriteResponse(response *Response) error {
	json, err := json.Marshal(response)
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
			fmt.Println(err)
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
		case response := <-c.inbox:
			c.conn.SetWriteDeadline(time.Now().Add(WriteWait))
			if err := c.WriteResponse(&response); err != nil {
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
