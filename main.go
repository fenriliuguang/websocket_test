package main

import (
	"fmt"
	"net/http"
	"t3/ws"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var liveLists ws.AliveList

func checkOrigin(r *http.Request) bool {
	return true
}

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     checkOrigin,
}

func cmdWebSocket(c *gin.Context) {

	wsConn, _ := upGrader.Upgrade(c.Writer, c.Request, nil)

	defer wsConn.Close()

	for {
		_, msg, _ := wsConn.ReadMessage()
		fmt.Printf("%s\n", msg)
	}
}

func main() {

	r := gin.Default()

	liveLists.Listen()

	r.GET("/room", cmdWebSocket)

	r.Run(":8800")
}
