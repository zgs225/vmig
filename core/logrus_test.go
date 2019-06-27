package core

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestStructToFields(t *testing.T) {
	person := struct {
		Name string
		Age  int
		Edu  struct {
			School string
			Major  string
		}
		Relationships map[string]string
	}{
		Name: "ZhangSan",
		Age:  18,
		Edu: struct {
			School string
			Major  string
		}{
			School: "THU",
			Major:  "IEEE",
		},
		Relationships: map[string]string{
			"Mother": "Hui Lan",
			"Father": "Zhang Mang",
		},
	}

	v, err := StructToFields(person)

	if err != nil {
		t.Error(err)
	}

	t.Log("Fields is:", v)

	name := v["Name"].(string)
	age := v["Age"].(int)
	edu := v["Edu"].(logrus.Fields)
	school := edu["School"].(string)
	major := edu["Major"].(string)

	if name != person.Name || age != person.Age || school != person.Edu.School || major != person.Edu.Major {
		t.Error("StructToFields value error:", v)
	}
}

func TestStructToFieldsWithUnexportableField(t *testing.T) {
	person := struct {
		NickName string
		Age      int
		realName string
	}{
		NickName: "Cat",
		Age:      18,
		realName: "ZhangSan",
	}

	v, err := StructToFields(person)

	if err != nil {
		t.Error(err)
	}

	t.Log("Fields is:", v)

	name := v["NickName"].(string)
	age := v["Age"].(int)
	_, ok := v["realName"]

	if name != person.NickName || age != person.Age {
		t.Error("StructToFields value error:", v)
	}

	if ok {
		t.Error("Unexported field should not exists.")
	}
}
