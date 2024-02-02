package decode

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/go-taken/gun/pkg"
	"github.com/go-taken/gun/src/html"
)

var Document = ""
var logger = pkg.NewLogger()

func NewJsonToHTML(path string) {

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	filename := strings.Replace(path, ".json", ".html", -1)
	// Decode JSON dari file
	var tag []html.HTML
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tag)
	if err != nil {
		logger.Error(err)
		return
	}
	Document += "<!DOCTYPE html> \n"
	GenerateTag(tag)
	GenerateFile(filename)
}

func GenerateTag(tags []html.HTML) {
	// pkg.Prety(tags)
	for _, tag := range tags {
		attrs := ""
		for _, attr := range tag.Attr {
			attrs += fmt.Sprintf(` %s="%s"`, attr.Name, attr.Value)
		}
		ok := pkg.ValidateVoidElement(tag.Tag)
		if ok {
			Document += fmt.Sprintf("\n<%s %s />", tag.Tag, attrs)
			Document += fmt.Sprintf("\n%s", tag.Value)
		} else {
			Document += fmt.Sprintf("\n<%s %s> ", tag.Tag, attrs)
			Document += fmt.Sprintf("\n%s", tag.Value)
		}
		if tag.Content != nil {
			GenerateTag(tag.Content)
			if ok := pkg.ValidateVoidElement(tag.Tag); !ok {
				Document += fmt.Sprintf("\n</%s>", tag.Tag)
			}
		} else {
			if ok := pkg.ValidateVoidElement(tag.Tag); !ok {
				Document += fmt.Sprintf("\n</%s>", tag.Tag)
			}
		}
	}
}

func GenerateFile(filename string) error {
	os.RemoveAll("dist")
	if err := os.Mkdir("dist", 0755); err != nil {
		return err
	}

	err := os.WriteFile("dist/"+filename, []byte(Document), 0644)
	if err != nil {
		return err
	}
	return nil
}
