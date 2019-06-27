package core

import (
	"strconv"
	"strings"
)

// DatabaseURLGenerator 通过 DatabaseConfig 生成数据库URL
type DatabaseURLGenerator interface {
	Generate(*DatabaseConfig) string
}

type MysqlURLGenerator struct {
}

func (g *MysqlURLGenerator) Generate(c *DatabaseConfig) string {
	buf := []string{"mysql://"}

	if len(c.User) > 0 {
		buf = append(buf, c.User)

		if len(c.Password) > 0 {
			buf = append(buf, ":", c.Password)
		}
	}

	buf = append(buf, "@tcp(")
	buf = append(buf, c.DBHost)
	if c.DBPort > 0 {
		buf = append(buf, ":", strconv.FormatInt(int64(c.DBPort), 10))
	}
	buf = append(buf, ")")
	if len(c.DBName) > 0 {
		buf = append(buf, "/", c.DBName)
	}
	if c.Extra != nil {
		first := true
		for k, v := range c.Extra {
			if first {
				buf = append(buf, "?")
				first = false
			} else {
				buf = append(buf, "&")
			}
			buf = append(buf, k, "=", v)
		}

	}

	return strings.Join(buf, "")
}
