// Code generated by sqlc. DO NOT EDIT.

package gen

import (
	"strings"
)

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

//authors.sql   DeleteAuthorIn
func int32Slice2interface(l []int32) []interface{} {
	v := make([]interface{}, len(l))
	for i, val := range l {
		v[i] = val

	}
	return v
}

//authors.sql   GetAuthorsInCompany
func stringSlice2interface(l []string) []interface{} {
	v := make([]interface{}, len(l))
	for i, val := range l {
		v[i] = val

	}
	return v
}
