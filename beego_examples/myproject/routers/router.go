package routers

import (
	"myproject/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.RegisterController{})

	// basic router examples [GET, POST, ANY]
	beego.Get("/hello", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello world!"))
	})
	beego.Post("/alice", func(ctx *context.Context) {
		ctx.Output.Body([]byte("bob"))
	})
	beego.Any("/foo", func(ctx *context.Context) {
		ctx.Output.Body([]byte("bar"))
	})
}
