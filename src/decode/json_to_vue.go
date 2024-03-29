package decode

import (
	"sync"

	"github.com/go-taken/gun/pkg"
	"github.com/go-taken/gun/src/html"
)

// func NewJsonToVue(path string) {

// 	file, err := os.Open(path)
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 		return
// 	}
// 	defer file.Close()
// 	filename := strings.Replace(path, ".json", ".vue", -1)
// 	// Decode JSON dari file
// 	var tag []html.HTML
// 	decoder := json.NewDecoder(file)
// 	err = decoder.Decode(&tag)
// 	if err != nil {
// 		logger.Error(err)
// 		return
// 	}
// 	Document += `<template>`
// 	GenerateVue(tag)
// 	Document += `</template>`
// 	GenerateFile(filename)
// }
// func GenerateVue(tags []html.HTML) {
// 	// pkg.Prety(tags)
// 	for _, tag := range tags {

// 		attrs := ""
// 		for _, attr := range tag.Attr {
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

func NewDecodeVue(name string) *Generate {
	return &Generate{
		Mutex: new(sync.Mutex),
		Name:  name,
		Vue:   "<template>",
	}
}

func (g *Generate) newComponentVue(tags []html.HTML) {
	for _, tag := range tags {
		attrs := ""
		for _, attr := range tag.Attr {
			attrs += ` ` + attr.Name + `="` + attr.Value + `"`
		}
		if ok := pkg.ValidateVoidElement(tag.Tag); ok {
			g.Vue += "\n<" + tag.Tag + attrs + " />"
			g.Vue += "\n" + tag.Value
		} else {
			g.Vue += "\n<" + tag.Tag + attrs + ">"
			g.Vue += "\n" + tag.Value
		}

		if tag.Content != nil {
			g.newComponentVue(tag.Content)
			if ok := pkg.ValidateVoidElement(tag.Tag); !ok {
				g.Vue += "\n</" + tag.Tag + ">"
			}
		} else {
			if ok := pkg.ValidateVoidElement(tag.Tag); !ok {
				g.Vue += "\n</" + tag.Tag + ">"
			}
		}
	}
}

func (g *Generate) ToVueString(tags []html.HTML) string {
	g.Mutex.Lock()
	g.newComponentVue(tags)
	defer g.Mutex.Unlock()
	return g.Vue + "\n</template>"
}
