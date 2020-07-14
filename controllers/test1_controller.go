package controllers

import (
	"github.com/astaxie/beego"
)

type Test1Controller struct {
	beego.Controller
}

var persons = make(map[string]string)

func (p *Test1Controller) Post() {
	name := p.GetString("name")
	age := p.GetString("age")
	persons[name] = age
	p.Data["json"] = "ok"
	p.ServeJSON()
}

