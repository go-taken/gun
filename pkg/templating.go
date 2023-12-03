package pkg

import (
	"fmt"
	"reflect"
	"regexp"
)

func Render(input any, replacements map[string]any) (string, error) {
	inputType := reflect.TypeOf(input)
	if inputType.Kind() != reflect.String {
		return "", fmt.Errorf("input must be of type string")
	}

	re := regexp.MustCompile(`\{(\w+)\}`)
	output := re.ReplaceAllStringFunc(input.(string), func(m string) string {
		key := m[1 : len(m)-1] // Extract key from match
		if val, ok := replacements[key]; ok {
			return fmt.Sprintf("%v", val)
		}
		return m
	})

	return output, nil
}
