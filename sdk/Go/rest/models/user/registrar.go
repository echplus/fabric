package user

import (
	"errors"

	"github.com/hyperledger/fabric/sdk/Go/rest/models/util"

	"github.com/astaxie/beego/logs"
	"github.com/hyperledger/fabric/sdk/Go/api"
	"github.com/hyperledger/fabric/sdk/Go/db"
	"github.com/hyperledger/fabric/sdk/Go/utils"
)

func (userImpl *UserImpl) Registrar(req *util.RegisterRQ) (string, error) {
	err := userImpl.Login(req.RegistarId, req.RegistarToken)
	if err != nil {
		logs.Error(err)
		return "null", err
	}
	sqlDB := &db.SqlDB{}
	if sqlDB == nil {
		logs.Error("Null pointer deference!")
		return "null", errors.New("Null pointer deference!")
	}

	dbuser := &db.UsersInfo{}
	err = sqlDB.ReadUser(dbuser, req.RegistarId)
	if err != nil {
		logs.Error(err)
		return "null", err
	}

	priv, err := utils.PrivStringToPointer(dbuser.Priv)
	if priv == nil {
		logs.Error("Null pointer deference!")
		return "null", errors.New("Null pointer deference!")
	}

	if err != nil {
		logs.Error(err)
		return "null", err
	}

	//注册用户
	registrar := api.User{EnrollID: req.RegistarId, EnrollPwd: []byte(req.RegistarToken), EnrollPrivKey: priv}
	userInfo := &api.User{EnrollID: req.UserID, Role: req.Role, Affiliation: req.Affiliation}
	mbsrvc := api.InitMbImpl()
	err = mbsrvc.Registrar(registrar, userInfo)
	if err != nil {
		logs.Error(err)
		return "null", err
	}

	token := string(userInfo.EnrollPwd)
	err = userImpl.Login(userInfo.EnrollID, token)
	if err != nil {
		logs.Error(err)
		return token, err
	}
	return token, nil
}
