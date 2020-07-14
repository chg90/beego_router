package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type Test4Controller struct {
	beego.Controller
}

func (p *Test4Controller) Get() {
	id := p.Ctx.Input.Param(":id")
	fmt.Println(id)
	p.Data["json"] = id
	p.ServeJSON()
}