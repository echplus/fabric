package config

import (
	_ "github.com/hyperledger/fabric/sdk/Go/core/crypto"
)

func InitCfg() error {
	initLoging()
	err := initViper()
	if err != nil {
		return nil
	}
	initlogLevel()
	return nil
}
