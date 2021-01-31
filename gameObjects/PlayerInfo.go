package gameObjects

type PlayerInfo struct {
	paddle     *Paddle
	endpoint   string
	havePaddle bool
	ready      bool
}

func (p PlayerInfo) IsSet() bool {
	if p.endpoint != "" {
		return true
	} else {
		return false
	}
}
