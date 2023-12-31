package pkg

var voidElement = []string{
	"area",
	"base",
	"br",
	"col",
	"embed",
	"hr",
	"img",
	"input",
	"link",
	"meta",
	"param",
	"source",
	"track",
	"wbr",
}

// ValidateVoidElement ...
// input validate void element (</br>)
// output true or false
func ValidateVoidElement(element string) bool {
	for _, v := range voidElement {
		if v == element {
			return true
		}
	}
	return false
}
