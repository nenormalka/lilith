package methods

import (
	"fmt"
	"reflect"
)

func ArrayStructToArrayValuesByField[S ~[]T, T, M any](s S, fields []string) []M {
	if len(s) == 0 || len(fields) == 0 {
		return nil
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	sl := make([]M, 0, len(s))

	for _, elem := range s {
		recursiveArrayToArrayValues(fields, &sl, len(fields), 0, elem, elem)
	}

	return sl
}

func recursiveArrayToArrayValues[T, M any](fields []string, sl *[]M, end, current int, e any, elem T) {
	for i, field := range fields {
		v := reflect.ValueOf(e)
		if !v.IsValid() {
			break
		}

		switch v.Kind() {
		case reflect.Slice:
			for j := 0; j < v.Len(); j++ {
				recursiveArrayToArrayValues(fields[i:], sl, end, current, v.Index(j).Interface(), elem)
			}
		case reflect.Map:
			for _, key := range v.MapKeys() {
				recursiveArrayToArrayValues(fields[i:], sl, end, current, v.MapIndex(key).Interface(), elem)
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

			val, ok := f.Interface().(M)
			if ok {
				*sl = append(*sl, val)
			}
		}
	}
}
