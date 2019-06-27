package core

import (
	"testing"
)

func TestSet(t *testing.T) {
	c := &Config{
		Current: &CurrentConfig{
			Env:     "development",
			Version: "v1.0.0",
		},
		Databases: map[string]*DatabaseConfig{
			"development": &DatabaseConfig{
				DBPort: 3306,
			},
		},
	}

	c.Set("current.env", "production")
	c.Set("databases.development.dbport", "3307")

	if c.Current.Env != "production" {
		t.Errorf("Config set error: expected %s, got %s", "production", c.Current.Env)
	}

	if c.Databases["development"].DBPort != 3307 {
		t.Errorf("Config set error: expected %v, got %v", 3307, c.Databases["development"].DBPort)
	}
}
