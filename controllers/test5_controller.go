package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type Test5Controller struct {
	beego.Controller
}

func (p *Test5Controller) TestMethod() {
	id := p.GetString("id")
	fmt.Println(id)
	p.Data["json"] = id
	p.ServeJSON()
}
