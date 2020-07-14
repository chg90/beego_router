package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/beego_router/controllers:Test7Controller"] = append(beego.GlobalControllerRouter["github.com/beego_router/controllers:Test7Controller"],
        beego.ControllerComments{
            Method: "TestMethod1",
            Router: `/method1/:key`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/beego_router/controllers:Test7Controller"] = append(beego.GlobalControllerRouter["github.com/beego_router/controllers:Test7Controller"],
        beego.ControllerComments{
            Method: "TestMethod2",
            Router: `/method2/:key`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
