package models

import (
	"go.uber.org/zap"
	"ketangpai/dao/redis"
	"log"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

const (
	KeyBroadcast = "broadcast"
	KeyEager     = "eager"
	KeyAnswer    = "answer"
	KeyChoose    = "choose"
)

type Message struct {
	Msg      []byte
	LessonID string
	UserID   int64
	Username string
	Conn     *Connection
}

type Connection struct {
	WsConn *websocket.Conn
	Send   chan []byte
}

func (m Message) Read(channel string) {
	c := m.Conn

	defer func() {
		H.Quit <- m
		c.WsConn.Close()
	}()

	c.WsConn.SetReadLimit(512)
	//读写超时设置为60s
	c.WsConn.SetReadDeadline(time.Now().Add(60 * time.Second))
	//设置pong处理方式
	c.WsConn.SetPongHandler(func(string) error {
		c.WsConn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	for {
		_, data, err := c.WsConn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}

		err = redis.TalkFrequency(strconv.FormatInt(m.UserID, 10), m.Username, m.LessonID)
		if err != nil {
			zap.L().Error("redis.TalkFrequency(strconv.FormatInt(m.UserID, 10), m.Username, m.LessonID) failed",zap.Error(err))
			return 
		}

		//传入广播通道
		msg := Message{data, m.LessonID, m.UserID, m.Username, c}
		switch channel {
		case KeyBroadcast:
			//广播通道
			H.Broadcast <- msg
			break
		case KeyEager:
			//抢答 提出问题
			H.Eager <- msg
			break
		case KeyAnswer:
			//回答问题
			H.Answer <- msg
			break
		case KeyChoose:
			H.Choose <- msg
		}

	}
}

func (m Message) Write() {
	c := m.Conn
	ticker := time.NewTicker(50 * time.Second)

	defer func() {
		ticker.Stop()
		c.WsConn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func (c *Connection) write(mt int, payload []byte) error {
	c.WsConn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	return c.WsConn.WriteMessage(mt, payload)
}
