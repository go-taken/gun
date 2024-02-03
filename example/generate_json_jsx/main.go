package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-taken/gun/src/decode"
	"github.com/go-taken/gun/src/html"
)

func main() {
	file, err := os.Open("./index.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	var tag []html.HTML
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tag)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsx := decode.NewDecodeJSX("Button")
	buttonJSX := jsx.ToJSXString(tag)
	fmt.Println(buttonJSX)

}
