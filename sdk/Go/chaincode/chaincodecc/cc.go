package main

import (
	"errors"
	"os"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/op/go-logging"
)

type CC struct {
}

var logger = logging.MustGetLogger("dummy cc")

func main() {

	go func() {
		time.Sleep(time.Second * 1)

		format := logging.MustStringFormatter(`[%{module}] %{time:2006-01-02 15:04:05} [%{level}] [%{longpkg} %{shortfile}] { %{message} }`)

		backendConsole := logging.NewLogBackend(os.Stderr, "", 0)
		backendConsole2Formatter := logging.NewBackendFormatter(backendConsole, format)

		logging.SetBackend(backendConsole2Formatter)
		logging.SetLevel(logging.INFO, "utxo")
	}()

	if err := shim.Start(new(CC)); err != nil {
		logger.Fatal(err)
	}
}

func (c *CC) Init(stub shim.ChaincodeStubInterface, function string, args []string) (resp []byte, e error) {

	defer func() {
		if err := recover(); err != nil {
			logger.Error("deploy panic", err)
			switch err.(type) {
			case string:
				e = errors.New(err.(string))
			case error:
				e = err.(error)
			}
		}
	}()

	logger.Info("deploy function:", function)
	logger.Info("deploy args:", args)

	return
}

func (c *CC) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) (resp []byte, e error) {

	defer func() {
		if err := recover(); err != nil {
			logger.Error("invoke panic", err)
			switch err.(type) {
			case string:
				e = errors.New(err.(string))
			case error:
				e = err.(error)
			}
		}
	}()

	logger.Info("invoke function:", function)
	logger.Info("invoke args:", args)

	if len(args) != 2 {
		e = errors.New("args must contains two parts")
		logger.Error(e)
		return
	}

	if e = stub.PutState(args[0], []byte(args[1])); e != nil {
		logger.Error(e)
		return
	}

	return
}

func (c *CC) Query(stub shim.ChaincodeStubInterface, function string, args []string) (resp []byte, e error) {

	defer func() {
		if err := recover(); err != nil {
			logger.Error("query panic", err)
			switch err.(type) {
			case string:
				e = errors.New(err.(string))
			case error:
				e = err.(error)
			}
		}
	}()

	logger.Info("query function:", function)
	logger.Info("query args:", args)

	if len(args) != 1 {
		e = errors.New("args must contains one part")
		logger.Error(e)
		return
	}

	resp, e = stub.GetState(args[0])
	return
}
