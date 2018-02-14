package logic

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

const (
	//写超时
	writeWait = 10 * time.Second

	//心跳超时
	pongWait = 60 * time.Second

	//心跳发送间隔
	pingPeriod = (pongWait * 9) / 10

	//消息最大长度
	maxMessageSize = 512
)

type Client struct {
	conn              *websocket.Conn
	recvChan          chan interface{}
	sendChan          chan interface{}
	closeNotifyerList []CloseNotifyer
}

// 新建链接
func NewClient(conn *websocket.Conn) (client *Client) {

	client = &Client{}
	client.conn = conn
	client.recvChan = make(chan interface{}, 256)
	client.sendChan = make(chan interface{}, 256)

	return
}

func (p *Player) isReady() bool {
	return p.status == PlayerStatusReady
}

func (p *Player) leaveDesk() {

	switch p.status {
	case PlayerStatusIdle:
		return
	case PlayerStatusSitDown:
		fallthrough
	case PlayerStatusReady:
		p.status = PlayerStatusIdle
	case PlayerStatusPlaying:
		p.status = PlayerStatusRobot
	}
	return
}

func (p *Player) send(message interface{}) (err error) {
	select {
	case p.sendChan <- message:
	default:
		return
	}
	return
}

func (p *Player) enterRoom(roomId int) (err error) {
	return
}

func (p *Player) process() {

	var loginedProc bool
	for req := range p.recvChan {
		if !loginedProc {
			err := p.handleLogin(req)
			if err != nil {
				fmt.Printf("handle login failed, err:%v\n", err)
				return
			}

			loginedProc = true
			continue
		}

		p.handleRequest(req)
	}
}

func (p *Player) handleLogin(req interface{}) (err error) {
	loginReq, ok := req.(*LoginReq)
	if !ok {
		err = ErrInvalidParameter
		p.conn.Close()
		return
	}

	return
}

func (p *Player) handleRequest(req interface{}) (err error) {
	switch v := req.(type) {
	case *RoomTalkReq:
		p.procRoomTalk(v)
	default:
		fmt.Printf("invalid req:%+v %T\n", v, v)
	}
	return
}

func (p *Player) procRoomTalk(roomTalkReq *RoomTalkReq) (err error) {

	var broadcastTalk = &RookTalkBroadcastResp{}
	broadcastTalk.Content = roomTalkReq.Content
	broadcastTalk.FromUid = roomTalkReq.UserId
	broadcastTalk.Nickname = p.name
	broadcastTalk.RoomId = roomTalkReq.RoomId

	//err = p.room.broadcastTalk(broadcastTalk)
	roomMgr.roomList[0].broadcastTalk(broadcastTalk)
	return
}

func (p *Player) run() {

	go p.write()
	go p.read()
	go p.process()
}

func (p *Player) AddCloseNotifyer(notifyer CloseNotifyer) {
	p.closeNotifyerList = append(p.closeNotifyerList, notifyer)
}

//read from network
func (p *Player) read() {
	defer func() {
		for _, v := range p.closeNotifyerList {
			v.OnClose(p)
		}

		p.conn.Close()
	}()

	p.conn.SetReadLimit(maxMessageSize)
	p.conn.SetReadDeadline(time.Now().Add(pongWait))

	pongHandler := func(string) error {
		p.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	}

	p.conn.SetPongHandler(pongHandler)
	for {
		var header CommHeader
		err := p.conn.ReadJSON(&header)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}

		log.Printf("recv header:%v\n", header)
		var req interface{}
		switch header.Cmd {
		case CmdRoomTalk:
			req = &RoomTalkReq{}
		default:
			break

		}

		err = p.conn.ReadJSON(req)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}

		log.Printf("recv req:%v\n", req)
		p.recvChan <- req
	}
}

//write to network
func (p *Player) write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		for _, v := range p.closeNotifyerList {
			v.OnClose(p)
		}
		ticker.Stop()
		p.conn.Close()
	}()
	for {
		select {
		case resp, ok := <-p.sendChan:
			p.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				p.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			err := p.conn.WriteJSON(resp)
			if err != nil {
				return
			}
		case <-ticker.C:
			p.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := p.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
