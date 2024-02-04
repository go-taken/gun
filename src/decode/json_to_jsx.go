package decode

import (
	"fmt"
	"sync"

	"github.com/go-taken/gun/pkg"
	"github.com/go-taken/gun/src/html"
)

// func NewJsonToJSX(path string) {

// 	file, err := os.Open(path)
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 		return
// 	}
// 	defer file.Close()
// 	filename := strings.Replace(path, ".json", ".jsx", -1)
// 	// Decode JSON dari file
// 	var tag []html.HTML
// 	decoder := json.NewDecoder(file)
// 	err = decoder.Decode(&tag)
// 	if err != nil {
// 		logger.Error(err)
// 		return
// 	}
// 	Document += `
// 		import React from 'react';

// 		export default function Page() {
// 				return (
// 					<>`
// 	GenerateJSX(tag)
// 	Document += `		</>
// 				)
// 			};`
// 	GenerateFile(filename)
// }
// func GenerateJSX(tags []html.HTML) {
// 	// pkg.Prety(tags)
// 	for _, tag := range tags {

// 		attrs := ""
// 		for _, attr := range tag.Attr {
// 			if attr.Name == "class" {
// 				attr.Name = "className"
// 			}
// 			attrs += fmt.Sprintf(` %s="%s"`, attr.Name, attr.Value)
// 		}
// 		if ok := pkg.ValidateVoidElement(tag.Tag); ok {
// 			Document += fmt.Sprintf("\n<%s %s />", tag.Tag, attrs)
// 			Document += fmt.Sprintf("\n%s", tag.Value)
// 		} else {
// 			Document += fmt.Sprintf("\n<%s %s> ", tag.Tag, attrs)
// 			Document += fmt.Sprintf("\n%s", tag.Value)
// 		}

// 		if tag.Content != nil {
// 			GenerateTag(tag.Content)
// 			if ok := pkg.ValidateVoidElement(tag.Tag); !ok {
// 				Document += fmt.Sprintf("\n</%s>", tag.Tag)
// 			}
// 		} else {
// 			if ok := pkg.ValidateVoidElement(tag.Tag); !ok {
// 				Document += fmt.Sprintf("\n</%s>", tag.Tag)
// 			}
// 		}
// 	}
// }

func NewDecodeJSX(name string) *Generate {
	return &Generate{
		Mutex: new(sync.Mutex),
		Name:  name,
		JSXComponent: `
import React from 'react';

export default function ` + name + `() { 
	return ( 
			<>`,
	}
}

func (g *Generate) newComponentJSX(data []html.HTML) {
	for _, tag := range data {
		attrs := ""
		for _, attr := range tag.Attr {
			if attr.Name == "class" {
				attr.Name = "className"
			}
			attrs += fmt.Sprintf(` %s="%s"`, attr.Name, attr.Value)
		}
		ok := pkg.ValidateVoidElement(tag.Tag)
		if ok {
			g.JSXComponent += fmt.Sprintf("\n<%s %s />", tag.Tag, attrs)
			g.JSXComponent += fmt.Sprintf("\n%s", tag.Value)
		} else {
			g.JSXComponent += fmt.Sprintf("\n<%s %s> ", tag.Tag, attrs)
			g.JSXComponent += fmt.Sprintf("\n%s", tag.Value)
		}
		if tag.Content != nil {
			g.newComponentJSX(tag.Content)
			if ok := pkg.ValidateVoidElement(tag.Tag); !ok {
				g.JSXComponent += fmt.Sprintf("\n</%s>", tag.Tag)
			}
		} else {
			if ok := pkg.ValidateVoidElement(tag.Tag); !ok {
				g.JSXComponent += fmt.Sprintf("\n</%s>", tag.Tag)
			}
		}
	}
}

func (g *Generate) ToJSXString(data []html.HTML) string {
	g.Mutex.Lock()
	g.newComponentJSX(data)
	defer g.Mutex.Unlock()
	return g.JSXComponent + "\n</>\n);};"
}
