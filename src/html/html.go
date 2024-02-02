package html

type HTML struct {
	Tag     string `json:"tag"`
	Attr    []Attr `json:"attr"`
	Value   string `json:"value"`
	Content []HTML `json:"content"`
}

type Attr struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func New() []HTML {
	html := []HTML{
		{
			Tag: "html",
			Attr: []Attr{
				{
					Name:  "lang",
					Value: "en",
				},
			},
			Content: []HTML{
				{
					Tag: "head",
					Content: []HTML{
						{
							Tag: "meta",
							Attr: []Attr{
								{
									Name:  "charset",
									Value: "UTF-8",
								},
							},
						},
						{
							Tag:     "title",
							Value:   "Title",
							Content: []HTML{},
						},
					},
				},
			},
		},
	}
	return html
}
