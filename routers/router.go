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
	beego.Router("/v1/person", &controllers.Test1Controller{})

	//正则路由
	//  /api_2/123可以匹配成功，变量”:id”值为”123”
	beego.Router("/api_2/?:id", &controllers.Test2Controller{})
	//  /api_3/123可以匹配成功，变量”:id”值为”123”，不匹配/api_3/
	beego.Router("/api_3/:id", &controllers.Test3Controller{})
	//  /api_4/123,变量”:id”值为”123”,不匹配/api_4/
	beego.Router("/api_4/:id([0-9]+)", &controllers.Test4Controller{})
	//正则字符串匹配,/user/astaxie,变量":username"值为astaxie
	//beego.Router("/user/:username([\\w]+)", &controllers.Test2Controller{})
	//*匹配方式,/download/file/api.xml,变量":path"值为"file/api",变量":ext"值为"xml"
	//beego.Router("/download/*.*", &controllers.Test2Controller{})
	//*全匹配方式,/download/ceshi/file/api.json,变量":splat"值为"file/api.json"
	//beego.Router("/download/ceshi/*", &controllers.Test2Controller{})
	//int 类型设置方式,":id"为int类型,框架帮你实现了正则([0-9]+)
	//beego.Router("/:id:int", &controllers.Test2Controller{})
	//string类型设置方式,":hi"为string类型,框架帮你实现了正则([\w]+)
	//beego.Router("/:hi:string", &controllers.Test2Controller{})
	//带有前缀的自定义正则,":id"为正则类型。匹配 cms_123.html这样的url,":id"值为"123"
	//beego.Router("/cms_:id([0-9]+).html", &controllers.Test2Controller{})

	//自定义方法及RESTful规则，第三个参数用来设置对应的请求method到方法名，*表示通配
	beego.Router("/test5", &controllers.Test5Controller{}, "*:TestMethod")

	//自动匹配
	beego.AutoRouter(&controllers.Test6Controller{})

	//注解,自动生成commentsRouter_.go
	beego.Include(&controllers.Test7Controller{})

	//namespace
	ns1 := beego.NewNamespace("/v1",
		beego.NSNamespace("/version",
			beego.NSInclude(
				&controllers.Test8Controller{},
			),
		),
	)
	ns2 := beego.NewNamespace("/v2",
		beego.NSInclude(
			&controllers.Test9Controller{},
		),
	)
	beego.AddNamespace(ns1, ns2)

}
