package controllers

import (
	"github.com/astaxie/beego"
)

type PersonController struct {
	beego.Controller
}

var persons = make(map[string]string)

func (p *PersonController) Post() {
	name := p.GetString("name")
	age := p.GetString("age")
	persons[name] = age
	p.Data["json"] = "ok"
	p.ServeJSON()
}

func (p *PersonController) Get() {
	name := p.GetString("name")
	age := persons[name]
	p.Data["json"] = age
	p.ServeJSON()
}
