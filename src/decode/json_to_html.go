package decode

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/go-taken/gun/pkg"
)

var html = ""
var logger = pkg.NewLogger()

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

func NewJsonToHTML(path string) {

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	filename := strings.Replace(path, ".json", ".html", -1)
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
	createDocument(filename, html)
}

func checkContent(tags any) {
	contens, err := unmarshalContent(tags)
	if err != nil {
		res, err2 := unmarshalContentArray(tags)
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

func unmarshalContent(htmlContent any) ([]Content, error) {
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

func unmarshalContentArray(htmlContent any) ([]string, error) {
	var stringArray []string
	dataJson, err := json.Marshal(htmlContent)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(dataJson, &stringArray)
	return stringArray, err
}

func createDocument(filename, html string) error {
	os.RemoveAll("dist")
	if err := os.Mkdir("dist", 0755); err != nil {
		return err
	}

	err := os.WriteFile("dist/"+filename, []byte(html), 0644)
	if err != nil {
		return err
	}
	return nil
}
