package core

import (
	"errors"
	"reflect"

	"github.com/sirupsen/logrus"
)

// StructToFields 将结构体转换成 logrus.Fields 对象
func StructToFields(i interface{}) (logrus.Fields, error) {
	v := reflect.Indirect(reflect.ValueOf(i))

	if v.Kind() != reflect.Struct {
		return nil, errors.New("Invalid argument type: struct required.")
	}

	r := logrus.Fields{}
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fv := v.Field(i)
		if fv.CanInterface() {
			fvv := reflect.Indirect(fv)
			if fvv.Kind() == reflect.Struct {
				fvv2, err := StructToFields(fvv.Interface())
				if err != nil {
					return nil, err
				}
				r[f.Name] = fvv2
			} else {
				r[f.Name] = fvv.Interface()
			}
		}
	}

	return r, nil
}
