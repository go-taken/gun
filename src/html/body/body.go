package body

import (
	"fmt"
)

type Body struct {
	// BodyDiv is a slice of strings that will be used to create the body of the HTML document.
	// Each string in the slice will be wrapped in a <div> tag.
	// Example:
	// BodyDiv = []string{"<h1>Hello World</h1>"}
	// Result:
	// <div>
	// 	<h1>Hello World</h1>
	// </div>
	Div []*string

	*Heading
}

func NewBody() *Body {
	return &Body{
		Heading: NewHeading(),
	}
}

func (h *Body) SetDiv(div string) *Body {
	val := fmt.Sprintf("\t <div> \n \t\t%s \n \t</div> \n", div)
	h.Div = append(h.Div, &val)
	return h
}

func (h *Body) GetDiv() string {
	var div string
	for _, v := range h.Div {
		div += *v
	}
	return div
}

func (h *Body) RenderBody() string {
	prt := h.Heading.RenderHeading()
	fmt.Println(prt)
	return fmt.Sprintf("<body>\n %s \n</body>", h.GetDiv())
}
