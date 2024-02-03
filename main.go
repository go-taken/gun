package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-taken/gun/src/decode"
	"github.com/go-taken/gun/src/html"
)

func main() {
	file, err := os.Open("./button.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	// filename := strings.Replace(path, ".json", ".html", -1)
	// Decode JSON dari file
	var tag []html.HTML
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tag)
	if err != nil {
		fmt.Println(err)
		return
	}
	html := decode.NewDecodeHTML("Button")
	button := html.ToHTMLComponentString(tag)
	fmt.Println(button)

	jsx := decode.NewDecodeJSX("Button")
	buttonJSX := jsx.ToJSXString(tag)
	fmt.Println(buttonJSX)

	vue := decode.NewDecodeVue("Button")
	buttonVue := vue.ToVueString(tag)
	fmt.Println(buttonVue)
	// go decode.NewJsonToHTML("./button.json")
	// pkg.Prety(html)
}
