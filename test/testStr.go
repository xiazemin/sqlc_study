package main

import (
	"fmt"
	"strings"
)

var listAuthors="12(?) dd (?)"
func main() {
	id := []int{1, 2, 3, 4}
	param := "?"
	for i := 0; i < len(id)-1; i++ {
		param += ",?"
	}
	fmt.Println(param)
	fmt.Println(strings.Replace(listAuthors, "(?)", "("+param+")", 1))
	listAuthors = strings.Replace(listAuthors, "(?)", "("+param+")", 0)
	fmt.Println(listAuthors)
	fmt.Println(strings.Replace(listAuthors, "(?)", "("+param+")", 2))
	fmt.Println(strings.Replace(listAuthors, "(?)", "("+param+")", 3))
	fmt.Println(replaceNth(listAuthors, "(?)", "("+param+")", 2))
	name:=[]string{"name1","name2"}
	val:=3
	myPrint(id,val,name)
}

func myPrint(v ...interface{}){
    fmt.Println(v)
}

// Replace the nth occurrence of old in s by new.
func replaceNth(s, old, new string, n int) string {
	i := 0
	for m := 1; m <= n; m++ {
		x := strings.Index(s[i:], old)
		if x < 0 {
			break
		}
		i += x
		if m == n {
			return s[:i] + new + s[i+len(old):]
		}
		i += len(old)
	}
	return s
}

