package main

import (
	"github.com/go-taken/gun/src/html"
)

func main() {
	html := html.New()
	head := html.Head
	head.SetTitle("Hello World")
	head.SetMeta(`name="viewport" content="width=device-width, initial-scale=1.0"`)
	head.SetScript(`src="https://cdn.jsdelivr.net/npm/vue/dist/vue.js"`)

	err := html.Start()
	if err != nil {
		panic(err)
	}
}
