package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {
	beego.GlobalControllerRouter["github.com/beego_router/controllers:Test8Controller"] = append(beego.GlobalControllerRouter["github.com/beego_router/controllers:VersionController"],
		beego.ControllerComments{
			Method:           "TestMethod8",
			Router:           ``,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})
	beego.GlobalControllerRouter["github.com/beego_router/controllers:Test9Controller"] = append(beego.GlobalControllerRouter["github.com/beego_router/controllers:TransactionForAndroidController"],
		beego.ControllerComments{
			Method:           "TestMethod9",
			Router:           `/transaction`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})
}
