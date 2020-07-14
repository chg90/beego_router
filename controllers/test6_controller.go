package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type Test6Controller struct {
	beego.Controller
}

func (p *Test6Controller) TestMethod() {
	mp := p.Ctx.Input.Params()
	id := mp["0"]
	name := mp["1"]
	fmt.Printf("0: %v,1: %v \n", id, name)
	p.Data["json"] = id
	p.ServeJSON()
}
