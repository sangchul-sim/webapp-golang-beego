package controllers

import (
	"io"
	"net/http"
	"time"

	"github.com/sangchul-sim/webapp-golang-beego/chat"
	"github.com/sangchul-sim/webapp-golang-beego/chat/redis_stream"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"github.com/heroku/x/redis"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	waitTimeout   = time.Minute * 1
	redisReceiver *redis_stream.RedisReceiver
	RedisSender   *redis_stream.RedisSender
)

// ChatController 는 /api 에 대한 controller 입니다.
type ChatController struct {
	BaseController
}

func init() {
	redisURL := beego.AppConfig.String("redis::url")
	redisChannelName := beego.AppConfig.String("redis::channel")

	redisPool, err := redis.NewRedisPoolFromURL(redisURL)
	if err != nil {
		beego.Error("Unable to create Redis pool", redisURL)
	}

	redisReceiver = redis_stream.NewRedisReceiver(redisPool)
	RedisSender = redis_stream.NewRedisSender(redisPool)

	go func() {
		for {
			waited, err := redis.WaitForAvailability(redisURL, waitTimeout, redisReceiver.Wait)
			if !waited || err != nil {
				beego.Error("Redis not available by timeout!\twaitTimeout:", waitTimeout, " error:", err)
			}

			redisReceiver.Broadcast(chat.AvailableMessage)
			err = redisReceiver.Run(redisChannelName)
			if err == nil {
				break
			}
			beego.Error(err)
		}
	}()

	go func() {
		for {
			waited, err := redis.WaitForAvailability(redisURL, waitTimeout, nil)
			if !waited || err != nil {
				beego.Error("Redis not available by timeout!\twaitTimeout:", waitTimeout, " error:", err)
			}
			err = RedisSender.Run(redisChannelName)
			if err == nil {
				break
			}
			beego.Error(err)
		}
	}()
}

func (dis *ChatController) Home() {
	dis.TplName = "chat/home.html"
	dis.Render()
}

func (dis *ChatController) Ws() {
	// https://godoc.org/github.com/gorilla/websocket
	// The Conn type represents a WebSocket connection.
	// A server application uses the Upgrade function from an Upgrader object
	// with a HTTP request handler to get a pointer to a Conn:
	w := dis.Ctx.ResponseWriter
	r := dis.Ctx.Request

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		m := "Unable to upgrade to websockets"

		beego.Error(m, err)
		http.Error(w, m, http.StatusBadRequest)
		return
	}

	id := redisReceiver.Register(ws)
	beego.Debug("id", id)

	for {
		// messageType : int
		// data : []byte
		messageType, data, err := ws.ReadMessage()

		if err != nil {
			LogData := make(map[string]interface{})
			LogData["messageType"] = messageType
			LogData["data"] = data

			if err == io.EOF {
				beego.Info("Websocket closed!", LogData)
			} else {
				beego.Error("Error reading websocket message", LogData)
			}
			break
		}

		switch messageType {
		// https://github.com/gorilla/websocket/blob/master/conn.go
		// TextMessage denotes a text data message. The text message payload is
		// interpreted as UTF-8 encoded text data.
		case websocket.TextMessage:
			// msg : chat.message
			msg, err := chat.ValidateMessage(data)
			if err != nil {
				beego.Error("Invalid Message\tmsg:", msg, " err:", err)
				break
			}

			// TODO
			// msg.Handle 을 . 기준으로 split 해서 interface 를 return 받는 방법 확인할 것
			// split 된 handle 의 첫번째 값이 room 이면 SetChannel 하도록 로직 변경할 것
			// if msg.Handle == "room.message" {
			// 	RedisSender.SetChannel(msg.Room)
			// }

			RedisSender.Publish(data)

		// BinaryMessage denotes a binary data message.
		case websocket.BinaryMessage:

		// CloseMessage denotes a close control message. The optional message
		// payload contains a numeric code and text. Use the FormatCloseMessage
		// function to format a close message payload.
		case websocket.CloseMessage:

		// PingMessage denotes a ping control message. The optional message payload
		// is UTF-8 encoded text.
		case websocket.PingMessage:

		// PongMessage denotes a ping control message. The optional message payload
		// is UTF-8 encoded text.
		case websocket.PongMessage:

		default:
			beego.Error("Unknown Message!")
			// l.Warning("Unknown Message!")
		}
	}

	redisReceiver.DeRegister(id)

	ws.WriteMessage(websocket.CloseMessage, []byte{})
}
