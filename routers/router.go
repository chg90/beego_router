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



	//beego.Router("/user/:username([\\w]+)", &controllers.Test2Controller{})
	//
	////正则字符串匹配 //匹配 /user/astaxie :username = astaxie
	//
	//beego.Router("/download/*.*", &controllers.Test2Controller{})
	//
	////*匹配方式 //匹配 /download/file/api.xml :path= file/api :ext=xml
	//
	//beego.Router("/download/ceshi/*", &controllers.Test2Controller{})
	//
	////*全匹配方式 //匹配 /download/ceshi/file/api.json :splat=file/api.json
	//
	//beego.Router("/:id:int", &controllers.Test2Controller{})
	//
	////int 类型设置方式，匹配 :id为int类型，框架帮你实现了正则([0-9]+)
	//
	//beego.Router("/:hi:string", &controllers.Test2Controller{})
	//
	////string 类型设置方式，匹配 :hi为string类型。框架帮你实现了正则([\w]+)
	//
	//beego.Router("/cms_:id([0-9]+).html", &controllers.Test2Controller{})
	//
	////带有前缀的自定义正则 //匹配 :id为正则类型。匹配 cms_123.html这样的url :id = 123

}
