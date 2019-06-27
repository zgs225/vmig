package core

import (
	"fmt"
	"strings"

	"github.com/zgs225/vmig/console"
)

const (
	DEFAULT_CONFIG_FILE = ".vmig"
)

var supportedDrivers = []string{"mysql", "postgres"}

// Config
type Config struct {
	Current   *CurrentConfig             `mapstructure:"current"`
	Databases map[string]*DatabaseConfig `mapstructure:"databases"`

	dirty bool // Whether config changed
}

// CurrentConfig 设置 vmig 默认使用的环境和版本
type CurrentConfig struct {
	Env     string `mapstructure:"env"`
	Version string `mapstructure:"version"`
}

type DatabaseConfig struct {
	Driver   string            `mapstructure:"driver"`
	DBHost   string            `mapstructure:"dbhost"`
	DBPort   int               `mapstructure:"dbport"`
	DBName   string            `mapstructure:"dbname"`
	User     string            `mapstructure:"user"`
	Password string            `mapstructure:"password"`
	Extra    map[string]string `mapstructure:"extra"`
}

// AddDatabaseConfig 添加一条数据库配置信息
func (c *Config) AddDatabaseConfig(env string, dc *DatabaseConfig, isDefault bool) error {
	if c.Databases == nil {
		c.Databases = make(map[string]*DatabaseConfig)
	}
	if c.Current == nil {
		c.Current = &CurrentConfig{}
	}
	_, ok := c.Databases[env]
	if ok {
		return fmt.Errorf("Duplicated environment: %s", env)
	}
	c.Databases[env] = dc
	if isDefault {
		c.Current.Env = env
	}
	c.dirty = true
	return nil
}

func (c *Config) SetDefaultVersion(v string) {
	if c.Current == nil {
		c.Current = &CurrentConfig{}
	}
	if c.Current.Version != v {
		c.Current.Version = v
		c.dirty = true
	}
}

func (c *Config) IsDirty() bool {
	return c.dirty
}

func (c *Config) GetCurrentDatabaseURL() (u string, err error) {
	var (
		g  DatabaseURLGenerator
		d  *DatabaseConfig
		ok bool
	)

	if d, ok = c.Databases[c.Current.Env]; !ok {
		err = fmt.Errorf("Database config for env %s not exists", c.Current.Env)
		return
	}

	switch d.Driver {
	case supportedDrivers[0]:
		g = &MysqlURLGenerator{}
		break
	default:
		err = fmt.Errorf("Unsupported driver: %s", d.Driver)
		return
	}
	u = g.Generate(d)
	return
}

// InitOption 用于初始化命令保存的参数
type InitOption struct {
	Env        string
	IsDefault  bool
	DBDriver   string
	DBHost     string
	DBPort     int
	DBName     string
	DBUser     string
	DBPassword string
}

// ReadFromConsole 从命令行中读取值
func (o *InitOption) ReadFromConsole() error {
	console.ReadStringVar(&o.Env, "development", "Please enter environment name:")
	console.ReadBoolVar(&o.IsDefault, false, "Set to default?")
	console.ReadStringVar(&o.DBDriver, "mysql", fmt.Sprintf("Please enter database driver (%s):", strings.Join(supportedDrivers, "/")))

	if err := o.checkDBDriver(); err != nil {
		return err
	}

	console.ReadStringVar(&o.DBHost, "localhost", "Please enter database host:")
	console.ReadIntVar(&o.DBPort, 3306, "Please enter database port:")
	console.ReadStringVar(&o.DBName, "my_database", "Please enter database name:")
	console.ReadStringVar(&o.DBUser, "", "Please enter database user:")
	console.ReadStringVar(&o.DBPassword, "", "Please enter database password:")

	return nil
}

func (o *InitOption) checkDBDriver() error {
	ok := false
	for _, d := range supportedDrivers {
		if o.DBDriver == d {
			ok = true
			break
		}
	}

	if !ok {
		return fmt.Errorf("Unsupported driver %s, supports (%s)", o.DBDriver, strings.Join(supportedDrivers, "/"))
	}

	return nil
}
