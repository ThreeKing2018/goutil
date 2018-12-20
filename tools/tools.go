package tools

import "strings"

//替换成空
func ReplaceNull(output string, oldSlice []string, newStr string) string {
	for _, old := range oldSlice {
		output = strings.Replace(output, old, newStr, -1)
	}
	return output
}
