package gameObjects

import "time"

type NetworkMessage struct {
	Packet      Packet
	sender      string
	RecieveTime time.Time
}
