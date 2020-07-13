package routers

import (
	"github.com/astaxie/beego/context"
	"github.com/beego_router/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//基础路由
	beego.Get("/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello get"))
	})

	beego.Post("/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello post"))
	})
	beego.Any("/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello any"))
	})

	//固定路由
	beego.Router("/v1/person", &controllers.PersonController{})

	//正则路由


	//

}
