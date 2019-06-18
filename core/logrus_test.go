package core

import (
	"testing"
)

func TestStructToFields(t *testing.T) {
	person := struct {
		Name string
		Age  int
	}{
		Name: "ZhangSan",
		Age:  18,
	}

	v, err := StructToFields(person)

	if err != nil {
		t.Error(err)
	}

	t.Log("Fields is:", v)

	name := v["Name"].(string)
	age := v["Age"].(int)

	if name != person.Name || age != person.Age {
		t.Error("StructToFields value error:", v)
	}
}
