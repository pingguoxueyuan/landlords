package logic

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	roomMgr   *RoomMgr
	playerMgr *PlayerMgr
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Init(roomConf []*RoomConf, maxPlayerNum int) (err error) {
	roomMgr, err = NewRoomMgr(roomConf)
	playerMgr = NewPlayerMgr(maxPlayerNum)
	return
}

func HandleConn(w http.ResponseWriter, r *http.Request) (err error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	player := NewPlayer(conn)
	playerMgr.AddPlayer(player)

	player.run()
	return
}
