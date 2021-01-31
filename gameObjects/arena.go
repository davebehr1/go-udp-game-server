package gameObjects

import (
	"fmt"
	"sync"
)

type ArenaState int

var nextId int = 1

const (
	NotRunning ArenaState = iota
	WaitingForPlayers
	NotifyingGameStart
	InGame
	GameOver
)

type Arena struct {
	sync.RWMutex
	State        ArenaState
	leftPlayer   PlayerInfo
	rightPlayer  PlayerInfo
	Id           int
	messageQueue *ConcurrentQueue
}

func newArena() *Arena {
	nextId++
	arena := &Arena{
		State: NotRunning,
		leftPlayer: PlayerInfo{
			paddle: newPaddle(LEFT, cfg),
		},
		rightPlayer: PlayerInfo{
			paddle: newPaddle(RIGHT, cfg),
		},
		Id:           nextId,
		messageQueue: NewConcurrentQueue(12),
	}

	return arena
}

func (a *Arena) tryAddPlayer(playerIp string) bool {
	if a.State == WaitingForPlayers {
		if a.leftPlayer.IsSet() == true {
			a.leftPlayer.endpoint = playerIp
			return true
		}

		if a.rightPlayer.IsSet() == true {
			a.rightPlayer.endpoint = playerIp
			return true
		}

	}
	return false
}

func (a *Arena) start() {
	a.Lock()
	a.State = WaitingForPlayers
	a.Unlock()

	go a.arenaRun()

}

func (a *Arena) handleConnectionSetup(player PlayerInfo, message NetworkMessage) {
	switch message.Packet.ptype {
	case RequestJoin:
		fmt.Println("Join request from player")
	case AcceptJoinAck:
	case Heartbeat:
	default:

	}

}

func (a *Arena) arenaRun() {
	var running bool = true
	for running {
		fmt.Println("running")

		switch a.State {
		case WaitingForPlayers:

			handleConnectionSetup(a.leftPlayer, message)
			fmt.Println("ja")
		case NotifyingGameStart:
			fmt.Println("no")
		case InGame:
			fmt.Println("in game")

		}
	}
}

func sendAcceptJoin(player PlayerInfo) {
	// ajp := AcceptJoinPacket{}
	// ajp.
}
