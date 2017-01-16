package chain

import (
	"github.com/astaxie/beego/logs"
	"github.com/hyperledger/fabric/sdk/Go/api"
)

func (chainImpl *ChainImpl) GetTransaction(id string) (string, error) {
	mbsrvc := api.InitMbImpl()
	resp, err := mbsrvc.GetTransactions(id)
	if err != nil {
		logs.Error(err)
	}
	return resp, err
}
