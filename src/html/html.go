package html

import (
	_ "embed"

	"github.com/go-taken/gun/pkg"
)

type HTML struct {
	Head string
	Body string
}

func New() *HTML {
	return &HTML{}
}

func (h *HTML) SetHead(head string) *HTML {
	h.Head = head
	return h
}

func (h *HTML) SetBody(body string) *HTML {
	h.Body = body
	return h
}

//go:embed app.html
var app string

func (h *HTML) Start() (any, error) {
	txt, err := pkg.ReplaceWithText(app, map[string]any{
		"head": h.Head,
		"body": h.Body,
	})
	return txt, err
}
