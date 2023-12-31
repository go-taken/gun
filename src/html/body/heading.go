package body

import "fmt"

type Heading struct {
	H1 []*string
	H2 []*string
	H3 []*string
	H4 []*string
	H5 []*string
	H6 []*string
}

func NewHeading() *Heading {
	return &Heading{}
}

func (h *Heading) SetH1(h1 string) *Heading {
	val := fmt.Sprintf("\t <h1> \n \t\t%s \n \t</h1> \n", h1)
	h.H1 = append(h.H1, &val)
	return h
}

func (h *Heading) SetH2(h2 string) *Heading {
	val := fmt.Sprintf("\t <h2> \n \t\t%s \n \t</h2> \n", h2)
	h.H2 = append(h.H2, &val)
	return h
}

func (h *Heading) SetH3(h3 string) *Heading {
	val := fmt.Sprintf("\t <h3> \n \t\t%s \n \t</h3> \n", h3)
	h.H3 = append(h.H3, &val)
	return h
}

func (h *Heading) SetH4(h4 string) *Heading {
	val := fmt.Sprintf("\t <h4> \n \t\t%s \n \t</h4> \n", h4)
	h.H4 = append(h.H4, &val)
	return h
}

func (h *Heading) SetH5(h5 string) *Heading {
	val := fmt.Sprintf("\t <h5> \n \t\t%s \n \t</h5> \n", h5)
	h.H5 = append(h.H5, &val)
	return h
}

func (h *Heading) SetH6(h6 string) *Heading {
	val := fmt.Sprintf("\t <h6> \n \t\t%s \n \t</h6> \n", h6)
	h.H6 = append(h.H6, &val)
	return h
}

func (h *Heading) GetH1() string {
	var h1 string
	for _, v := range h.H1 {
		h1 += *v
	}
	return h1
}

func (h *Heading) GetH2() string {
	var h2 string
	for _, v := range h.H2 {
		h2 += *v
	}
	return h2
}

func (h *Heading) GetH3() string {
	var h3 string
	for _, v := range h.H3 {
		h3 += *v
	}
	return h3
}

func (h *Heading) GetH4() string {
	var h4 string
	for _, v := range h.H4 {
		h4 += *v
	}
	return h4
}

func (h *Heading) GetH5() string {
	var h5 string
	for _, v := range h.H5 {
		h5 += *v
	}
	return h5
}

func (h *Heading) GetH6() string {
	var h6 string
	for _, v := range h.H6 {
		h6 += *v
	}
	return h6
}

func (h *Heading) RenderHeading() string {
	return fmt.Sprintf("<h1>\n %s \n</h1>", h.GetH1())
}
