package html

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/go-taken/gun/pkg"
	"github.com/go-taken/gun/src/html/head"
)

type HTML struct {
	Head   *head.Head
	Body   string
	render func(string, map[string]any) (string, error)
}

func New() *HTML {
	return &HTML{
		Head: head.NewHead(),
	}
}

func (h *HTML) SetBody(body string) *HTML {
	h.Body = body
	return h
}

//go:embed app.html
var app string

func (h *HTML) Start() error {
	os.RemoveAll("dist")
	if err := os.Mkdir("dist", 0755); err != nil {
		return err
	}
	html, err := h.run().render(app, map[string]any{
		"head": h.Head.RenderHead(),
		"body": h.Body,
	})
	if err != nil {
		return err
	}
	fmt.Println(html)
	os.WriteFile("dist/index.html", []byte(html), 0644)
	return err
}

func (h *HTML) run() *HTML {
	h.render = func(s string, m map[string]any) (string, error) {
		return pkg.Render(s, m)
	}

	return h
}
