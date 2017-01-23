package chain

import (
	"github.com/astaxie/beego/logs"
	"github.com/hyperledger/fabric/sdk/Go/api"
)

func (chainImpl *ChainImpl) GetBlocksCount() (string, error) {
	mbsrvc := api.InitMbImpl()
	resp, err := mbsrvc.GetBlocksCount()
	if err != nil {
		logs.Error(err)
	}
	return resp, err
}
