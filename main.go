package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/go-taken/gun/pkg"
)

var html = ""
var logger = pkg.NewLogger()

func main() {
	// html := html.New()
	// html.SetHead()
	// html.SetBody()
	// html.Body.SetDiv("blablabal").SetDiv("asdfasdf")

	// head := html.Head
	// head.SetTitle("Hello World")
	// head.SetMeta(`name="viewport" content="width=device-width, initial-scale=1.0"`)
	// head.SetScript(`src="https://cdn.jsdelivr.net/npm/vue/dist/vue.js"`)

	// err := html.Start()
	// if err != nil {
	// 	panic(err)
	// }
	file, err := os.Open("./data.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Decode JSON dari file
	var tag []HTML
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tag)
	if err != nil {
		logger.Error(err)
		return
	}
	html += "<!DOCTYPE html> \n"
	checkContent(tag)
	CreateDocument(html)
}

func checkContent(tags any) {
	contens, err := UnmarshalContent(tags)
	if err != nil {
		res, err2 := UnmarshalContentArray(tags)
		if err2 != nil {
			logger.Error(err2)
		}
		html += fmt.Sprintf("\n%s", strings.Join(res, " "))
	}
	for _, content := range contens {
		if tag := content.Tag != ""; tag {
			attrs := ""
			for _, attr := range content.Attr {
				attrs += fmt.Sprintf(` %s="%s"`, attr.Name, attr.Value)
			}
			if ok := pkg.ValidateVoidElement(content.Tag); ok {
				html += fmt.Sprintf("\n<%s %s />", content.Tag, attrs)
			} else {
				html += fmt.Sprintf("\n<%s %s>", content.Tag, attrs)
			}
		}
		if content.Content != nil {
			checkContent(content.Content)
			if ok := pkg.ValidateVoidElement(content.Tag); !ok {
				html += fmt.Sprintf("\n</%s>", content.Tag)
			}
		}
	}
}

func UnmarshalContent(htmlContent any) ([]Content, error) {
	content := []Content{}
	dataJson, err := json.Marshal(htmlContent)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(dataJson, &content)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func UnmarshalContentArray(htmlContent any) ([]string, error) {
	var stringArray []string
	dataJson, err := json.Marshal(htmlContent)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(dataJson, &stringArray)
	return stringArray, err
}

func CreateDocument(html string) error {
	os.RemoveAll("dist")
	if err := os.Mkdir("dist", 0755); err != nil {
		return err
	}

	err := os.WriteFile("dist/test.html", []byte(html), 0644)
	if err != nil {
		return err
	}
	return nil
}

type HTML struct {
	Tag     string `json:"tag"`
	Attr    []Attr `json:"attr"`
	Content []any  `json:"content"`
}

type Attr struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Content struct {
	Tag     string `json:"tag"`
	Class   string `json:"class"`
	ID      string `json:"id"`
	Attr    []Attr `json:"attr"`
	Content []any  `json:"content"`
}
