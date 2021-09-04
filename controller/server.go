package controller

import (
	"go.uber.org/zap"
	"ketangpai/models"
	"ketangpai/pkg/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

//设置参数
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Server 服务器
func Server(c *gin.Context) {

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		}, //跨域true忽略
		Subprotocols: []string{c.GetHeader("Sec-WebSocket-Protocol")}, // 处理 Sec-WebSocket-Protocol Header
	}
	upgradeHeader := http.Header{}
	if hdr := c.Request.Header.Get("Sec-Websocket-Protocol"); hdr != "" {
		upgradeHeader.Set("Sec-Websocket-Protocol", hdr)
	}

	wsConn, err := upgrader.Upgrade(c.Writer, c.Request, upgradeHeader)
	if err != nil {
		zap.L().Error("socket connect failed", zap.Error(err))
		return
	}

	//获取进入的房间
	channel:=c.Param("channel")
	roomid := c.Param("roomid")
	token := c.Param("token")
	Claims, err := jwt.ParseToken(token)
	if err != nil {
		zap.L().Error("get token failed", zap.Error(err))
		return
	}

	//把信息传入新进入聊天室的通道 进行广播
	con := &models.Connection{Send: make(chan []byte, 256), WsConn: wsConn}
	msg := models.Message{Roomid: roomid, Username: Claims.UserName, Conn: con}



	models.H.Join <- msg
	go msg.Write()
	go msg.Read(channel)

}
