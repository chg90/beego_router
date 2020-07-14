package controllers

import (
	"github.com/astaxie/beego"
)

type Test7Controller struct {
	beego.Controller
}

func (p *Test7Controller) URLMapping() {
	p.Mapping("method1", p.TestMethod1)
	p.Mapping("method2", p.TestMethod2)
}

// @router /method1/:key [get]
func (p *Test7Controller) TestMethod1() {
	p.Data["json"] = "method1"
	p.ServeJSON()
}

// @router /method2/:key [post]
func (p *Test7Controller) TestMethod2() {
	p.Data["json"] = "method2"
	p.ServeJSON()
}
