package pkg

import "encoding/json"

func Prety(i any) {
	text, _ := json.MarshalIndent(i, "", " ")
	println(string(text))
}
