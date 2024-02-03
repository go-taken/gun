package decode

import (
	"sync"
)

type Generate struct {
	*sync.Mutex
	Vue           string
	JSX           string
	HTML          string
	VueComponent  string
	JSXComponent  string
	HTMLComponent string
	Name          string
}
