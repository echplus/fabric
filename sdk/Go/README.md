Hyperledger Fabric Client SDK for GoLang

## 开发手册

本SDK包括三个方面进行阐述，db的脚本、配置文件、sdk导出接口

db存储脚本文件为db.sql (数据库类型为mysql)

配置文件说明

mbsrvc为连接到 [membersrvc]节点

peer为链接到peer节点

db为数据库链接信息

rest为节点提供的rest

SDK导出接口

Registrar(registrar User, user \*User) error

功能：用户注册接口

参数：
registrar 具有注册权限的用户信息（admin）

user 注册者用户信息

返回值：

错误提示


Enroll(user \*User) (interface{}, []byte, []byte, error)

   功能：用户登录接口
参数：
    user 用户信息
返回值：
用户私钥
证书签名
Pkchain（开启隐私的公钥）

错误提示


Deploy(path string, args []string, metadata []byte, userid string) (\*pb.Response, error)

功能：用户部署接口

参数：

path 部署的路径

args 参数组（第一个值为chaincode方法名称，第二个起为参数值）

metadata 原始数据

userid 为用户id

返回值

返回参数结构体

错误提示


Invoke(name string, args []string, txid string, metadata []byte, userid string) (\*pb.Response, error)

功能：执行交易接口

参数：

name chaincodeID名称

args 参数组（第一个值为chaincode方法名称，第二个起为参数值）

txid 交易id号

metadata 原始数据

userid 为用户id

返回值

返回参数结构体

错误提示


Query(name string, args []string, txid string, metadata []byte, userid string) (\*pb.Response, error)

功能：查询交易接口

参数：
name chaincodeID名称

args 参数组（第一个值为chaincode方法名称，第二个起为参数值）

txid 交易id号

metadata 原始数据

userid 为用户id

返回值

返回参数结构体

错误提示


Network() (string, error)

功能：获取网络节点信息

返回值：

网络节点信息

错误提示


GetBlocks(id string) (string, error)

功能：根据区块id获取详细的区块信息

返回值：

区块的详细信息

错误提示


GetChain() (string, error)

功能：获得区块高度等信息

返回值：

区块高度信息

错误提示


GetTransactions(txid string) (string, error)

功能：获取交易信息

参数：

txid交易ID

返回值：

交易信息

错误提示
