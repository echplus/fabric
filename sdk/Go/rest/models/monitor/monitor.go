package monitor

import (
	"github.com/astaxie/beego/logs"
	"github.com/hyperledger/fabric/sdk/Go/api"
)

type MonitorImpl struct {
}

func (monitorImpl *MonitorImpl) GetNetwork() (string, error) {
	mbsrvc := api.InitMbImpl()
	resp, err := mbsrvc.Network()
	if err != nil {
		logs.Error(err)
	}
	logs.Debug("GetNetwork success")
	return resp, err
}
