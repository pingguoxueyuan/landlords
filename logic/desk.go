package logic

const (
	DeskStatusEmpty = iota
	DeskStatusNotReady
	DeskStatusStart
	DeskStatusSelectLandlord
	DeskStatusPlaying
	DeskStatusEnd
)

type Desk struct {
	seats         [3]*Seat
	status        int
	turn          int
	landlordIndex int
}

func NewDesk() (d *Desk) {
	d = &Desk{}

	d.seats[0] = NewSeat(d)
	d.seats[1] = NewSeat(d)
	d.seats[2] = NewSeat(d)
	return
}

func (d *Desk) isGameStart() bool {
	return d.status == DeskStatusStart
}
