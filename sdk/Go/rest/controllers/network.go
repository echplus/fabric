package controllers

import (
	"runtime/debug"

	"github.com/hyperledger/fabric/sdk/Go/rest/models/monitor"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type NetworkController struct {
	beego.Controller
}

func (this *NetworkController) Get() {
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err, string(debug.Stack()))
			this.Ctx.WriteString(string(debug.Stack()))
		}
	}()
	monitorImpl := &monitor.MonitorImpl{}
	resp, err := monitorImpl.GetNetwork()
	if err != nil {
		logs.Error(err)
	}
	this.Ctx.WriteString(resp)
	this.TplName = "network.tpl"
}
