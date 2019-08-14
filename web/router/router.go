package router

import (
	"github.com/kataras/iris/mvc"
	"iris-blog/bootstrap"
	"iris-blog/web/controller"
)

func Configure(b *bootstrap.BootStrapper) {
	r := mvc.New(b.Party("/"))

	example := r.Party("/example")
	example.Handle(new(controller.ExampleController))

}
