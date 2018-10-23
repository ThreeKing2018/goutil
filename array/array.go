package array

import (
	"strconv"
	"strings"
)

//int 转 string
func IntArrToInString(i []int) string {
	s := make([]string,0,len(i))
	for _,o := range i {
		s = append(s,strconv.Itoa(o))
	}
	return strings.Join(s,",")
}
//int64 change string
func Int64ArrToInString(i64 []int64) string {
	s := make([]string,0,len(i64))
	for _,o := range i64 {
		s = append(s,strconv.FormatInt(o,10))
	}
	return strings.Join(s,",")
}
//string arr change int
func StringArrToInString(s []string) string {
	return `"`+strings.Join(s,`","`)+`"`
}
//判断元素是否包含
func InArray(s string, arr []string) bool {
	for _, val :=range arr {
		if s == val {
			return true
		}
	}
	return false
}

