package controllers

import (
	"encoding/json"
	"runtime/debug"

	"github.com/hyperledger/fabric/sdk/Go/rest/models/chain"
	"github.com/hyperledger/fabric/sdk/Go/rest/models/util"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type InvokeController struct {
	beego.Controller
}

func (this *InvokeController) Post() {
	var ob util.InvokeRQ
	var resp util.InvokeRS
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err, string(debug.Stack()))
			resp.ErrorCode = 10001
			resp.ErrorMsg = string(debug.Stack())
			resp.Msg = "null"
		}
	}()
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	if err != nil {
		logs.Error(err)
	}
	chainImpl := &chain.ChainImpl{}
	msg, err := chainImpl.ChainInvoke(&ob)
	resp.Msg = msg
	if err != nil {
		logs.Error(err)
		resp.ErrorCode = 10001
		resp.ErrorMsg = err.Error()
	} else {
		resp.ErrorCode = 10000
		resp.ErrorMsg = "success"
	}
	jsRe, _ := json.Marshal(&resp)
	this.Ctx.WriteString(string(jsRe))
}
