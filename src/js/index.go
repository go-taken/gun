package js

import "github.com/go-taken/gun/src/js/console"

type JS struct {
	Console *console.Console
}

func New() *JS {
	return &JS{
		Console: console.New(),
	}
}


