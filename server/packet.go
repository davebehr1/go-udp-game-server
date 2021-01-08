package main

import "time"

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
	Payload   []byte
	timestamp time.Time
	ptype     PacketType
}

func (packet Packet) GetBytes() []byte {
	return packet.Payload
}

type RequestJoinPacket struct {
	Packet
}

type AcceptJoinPacket struct {
	Packet
}

func main() {
	c1 := Packet{make([]byte, 0), time.Now(), RequestJoin}
}
