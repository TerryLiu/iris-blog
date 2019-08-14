package main

import (
	"iris-blog/bootstrap"
	"iris-blog/web/router"
)

func NewApp(code string) *bootstrap.BootStrapper {
	app := bootstrap.NewBootStrapper(code)
	app.Setup()
	app.Configure(router.Configure)
	return app
}

func main() {
	app := NewApp("")
	app.Listen(":9090")

}
