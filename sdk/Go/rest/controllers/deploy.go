package controllers

import (
	"encoding/json"
	"runtime/debug"

	"github.com/hyperledger/fabric/sdk/Go/rest/models/chain"
	"github.com/hyperledger/fabric/sdk/Go/rest/models/util"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type DeployController struct {
	beego.Controller
}

func (this *DeployController) Post() {
	var ob util.DeployRQ
	var resp util.DeployRS
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err, string(debug.Stack()))
			resp.ErrorCode = 10001
			resp.ErrorMsg = string(debug.Stack())
			resp.ChaincodeID = "null"
		}
	}()
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	if err != nil {
		logs.Error(err)
	}
	logs.Debug("-----", string(this.Ctx.Input.RequestBody))
	logs.Debug("------", ob)
	chainImpl := &chain.ChainImpl{}
	msg, err := chainImpl.ChainDeploy(&ob)
	resp.ChaincodeID = msg
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
