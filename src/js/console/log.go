package console

import "log"

type Console struct {
}

func New() *Console {
	return &Console{}
}

func (c *Console) Log(logs ...any) {
	logger := log.Default()
	logger.Println(logs...)
}
