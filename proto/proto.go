package proto

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

const (
	MaxPacketLen = 16 * 1024
)

type Packet struct {
	Length uint32 `json:"length"`
	Cmd    uint32 `json:"cmd"`
	Data   []byte `json:"data"`
}

func Read(conn net.Conn) (pack *Packet, err error) {

	//读取长度
	var buf [4]byte
	_, err = io.ReadFull(conn, buf[:])
	if err != nil {
		return
	}

	pack = &Packet{}
	pack.Length = binary.BigEndian.Uint32(buf[:])
	if pack.Length > MaxPacketLen {
		err = fmt.Errorf("[%d]reach max packet length:[%d]", pack.Length, MaxPacketLen)
		return
	}

	//读取命令号
	_, err = io.ReadFull(conn, buf[:])
	if err != nil {
		return
	}
	pack.Cmd = binary.BigEndian.Uint32(buf[:])

	pack.Data = make([]byte, pack.Length)
	_, err = io.ReadFull(conn, pack.Data)
	return
}

func Write(conn net.Conn, pack *Packet) (err error) {

	pack.Length = uint32(len(pack.Data))
	var buffer bytes.Buffer
	var buf [4]byte

	binary.BigEndian.PutUint32(buf[:], pack.Length)
	_, err = buffer.Write(buf[:])
	if err != nil {
		return
	}

	binary.BigEndian.PutUint32(buf[:], pack.Cmd)
	_, err = buffer.Write(buf[:])
	if err != nil {
		return
	}

	_, err = buffer.Write(pack.Data)
	if err != nil {
		return
	}

	_, err = buffer.WriteTo(conn)
	if err != nil {
		return
	}
	return
}
