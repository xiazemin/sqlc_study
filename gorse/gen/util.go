// Code generated by sqlc. DO NOT EDIT.

package gorse

import (
	"reflect"
	"regexp"
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

var (
	fullValues   = regexp.MustCompile("(VALUES|values|VALUE|value)[\\t\\s\\n]*\\([\\t\\s\\n]*(\\?|[\\w\\W]+)[\\t\\s\\n]*(,[\\t\\s\\n]*(\\?|[\\w\\W]+))*\\)[\\t\\s\\n]*(,[\\t\\s\\n]*\\((\\?|[\\w\\W]+)[\\t\\s\\n]*(,[\\t\\s\\n]*(\\?|[\\w\\W]+))*\\))*")
	singleValues = regexp.MustCompile("\\((\\?|[\\w\\W]+)(,(\\?|[\\w\\W]+))*\\)")
)

func repeatN(val string, n int) string {
	str := fullValues.FindString(strings.Replace(strings.Replace(strings.Replace(val, "\t", "", -1), "\n", "", -1), " ", "", -1))
	oneStr := singleValues.FindString(str)
	values := strings.Split(oneStr, "),(")
	if len(values) > 0 {
		oneStr = "(" + values[0][1:] + ")"
	}

	var oneStrArr []string
	for i := 0; i < n; i++ {
		oneStrArr = append(oneStrArr, oneStr)
	}
	return fullValues.ReplaceAllString(val, "VALUES"+strings.Join(oneStrArr, ","))
}

type IsZeroer interface {
	IsEmpty() bool
}

var IsZeroType = reflect.TypeOf((*IsZeroer)(nil)).Elem()

// IsEmpty reports whether a value is a zero value
// Including support: Bool, Array, String, Float32, Float64, Int, Int8, Int16, Int32, Int64, Uint, Uint8, Uint16, Uint32, Uint64, Uintptr
// Map, Slice, Interface, Struct
func IsEmpty(v reflect.Value) bool {
	if v.IsValid() && v.Type().Implements(IsZeroType) {
		return v.Interface().(IsZeroer).IsEmpty()
	}
	switch v.Kind() {
	case reflect.Bool:
		return !v.Bool()
	case reflect.Array, reflect.String:
		return v.Len() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Map, reflect.Slice:
		return v.IsNil() || v.Len() == 0
	case reflect.Interface:
		return v.IsNil()
	case reflect.Invalid:
		return true
	}

	if v.Kind() != reflect.Struct {
		return false
	}

	// Traverse the Struct and only return true
	// if all of its fields return IsEmpty == true
	n := v.NumField()
	for i := 0; i < n; i++ {
		vf := v.Field(i)
		if !IsEmpty(vf) {
			return false
		}
	}
	return true
}

//items.sql   ListItemsByID
func stringSlice2interface(l []string) []interface{} {
	v := make([]interface{}, len(l))
	for i, val := range l {
		v[i] = val

	}
	return v
}

//if len(IDs)>65536,mysql will return error
func BatchRunstring(batch int, IDs []string, fn func([]string) error) error {
	for i := 0; i <= len(IDs)/batch; i++ {
		l := i * batch
		r := (i + 1) * batch
		if r > len(IDs) {
			r = len(IDs)
		}
		if r > l {
			batchIDs := IDs[l:r]
			if err := fn(batchIDs); err != nil {
				return err
			}
		}
	}
	return nil
}

//measurements.sql   ListMeasurements
func float64Slice2interface(l []float64) []interface{} {
	v := make([]interface{}, len(l))
	for i, val := range l {
		v[i] = val

	}
	return v
}

//if len(IDs)>65536,mysql will return error
func BatchRunfloat64(batch int, IDs []float64, fn func([]float64) error) error {
	for i := 0; i <= len(IDs)/batch; i++ {
		l := i * batch
		r := (i + 1) * batch
		if r > len(IDs) {
			r = len(IDs)
		}
		if r > l {
			batchIDs := IDs[l:r]
			if err := fn(batchIDs); err != nil {
				return err
			}
		}
	}
	return nil
}
