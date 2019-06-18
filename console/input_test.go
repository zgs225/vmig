package console

import (
	"testing"
)

func TestReadStringVar(t *testing.T) {
	var env string

	if err := ReadStringVar(&env, "development", "Please enter environment name:"); err != nil {
		t.Fatal(err)
	}

	if env != "development" {
		t.Errorf("env value from ReadStringVar with empty input shoud be default value, expect: %s, got: %s", "development", env)
	}

	if err := ReadStringVar(&env, "development", "Please enter environment name: production"); err != nil {
		t.Fatal(err)
	}

	if env != "production" {
		t.Errorf("env value from ReadStringVar with empty input shoud be default value, expect: %s, got: %s", "production", env)
	}
}

func TestReadBoolVar(t *testing.T) {
	var v bool

	ReadBoolVar(&v, true, "Test ReadBoolVar with empty input")
	if !v {
		t.Errorf("Value from ReadBoolVar expect %v, got %v", true, v)
	}

	ReadBoolVar(&v, true, "Test ReadBoolVar with N")
	if v {
		t.Errorf("Value from ReadBoolVar expect %v, got %v", false, v)
	}
}

func TestReadIntVar(t *testing.T) {
	var v int

	ReadIntVar(&v, 3306, "Test ReadBoolVar with empty input")
	if v != 3306 {
		t.Errorf("Value from ReadBoolVar expect %v, got %v", 3306, v)
	}

	ReadIntVar(&v, 3306, "Test ReadBoolVar with 27017")
	if v != 27017 {
		t.Errorf("Value from ReadBoolVar expect %v, got %v", 27017, v)
	}
}
