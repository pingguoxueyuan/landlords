package logic

import (
	"sync"
)

const (
	SeatStatusEmpty = iota
	SeatStatusUsed
)

type Seat struct {
	player *Player
	status int
	lock   sync.Mutex
	desk   *Desk
}

func NewSeat(desk *Desk) (s *Seat) {
	s = &Seat{
		desk: desk,
	}

	return s
}

func (s *Seat) isEmpty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.status == SeatStatusEmpty
}

func (s *Seat) isUsed() bool {
	return s.isEmpty() == false
}

func (s *Seat) sitDown(player *Player) (err error) {

	s.lock.Lock()
	defer s.lock.Unlock()

	if s.isUsed() {
		err = ErrSeatIsUsed
		return
	}

	s.player = player
	s.status = SeatStatusUsed
	return
}

func (s *Seat) isReady() bool {
	if s.isEmpty() {
		return false
	}

	return s.player.isReady()
}

func (s *Seat) leave() (err error) {
	if s.isEmpty() {
		err = ErrInvalidSeatStatus
		return
	}

	if s.player != nil {
		s.player.leaveDesk()
		s.player = nil
	}

	s.status = SeatStatusEmpty
	return
}
