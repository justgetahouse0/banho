package packets

import (
	"encoding/binary"
	"bytes"
)

type Packet []byte
type PacketType int16

func Make(h PacketType, dataLen uint32, data interface{}) (Packet) {
	buff := new(bytes.Buffer)
	binary.Write(buff, binary.LittleEndian, h)
	binary.Write(buff, binary.LittleEndian, byte(0))
	binary.Write(buff, binary.LittleEndian, dataLen)
	binary.Write(buff, binary.LittleEndian, data)
	packet := Packet(buff.Bytes())
	return packet
}