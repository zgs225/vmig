package core

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

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
	err = app.initLogger()
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

// CreateVersion 创建新版本
func (app *App) CreateVersion(version string, isDefault bool) error {
	dirName := filepath.Join(app.Root, version)
	if _, err := os.Stat(dirName); !os.IsNotExist(err) {
		return fmt.Errorf("Version already exists: " + version)
	}

	if err := os.Mkdir(dirName, 0755); err != nil {
		return err
	}

	app.Logger.WithField("directory", dirName).Debug("Version directory created.")

	if isDefault {
		app.Config.SetDefaultVersion(version)
		app.Logger.Debug("Current version set to ", version)
	}

	app.Logger.Infof("Version %s created.", version)

	return nil
}

// New 创建新的迁移文件
func (app *App) New(title, version string) error {
	if err := app.checkVersion(version); err != nil {
		return err
	}

	args := []string{"create", "-ext", "sql", "-dir", version, title}
	stdout, stderr, err := vmig_exec("migrate", args)

	if len(stdout) > 0 {
		app.Logger.Debug(stdout)
	}

	if len(stderr) > 0 {
		app.Logger.Error(stderr)
	}

	if err != nil {
		return err
	}
	app.Logger.WithField("cmd", "migrate").WithField("args", args).Debug("golang-migrate create command executed.")
	app.Logger.WithField("Version", version).Info("migration files created.")
	return nil
}

// Up 应用指定版本的所有迁移文件
func (app *App) Up(env, version string) error {
	return app.UpN(env, version, 0)
}

// Down 回滚指定版本的所有迁移文件
func (app *App) Down(env, version string) error {
	return app.DownN(env, version, 0)
}

// UpN 应用指定版本指定数量迁移文件，n <= 0 代表所有
func (app *App) UpN(env, version string, n int) error {
	if version == "" {
		return errors.New("Version required.")
	}

	if env == "" {
		return errors.New("Environment required.")
	}

	if err := app.checkVersion(version); err != nil {
		return err
	}

	cmd := "migrate"
	database, err := app.Config.GetCurrentDatabaseURL()
	if err != nil {
		return err
	}
	args := []string{"-path", version, "-database", database, "up"}
	if n > 0 {
		args = append(args, strconv.FormatInt(int64(n), 10))
	}
	stdout, stderr, err := vmig_exec(cmd, args)

	if len(stdout) > 0 {
		app.Logger.Debug(stdout)
	}

	if len(stderr) > 0 {
		p := regexp.MustCompile("\\d+\\/u\\s.+\\(.*\\)")
		m := p.FindAllString(stderr, -1)
		if m != nil {
			for _, l := range m {
				app.Logger.WithField("Env", env).WithField("Version", version).Info(l)
			}
		} else {
			app.Logger.Error(stderr)
		}
	}

	if err != nil {
		return err
	}
	app.Logger.WithField("cmd", "migrate").WithField("args", args).Debug("golang-migrate up command executed.")
	app.Logger.WithField("Env", env).WithField("Version", version).Info("Migrated.")
	return nil
}

// DownN 回滚指定版本指定数量迁移文件，n <= 0 代表所有
func (app *App) DownN(env, version string, n int) error {
	if version == "" {
		return errors.New("Version required.")
	}

	if env == "" {
		return errors.New("Environment required.")
	}

	if err := app.checkVersion(version); err != nil {
		return err
	}

	cmd := "migrate"
	database, err := app.Config.GetCurrentDatabaseURL()
	if err != nil {
		return err
	}
	args := []string{"-path", version, "-database", database, "down"}
	if n > 0 {
		args = append(args, strconv.FormatInt(int64(n), 10))
	}
	stdout, stderr, err := vmig_exec(cmd, args)

	if len(stdout) > 0 {
		app.Logger.Debug(stdout)
	}

	if len(stderr) > 0 {
		p := regexp.MustCompile("\\d+\\/d\\s.+\\(.*\\)")
		m := p.FindAllString(stderr, -1)
		if m != nil {
			for _, l := range m {
				app.Logger.WithField("Env", env).WithField("Version", version).Info(l)
			}
		} else {
			app.Logger.Error(stderr)
		}
	}

	if err != nil {
		return err
	}
	app.Logger.WithField("cmd", "migrate").WithField("args", args).Debug("golang-migrate down command executed.")
	app.Logger.WithField("Env", env).WithField("Version", version).Info("Rollback.")
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

func (app *App) checkVersion(v string) error {
	// Check exists
	dir := filepath.Join(app.Root, v)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return fmt.Errorf("Version %s does not exist.", v)
	}

	return nil
}
