package chain

import (
	"github.com/hyperledger/fabric/sdk/Go/rest/models/util"

	"github.com/hyperledger/fabric/sdk/Go/rest/models/user"

	"github.com/astaxie/beego/logs"
	"github.com/hyperledger/fabric/sdk/Go/api"
)

func (chainImpl *ChainImpl) ChainDeploy(req *util.DeployRQ) (string, error) {
	userImpl := &user.UserImpl{}
	mbsrvc := api.InitMbImpl()
	err := userImpl.Login(req.UserID, req.UserToken)
	if err != nil {
		logs.Error(err)
		return "null", err
	}
	resp, err := mbsrvc.Deploy(req.Path, req.Args, []byte(req.Meta), req.UserID)
	if err != nil {
		logs.Error(err)
	}
	return string(resp.Msg), err
}
