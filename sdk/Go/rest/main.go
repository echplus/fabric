package main

import (
	"github.com/astaxie/beego"
	_ "github.com/hyperledger/fabric/sdk/Go/rest/routers"
)

func main() {
	beego.SetStaticPath("/static", "static")
	beego.Run()
}
