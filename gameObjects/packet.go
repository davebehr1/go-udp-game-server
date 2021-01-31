package gameObjects

import (
	"bytes"
	"encoding/gob"
	"time"
)

type PacketType uint8

const (
	RequestJoin PacketType = iota
	AcceptJoin
	AcceptJoinAck
	Heartbeat
	HeartbeatAck
	GameStart
	PaddlePosition
	GameState
	PlaySoundEffect
	Bye
)

type Packet struct {
	Payload   string
	timestamp time.Time
	ptype     PacketType
}

func (packet Packet) getBytes() []byte {

	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)

	err := encoder.Encode(packet)

	if err != nil {
		panic(err)

	}
	return buffer.Bytes()
}

type RequestJoinPacket struct {
	Packet
}

type AcceptJoinPacket struct {
	Packet
}
