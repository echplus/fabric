package controllers

import (
	"runtime/debug"

	"github.com/hyperledger/fabric/sdk/Go/rest/models/chain"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type ChainController struct {
	beego.Controller
}

func (this *ChainController) Get() {

	defer func() {
		if err := recover(); err != nil {
			logs.Error(err, string(debug.Stack()))
			this.Ctx.WriteString(string(debug.Stack()))
		}
	}()
	chainImpl := &chain.ChainImpl{}
	resp, err := chainImpl.GetChain()
	if err != nil {
		logs.Error(err)
	}
	this.Ctx.WriteString(resp)
}
