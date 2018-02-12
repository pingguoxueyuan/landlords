package logic

type CloseNotifyer interface {
	OnClose(interface{})
}
