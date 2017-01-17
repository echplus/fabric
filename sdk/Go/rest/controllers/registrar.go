package controllers

import (
	"encoding/json"

	"github.com/hyperledger/fabric/sdk/Go/rest/models/user"
	"github.com/hyperledger/fabric/sdk/Go/rest/models/util"

	"runtime/debug"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type RegistrarController struct {
	beego.Controller
}

func (this *RegistrarController) Post() {
	var ob util.RegisterRQ
	var resp util.RegisterRS
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err, string(debug.Stack()))
			resp.ErrorCode = 10001
			resp.ErrorMsg = string(debug.Stack())
			resp.Token = "null"

		}
	}()
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	if err != nil {
		logs.Error(err)
	}
	userImpl := &user.UserImpl{}
	token, err := userImpl.Registrar(&ob)
	resp.Token = token
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
