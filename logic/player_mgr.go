package logic

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type PlayerMgr struct {
	playerMap    map[*Player]bool
	maxPlayerNum int32
	lock         sync.Mutex
}

func NewPlayerMgr(maxPlayerNum int) *PlayerMgr {
	playerMgr := &PlayerMgr{
		playerMap:    make(map[*Player]bool, 32),
		maxPlayerNum: int32(maxPlayerNum),
	}

	return playerMgr
}

func (p *PlayerMgr) AddPlayer(player *Player) (err error) {
	if len(p.playerMap) > int(p.maxPlayerNum) {
		err = ErrReachMaxPlayer
		return
	}

	atomic.AddInt32(&p.maxPlayerNum, 1)
	p.lock.Lock()
	defer p.lock.Unlock()

	p.playerMap[player] = true
	player.AddCloseNotifyer(p)
	return
}

func (p *PlayerMgr) DelPlayer(player *Player) (err error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	_, ok := p.playerMap[player]
	if !ok {
		err = ErrNotFoundPlayer
		return
	}

	delete(p.playerMap, player)
	return
}

func (p *PlayerMgr) OnClose(obj interface{}) {
	player, ok := obj.(*Player)
	if !ok {
		fmt.Println("player:", player, " not ok")
		return
	}

	fmt.Println("player:", player, " is closed")
	p.DelPlayer(player)
}
