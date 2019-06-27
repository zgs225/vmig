package core

import (
	"testing"
)

func TestMysqlURLGenerator(t *testing.T) {
	c := &DatabaseConfig{
		Driver:   "mysql",
		DBHost:   "localhost",
		DBPort:   3306,
		DBName:   "project_x",
		User:     "root",
		Password: "public",
	}
	u := "mysql://root:public@tcp(localhost:3306)/project_x"
	g := &MysqlURLGenerator{}

	v := g.Generate(c)

	if v != u {
		t.Errorf("Error. Expect %s, got %s", u, v)
	}
}
