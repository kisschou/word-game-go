package core

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type (
	WebSocket struct {
	}
)

var (
	upgrader = websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// http.HandleFunc("/ws", wsHandler)
// http.ListenAndServe("0.0.0.0:7777", nil)
func (ws *WebSocket) Handler(w http.ResponseWriter, r *http.Request) {
	//收到http请求(upgrade),完成websocket协议转换
	//在应答的header中放上upgrade:websoket
	var (
		conn *websocket.Conn
		err  error
		//msgType int
		data []byte
	)
	if conn, err = upgrader.Upgrade(w, r, nil); err != nil {
		//报错了，直接返回底层的websocket链接就会终断掉
		return
	}
	//得到了websocket.Conn长连接的对象，实现数据的收发
	for {
		//Text(json), Binary
		if _, data, err = conn.ReadMessage(); err != nil {
			//报错关闭websocket
			goto ERR

		}
		//发送数据，判断返回值是否报错
		if err = conn.WriteMessage(websocket.TextMessage, data); err != nil {
			//报错了
			goto ERR
		}
	}
	//error的标签
ERR:
	conn.Close()
}
