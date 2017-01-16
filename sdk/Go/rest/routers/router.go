package routers

import (
	"github.com/hyperledger/fabric/sdk/Go/rest/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/user/registrar", &controllers.RegistrarController{})
	beego.Router("/chain/deploy", &controllers.DeployController{})
	beego.Router("/chain/invoke", &controllers.InvokeController{})
	beego.Router("/chain/query", &controllers.QueryController{})
	beego.Router("/monitor/network", &controllers.NetworkController{})
	beego.Router("/chain/blocks/:id", &controllers.BlockController{})
	beego.Router("/chain", &controllers.ChainController{})
	beego.Router("/chain/transactions/:id", &controllers.TransactionController{})
}
