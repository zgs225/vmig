package cmd

import (
	"github.com/zgs225/vmig/core"
)

var __app *core.App

func GetVmigApp() *core.App {
	return __app
}

func SetVmigApp(app *core.App) {
	__app = app
}
