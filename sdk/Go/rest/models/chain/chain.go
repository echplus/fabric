package chain

import (
	"github.com/astaxie/beego/logs"
	"github.com/hyperledger/fabric/sdk/Go/api"
)

type ChainImpl struct {
}

func (chainImpl *ChainImpl) GetChain() (string, error) {
	mbsrvc := api.InitMbImpl()
	resp, err := mbsrvc.GetChain()
	if err != nil {
		logs.Error(err)
	}
	return resp, err
}
