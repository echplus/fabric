package controllers

import (
	"runtime/debug"

	"github.com/hyperledger/fabric/sdk/Go/rest/models/chain"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type TransactionController struct {
	beego.Controller
}

func (this *TransactionController) Get() {
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err, string(debug.Stack()))
			this.Ctx.WriteString(string(debug.Stack()))
		}
	}()
	id := this.Ctx.Input.Param(":id")
	chainImpl := &chain.ChainImpl{}
	resp, err := chainImpl.GetTransaction(id)
	if err != nil {
		logs.Error(err)
	}
	this.Ctx.WriteString(resp)
}
