package user

import (
	"errors"

	"github.com/hyperledger/fabric/sdk/Go/api"

	"github.com/astaxie/beego/logs"
	"github.com/hyperledger/fabric/sdk/Go/db"
)

type UserImpl struct {
}

func (userImpl *UserImpl) readuser(userID string, pwd string) error {
	dbuser := &db.UsersInfo{}
	sqlDB := &db.SqlDB{}
	if sqlDB == nil {
		logs.Error("Null pointer deference!")
		return errors.New("Null pointer deference!")
	}
	err := sqlDB.ReadUser(dbuser, userID)
	if err != nil {
		logs.Error("read eca errors:", err.Error)
		return err
	}

	logs.Debug("read ecaInfo success!")
	bFlag := (dbuser.UserId != userID) || (dbuser.Token != pwd)
	if bFlag {
		logs.Error("user name or token error!")
		return errors.New("user name or token error!")
	}
	return nil
}

func (userImpl *UserImpl) Login(userID string, pwd string) error {
	user := &api.User{EnrollID: userID, EnrollPwd: []byte(pwd)}
	err := userImpl.readuser(userID, pwd)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			logs.Error(err)
			return err
		}
	} else {
		return nil
	}
	mbsrvc := api.InitMbImpl()
	_, _, _, err = mbsrvc.Enroll(user)
	if err != nil {
		logs.Error(err)
		return err
	}
	return nil
}
