package logic

import (
	"fmt"
)

type Room struct {
	id           string
	name         string
	deskList     []Desk
	playerMap    map[int]*Player
	playerChan   chan *Player
	roomTalkChan chan *RookTalkBroadcastResp
	maxPlayerNum int
}

func NewRoom(id, name string, deskNum, maxPlayerNum int) (r *Room) {
	r = &Room{
		id:           id,
		name:         name,
		deskList:     make([]Desk, deskNum),
		maxPlayerNum: maxPlayerNum,
		playerMap:    make(map[int]*Player, 1024),
		playerChan:   make(chan *Player, 1024),
		roomTalkChan: make(chan *RookTalkBroadcastResp, 1024),
	}

	return
}

func (r *Room) AddPlayer(player *Player) (err error) {
	if len(r.playerMap) >= r.maxPlayerNum {
		err = ErrReachMaxPlayer
		return
	}

	select {
	case r.playerChan <- player:
	default:
		err = ErrReachMaxPlayer
	}
	return
}

func (r *Room) run() {
	for {
		select {
		case p := <-r.playerChan:
			r.playerMap[p.userId] = p
		case talkReq := <-r.roomTalkChan:
			r.handleBroadcaseTalk(talkReq)
		}
	}
}

func (r *Room) handleBroadcaseTalk(resp *RookTalkBroadcastResp) {

	fmt.Println("Handle resp")
	for _, v := range r.playerMap {
		if v.userId == resp.FromUid {
			continue
		}

		header := CommHeader{
			Cmd: CmdRoomTalkBroadcast,
		}

		err := v.send(header)
		if err != nil {
			//TODO: write log
		}

		fmt.Println("send:", resp)
		err = v.send(resp)
		if err != nil {
			//TODO: write log
		}
	}

}

func (r *Room) broadcastTalk(req *RookTalkBroadcastResp) (err error) {
	select {
	case r.roomTalkChan <- req:
	default:
		err = ErrServerBusy
		return
	}
	return
}
