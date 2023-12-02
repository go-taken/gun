package main

import (
	_ "embed"
	"fmt"

	"github.com/go-taken/gun/src/html"
)

func main() {
	html := html.New()
	html.SetHead("head").SetBody("body")
	txt, _ := html.Start()
	fmt.Println(txt)
}
