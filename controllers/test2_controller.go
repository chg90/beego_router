package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type Test2Controller struct {
	beego.Controller
}

func (p *Test2Controller) Get() {
	id := p.Ctx.Input.Param(":id")
	fmt.Println(id)
	p.Data["json"] = id
	p.ServeJSON()
}
