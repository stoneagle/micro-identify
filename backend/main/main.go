package main

import (
	"identify/backend/bootstrap"
	"identify/backend/bootstrap/database"
	"identify/backend/bootstrap/image"
	"identify/backend/bootstrap/route"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("identify", "wuzhongyang@wzy.com")
	app.Bootstrap()
	app.Configure(database.Configure, route.Configure, image.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")

}
