package controllers

import (
	"github.com/astaxie/beego"
)

type Test8Controller struct {
	beego.Controller
}

func (p *Test8Controller) URLMapping() {
	p.Mapping("test8", p.TestMethod8)
}

func (p *Test8Controller) TestMethod8() {
	p.Data["json"] = "method8"
	p.ServeJSON()
}
