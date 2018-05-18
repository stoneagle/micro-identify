package main

import (
	"identify/backend/bootstrap"
	"identify/backend/bootstrap/database"
	"identify/backend/bootstrap/ipc"
	"identify/backend/bootstrap/route"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("identify", "wuzhongyang@wzy.com")
	app.Bootstrap()
	app.Configure(database.Configure, route.Configure, ipc.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")
}
