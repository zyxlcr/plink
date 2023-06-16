package websocket

import (
	"chatcser/pkg/plink/config"
	"chatcser/pkg/plink/connection"
	"chatcser/pkg/plink/iface"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/gorilla/websocket"
)

type WsServer struct {
	Config  *config.Config
	PServer iface.IServer

	Router iface.IRouter

	Gin      *gin.Engine
	WsRouter http.Handler

	//当前Server的消息管理模块，用来绑定MsgId和对应的处理方法
	MsgHandler iface.IMsgHandle
	//当前Server的链接管理器
	ConnMgr iface.IConnManager
	//该Server的连接创建时Hook函数
	OnConnStart func(conn iface.IConnection)
	//该Server的连接断开时的Hook函数
	OnConnStop func(conn iface.IConnection)

	Upgrader websocket.Upgrader
}

func NewWsServer(s iface.IServer) *WsServer {
	ws := &WsServer{
		Config:     s.GetConfig(),
		PServer:    s,
		Router:     s.GetRouter(),
		Gin:        gin.Default(),
		MsgHandler: s.GetMsgHandler(),
		ConnMgr:    s.GetConnMgr(),
		Upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true },
		},
	}
	h := ws.Gin
	h.GET("/ws", func(ctx *gin.Context) {
		handleWebSocketRequest(ctx.Writer, ctx.Request, ws)
	})
	ws.WsRouter = h
	return ws
}

func Start(s iface.IServer) *WsServer {
	ws := NewWsServer(s)
	//http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
	//	handleWebSocketRequest(w, r, ws)
	//})
	return ws
}

func Listen(s iface.IServer, handler http.Handler) error {
	fmt.Printf("[START] Webocket Server name: %s,listenner at port: %s, is starting\n", s.GetConfig().Name, s.GetConfig().WsConfig.WsPort)
	return http.ListenAndServe(s.GetConfig().WsConfig.WsPort, handler)

}

func handleWebSocketRequest(w http.ResponseWriter, r *http.Request, s *WsServer) {

	// Upgrade HTTP request to WebSocket
	conn, err := s.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}

	//3.2 设置服务器最大连接控制,如果超过最大连接，那么则关闭此新的连接
	if s.ConnMgr.LenWs() >= int(s.Config.TcpConfig.MaxConn) { //GlobalObject.MaxConn
		fmt.Println("超了")
		conn.Close()
	}

	//3.3 处理该新连接请求的 业务 方法， 此时应该有 handler 和 conn是绑定的
	dealConn := connection.NewWsConntion(s.PServer, conn)

	//3.4 启动当前链接的处理业务
	go dealConn.StartWs()

}

func handleMessage(conn *websocket.Conn, msg []byte, msgType int) {
	// Process message here...

	// Split message into chunks and send with header to avoid packet sticking
	chunkSize := 512
	for i := 0; i < len(msg); i += chunkSize {
		end := i + chunkSize
		if end > len(msg) {
			end = len(msg)
		}

		// Add header to each chunk
		header := make([]byte, 2)
		header[0] = byte(msgType)
		header[1] = byte(i / chunkSize)

		// Send chunk with header
		conn.WriteMessage(websocket.BinaryMessage, append(header, msg[i:end]...))
	}
}
