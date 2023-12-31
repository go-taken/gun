package head

import (
	"fmt"
	"strings"
)

type Head struct {
	// Title is the title of the document.
	// Example:
	// <title>{slot}</title>
	Title string

	// Meta is the meta of the document.
	// Example:
	// <meta name="viewport" content="width=device-width, initial-scale=1.0">
	Meta []string

	// Script is the script of the document.
	// Example:
	// <script src="https://cdn.jsdelivr.net/npm/vue/dist/vue.js"></script>
	Script []string

	// Link is the link of the document.
	// Example:
	// <link rel="icon" href="https://cdn.jsdelivr.net/npm/vuetify/dist/vuetify.min.css">
	Link []string
}

func NewHead() *Head {
	return &Head{}
}

func (h *Head) SetTitle(title string) *Head {
	h.Title = fmt.Sprintf("\t <title>%s</title> \n", title)
	return h
}

func (h *Head) SetMeta(meta string) *Head {
	h.Meta = append(h.Meta, fmt.Sprintf("\t <meta %s> \n", meta))
	return h
}

func (h *Head) SetScript(script string) *Head {
	h.Script = append(h.Script, fmt.Sprintf("\t <script %s></script> \n", script))
	return h
}

func (h *Head) SetLink(link string) *Head {
	h.Link = append(h.Link, fmt.Sprintf("\t <link %s> \n", link))
	return h
}

func (h *Head) GetTitle() string { return h.Title }

func (h *Head) GetScript() string { return strings.Join(h.Script, "") }

func (h *Head) GetMeta() string { return strings.Join(h.Meta, "") }

func (h *Head) GetLink() string { return strings.Join(h.Link, "") }

func (h *Head) RenderHead() string {
	return fmt.Sprintf("<head>\n%s%s%s%s</head>", h.GetTitle(), h.GetMeta(), h.GetLink(), h.GetScript())
}
