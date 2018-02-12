package logic

const (
	//聊天相关命令
	CmdRoomTalk          = 1001
	CmdDeskTalk          = 1002
	CmdDeskTalkBroadcast = 1003
	CmdRoomTalkBroadcast = 1004

	//房间相关命令
	CmdEnterRoom = 2001
	CmdLeaveRoom = 2002

	//游戏相关命令
	CmdSitdown              = 3001
	CmdLeaveDesk            = 3002
	CmdGameReady            = 3003
	CmdGameCancelReady      = 3004
	CmdGameStartBroadcat    = 3005
	CmdSelectLandlord       = 3006
	CmdSelectLandlordResult = 3007
)
