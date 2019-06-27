package core

import (
	"errors"
	"reflect"

	"github.com/sirupsen/logrus"
)

// StructToFields 将结构体转换成 logrus.Fields 对象
func StructToFields(i interface{}) (logrus.Fields, error) {
	return StructToFields2(i)
}

func StructToFields2(i interface{}) (logrus.Fields, error) {
	v := reflect.Indirect(reflect.ValueOf(i))

	switch v.Kind() {
	case reflect.Struct:
		rt := logrus.Fields{}
		t := v.Type()

		for i := 0; i < t.NumField(); i++ {
			fd := t.Field(i)
			fv := reflect.Indirect(v.Field(i))
			if fv.IsValid() && fv.CanInterface() {
				switch fv.Kind() {
				case reflect.Struct, reflect.Map:
					fvv, err := StructToFields2(fv.Interface())
					if err != nil {
						return nil, err
					}
					rt[fd.Name] = fvv
					break
				default:
					rt[fd.Name] = fv.Interface()
					break
				}
			}
		}

		return rt, nil
		break
	case reflect.Map:
		rt := logrus.Fields{}
		iter := v.MapRange()

		for iter.Next() {
			mk := iter.Key()
			mv := reflect.Indirect(iter.Value())

			if mk.Kind() == reflect.String && mv.IsValid() && mv.CanInterface() {
				switch mv.Kind() {
				case reflect.Struct, reflect.Map:
					mvv, err := StructToFields2(mv.Interface())
					if err != nil {
						return nil, err
					}
					rt[mk.String()] = mvv
					break
				default:
					rt[mk.String()] = mv.Interface()
					break
				}
			}
		}
		return rt, nil
		break
	}
	return nil, errors.New("Invalid argument type: struct or map required.")
}
