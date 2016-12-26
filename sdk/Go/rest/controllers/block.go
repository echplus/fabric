package controllers

import (
	"runtime/debug"

	"github.com/hyperledger/fabric/sdk/Go/rest/models/chain"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type BlockController struct {
	beego.Controller
}

func (this *BlockController) Get() {
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err, string(debug.Stack()))
			this.Ctx.WriteString(string(debug.Stack()))
		}
	}()
	id := this.Ctx.Input.Param(":id")
	chainImpl := &chain.ChainImpl{}
	resp, err := chainImpl.GetBlock(id)
	if err != nil {
		logs.Error(err)
	}
	this.Ctx.WriteString(resp)
}
