package core

import (
	"fmt"
	"reflect"
	"strconv"
)

// convertFuncTable 类型转换表，分别是源类型 -> 目标类型
var convertFuncTable = map[reflect.Kind]map[reflect.Kind]func(reflect.Value) (reflect.Value, error){
	reflect.String: map[reflect.Kind]func(reflect.Value) (reflect.Value, error){
		reflect.Int: func(v reflect.Value) (reflect.Value, error) {
			s := v.String()
			i, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				return reflect.Value{}, err
			}
			return reflect.ValueOf(int(i)), nil
		},
		reflect.Bool: func(v reflect.Value) (reflect.Value, error) {
			s := v.String()
			b, err := strconv.ParseBool(s)
			if err != nil {
				return reflect.Value{}, err
			}
			return reflect.ValueOf(b), nil
		},
	},
}

func ConvertValue(v reflect.Value, t reflect.Type) (reflect.Value, error) {
	sk := v.Kind()
	tk := t.Kind()
	m, ok := convertFuncTable[sk]
	if !ok {
		return reflect.Value{}, fmt.Errorf("Cannot convert type %v to %v", v.Type(), t)
	}
	f, ok := m[tk]
	if !ok {
		return reflect.Value{}, fmt.Errorf("Cannot convert type %v to %v", v.Type(), t)
	}
	return f(v)
}
