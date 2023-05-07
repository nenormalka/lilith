package methods

import (
	"fmt"
	"reflect"
)

func ArrayStructToMapByField[S ~[]T, T any, O comparable](s S, fields []string) map[O]T {
	if len(s) == 0 || len(fields) == 0 {
		return nil
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	m := make(map[O]T, len(s))

	for _, elem := range s {
		recursiveArrayToMap(fields, m, len(fields), 0, elem, elem)
	}

	return m
}

func recursiveArrayToMap[T any, O comparable](fields []string, m map[O]T, end, current int, e any, elem T) {
	for i, field := range fields {
		v := reflect.ValueOf(e)
		if !v.IsValid() {
			break
		}

		switch v.Kind() {
		case reflect.Slice:
			for j := 0; j < v.Len(); j++ {
				recursiveArrayToMap(fields[i:], m, end, current, v.Index(j).Interface(), elem)
			}
		case reflect.Map:
			for _, key := range v.MapKeys() {
				recursiveArrayToMap(fields[i:], m, end, current, v.MapIndex(key).Interface(), elem)
			}
		default:
			current++

			isPtr := v.Kind() == reflect.Ptr

			if isPtr && v.IsNil() {
				break
			}

			if isPtr {
				v = v.Elem()
			}

			f := v.FieldByName(field)
			if !f.IsValid() {
				break
			}

			isPtr = f.Kind() == reflect.Ptr
			if isPtr && f.IsNil() {
				break
			}

			if isPtr {
				f = f.Elem()
			}

			if current != end {
				e = f.Interface()
				continue
			}

			val, ok := f.Interface().(O)
			if ok {
				m[val] = elem
			}
		}
	}
}
