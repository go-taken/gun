package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var html = ""

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
	var tag HTML
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tag)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	html += "<!DOCTYPE html> \n"
	html += "<html>"
	checkContent(tag.Content)
	html += "\n</html>"
	CreateDocument(html)
}

func checkContent(htmlContent any) {
	text := ""
	contens, err := UnmarshalContent(htmlContent)
	if err != nil {
		res, err := UnmarshalContentArray(htmlContent)
		if err != nil {
			fmt.Println(err)
		}
		text = strings.Join(res, " ")
		html += strings.Join(res, "")
	}
	for _, content := range contens {
		if tag := content.Tag != ""; tag {
			class := ""
			if content.Attr.Class != "" {
				class = fmt.Sprintf("class=\"%s\"", content.Attr.Class)
			}
			
			id := ""
			if content.Attr.ID != "" {
				id = fmt.Sprintf("id=\"%s\"", content.Attr.ID)
			}

			typeTag := ""
			if content.Attr.Type != "" {
				typeTag = fmt.Sprintf("type=\"%s\"", content.Attr.Type)
			}

			name := ""
			if content.Attr.Name != "" {
				name = fmt.Sprintf("name=\"%s\"", content.Attr.Name)
			}

			value := ""
			if content.Attr.Value != "" {
				value = fmt.Sprintf("value=\"%s\"", content.Attr.Value)
			}

			html += fmt.Sprintf("\n<%s %s %s %s %s %s>", content.Tag, class, id, typeTag, name, value)
		}
		if content.Content != nil {
			checkContent(content.Content)
			html += text
			html += fmt.Sprintf("\n</%s>  ", content.Tag)
			text = ""
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
	Class   string `json:"class"`
	ID      string `json:"id"`
	Attr    Attr   `json:"attr"`
	Content []any  `json:"content"`
}

type Attr struct {
	Class string `json:"class"`
	ID    string `json:"id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Content struct {
	Tag     string `json:"tag"`
	Class   string `json:"class"`
	ID      string `json:"id"`
	Attr    Attr   `json:"attr"`
	Content []any  `json:"content"`
}
type ContentString struct {
	Content []string `json:"content"`
}
