package logic

type RoomMgr struct {
	roomList []*Room
	roomNum  int
}

func NewRoomMgr(roomConf []*RoomConf) (roomMgr *RoomMgr, err error) {
	if len(roomConf) == 0 {
		err = ErrInvalidParameter
		return
	}

	roomMgr = &RoomMgr{
		roomNum: len(roomConf),
	}

	for _, conf := range roomConf {
		room := NewRoom(conf.Id, conf.Name, conf.DeskNum, conf.MaxPlayerNum)
		go room.run()
		roomMgr.roomList = append(roomMgr.roomList, room)
	}
	return
}

func (r *RoomMgr) GetAllRoomList() []*Room {
	return r.roomList
}

func (r *RoomMgr) GetRoom(id string) {
	return
}
