package core

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// App
type App struct {
	Root    string
	Config  *Config
	Logger  *logrus.Logger
	Verbose bool
}

// NewApp 初始化App
// 参数是项目的根目录，如果值是空字符串，则按照 pwd 值算
func NewApp(root string) (app *App, err error) {
	if root == "" {
		if root, err = os.Getwd(); err != nil {
			return
		}
	}
	app = &App{Root: root}
	if err = app.initLogger(); err != nil {
		return
	}

	return
}

// SetVerbose 打开Debug 信息输出
func (app *App) SetVerbose(v bool) {
	app.Verbose = v
	if app.Verbose {
		app.Logger.SetLevel(logrus.DebugLevel)
		app.Logger.Debugln("Vmig app output DEBUG message.")
	} else {
		app.Logger.SetLevel(logrus.InfoLevel)
	}
}

// Init 初始化
func (app *App) Init() error {
	o := &InitOption{}
	if err := o.ReadFromConsole(); err != nil {
		return err
	}

	var l *logrus.Entry

	if f, err := StructToFields(o); err == nil {
		l = app.Logger.WithFields(f)
		l.Debugln("Init options detail")
	} else {
		return err
	}

	err := app.Config.AddDatabaseConfig(o.Env, &DatabaseConfig{
		Driver:   o.DBDriver,
		DBHost:   o.DBHost,
		DBPort:   o.DBPort,
		DBName:   o.DBName,
		User:     o.DBUser,
		Password: o.DBPassword,
	}, o.IsDefault)

	if err != nil {
		return err
	}

	l.Debug("configuration added.")

	return nil
}

// LoadConfigFromViper 通过 Viper 加载配置文件
func (app *App) LoadConfigFromViper(v *viper.Viper) error {
	app.Config = &Config{}
	if err := v.Unmarshal(app.Config); err == nil {
		if f, err := StructToFields(app.Config); err == nil {
			app.Logger.WithFields(f).Debug("Config loaded.")
		} else {
			return err
		}
		return nil
	} else {
		return err
	}
}

// DumpConfigByViper 通过 Viper 将配置信息保存到配置文件中
func (app *App) DumpConfigByViper(v *viper.Viper) error {
	var fileName string

	if fileName = v.ConfigFileUsed(); fileName == "" {
		fileName = filepath.Join(app.Root, DEFAULT_CONFIG_FILE+".yaml")
	}

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		app.Logger.Debugln("Config file not exists. Creating", fileName)
		if _, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644); err != nil {
			return err
		}
	}

	v.Set("current", app.Config.Current)
	v.Set("databases", app.Config.Databases)

	return v.WriteConfig()
}

func (app *App) initLogger() error {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true,
	})
	l.SetReportCaller(false)
	l.SetLevel(logrus.InfoLevel)
	app.Logger = l
	return nil
}