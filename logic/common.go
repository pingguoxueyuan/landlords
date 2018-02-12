package logic

type CommHeader struct {
	Cmd int `json:"cmd"`
}

type RespBase struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
}

type RoomTalkReq struct {
	RoomId  string `json:"room_id"`
	Content string `json:"content"`
	UserId  int    `json:"user_id"`
}

type RoomTalkResp struct {
	RespBase
}

type RookTalkBroadcastResp struct {
	RoomId   string `json:"room_id"`
	Content  string `json:"content"`
	FromUid  int    `json:"from_uid"`
	Nickname string `json:"nick_name"`
}

type LoginReq struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
	DeviceId string `json:"device_id"`
	Platform string `json:"platform"`
	Version  int    `json:"version"`
	IP       string `json:"ip"`
}

type LoginResp struct {
	RespBase
	LastLogin string `json:"ip"`
	Token     string `json:"token"`
}
