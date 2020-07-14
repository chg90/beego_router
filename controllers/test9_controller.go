package controllers

import (
	"github.com/astaxie/beego"
)

type Test9Controller struct {
	beego.Controller
}

func (p *Test9Controller) URLMapping() {
	p.Mapping("method9", p.TestMethod9)
}

func (p *Test9Controller) TestMethod9() {
	p.Data["json"] = "method9"
	p.ServeJSON()
}
