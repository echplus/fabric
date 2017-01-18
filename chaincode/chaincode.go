package main

import (
	_ "encoding/base64"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/crypto/primitives"
	"github.com/op/go-logging"
)

type Institution struct {
	Name        string `json:"Name"`
	TotalNumber string `json:"TotalNumber"`
	RestNumber  string `json:"RestNumber"`
	ID          string `json:"ID"`
}

type Institutions struct {
	InstitutionIndex []Institution `json:"Institution"`
}

type User struct {
	Name   string `json:"Name"`
	Number string `json:"Number"`
	ID     string `json:"ID"`
}

type Users struct {
	UserIndex []User `json:"User"`
}

type Transaction struct {
	FromType string `json:"FromType"` //发送方角色	,institution:0,user:1
	FromID   string `json:"FromID"`   //发送方ID
	ToType   string `json:"ToType"`   //接收方角色	,institution:0,user:1
	ToID     string `json:"ToID"`     //接收方ID
	IDTime   string `json:"IDTime"`   //交易时间
	Number   string `json:"Number"`   //交易数额
	Rate     string `json:"Rate"`     //交易汇率
	ID       string `json:"ID"`       //交易ID
}

type Transactions struct {
	TransactionIndex []Transaction `json:"Transaction"`
}

var logger = logging.MustGetLogger("IntegrationManage")

type IntegrationChaincode struct {
}

func (this *IntegrationChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}

	// Create Institution table
	err := stub.CreateTable("Institution", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "ID", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "Name", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "TotalNumber", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "RestNumber", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed create Institution table")
	}

	// Create User table
	err = stub.CreateTable("User", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "ID", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "Name", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "Number", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed create User table")
	}

	// Create Transaction table
	err = stub.CreateTable("Transaction", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "ID", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "IDTime", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "FromType", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "FromID", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "ToType", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "ToID", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "Number", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "Rate", Type: shim.ColumnDefinition_STRING, Key: false},
	})

	if err != nil {
		return nil, errors.New("Failed create Transaction table")
	}

	return nil, nil
}

func (this *IntegrationChaincode) isCaller(stub shim.ChaincodeStubInterface, certificate []byte) (bool, error) {
	logger.Debug("Check caller...Begin!")

	// In order to enforce access control, we require that the
	// metadata contains the signature under the signing key corresponding
	// to the verification key inside certificate of
	// the payload of the transaction (namely, function name and args) and
	// the transaction binding (to avoid copying attacks)

	// Verify \sigma=Sign(certificate.sk, tx.Payload||tx.Binding) against certificate.vk
	// \sigma is in the metadata

	//	sigma, err := stub.GetCallerMetadata()
	//	if err != nil {
	//		return false, errors.New("Failed get metadata")
	//	}

	sigma, err := stub.GetCallerMetadata()
	if err != nil {
		logger.Error("Failed get metadata:", err)
		return false, errors.New("Failed get metadata")
	}
	payload, err := stub.GetPayload()
	if err != nil {
		logger.Error("Failed get payload:", err)
		return false, errors.New("Failed get payload")
	}
	binding, err := stub.GetBinding()
	if err != nil {
		logger.Error("Failed get binding:", err)
		return false, errors.New("Failed get binding")
	}

	ok, err := stub.VerifySignature(
		certificate,
		sigma,
		append(payload, binding...),
	)
	if err != nil {
		logger.Errorf("Failed checking signature [%s]", err)
		return ok, err
	}
	if !ok {
		logger.Error("Invalid signature")
	}

	logger.Debug("Check caller...Verified!")
	return ok, err
}

func (this *IntegrationChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	logger.Debugf("Invoke [%s]", function)
	//Verify the identity of the caller
	//	owner, err := base64.StdEncoding.DecodeString(args[0])
	//	if err != nil {
	//		logger.Error("Failed decode owner:", err)
	//		return nil, errors.New("Failed decode owner")
	//	}
	//	ok, err := this.isCaller(stub, owner)
	//	if err != nil {
	//		logger.Error("Failed check asset owner identity:", err)
	//		return nil, errors.New("Failed check asset owner identity")
	//	}
	//	if !ok {
	//		logger.Error("The caller is not the owner of the asset:", err)
	//		return nil, errors.New("The caller is not the owner of the asset")
	//	}
	if function == "createInstitution" {
		return this.createInstitution(stub, function, args)
	} else if function == "issueCoin" {
		return this.issueCoin(stub, function, args)
	} else if function == "issueCoinToUser" {
		return this.issueCoinToUser(stub, function, args)
	} else if function == "transfer" {
		// Transfer ownership
		return this.transfer(stub, function, args)
	}
	return nil, errors.New("Received unknown function invocation")
}

func (this *IntegrationChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	logger.Debugf("query [%s]", function)
	// Verify the identity of the caller
	//	owner, err := base64.StdEncoding.DecodeString(args[0])
	//	if err != nil {
	//		logger.Error("Failed decode owner:", err)
	//		return nil, errors.New("Failed decode owner")
	//	}
	//	ok, err := this.isCaller(stub, owner)
	//	if err != nil {
	//		logger.Error("Failed check asset owner identity:", err)
	//		return nil, errors.New("Failed check asset owner identity")
	//	}
	//	if !ok {
	//		logger.Error("The caller is not the owner of the asset:", err)
	//		return nil, errors.New("The caller is not the owner of the asset")
	//	}
	if function == "getInstitutionById" {
		return this.getInstitutionById(stub, args[1])
	} else if function == "getUserById" {
		return this.getUserById(stub, args[1])
	} else if function == "getTransactionById" {
		return this.getTransactionById(stub, args[1])
	} else if function == "getInstitution" {
		return this.getInstitution(stub)
	} else if function == "getUser" {
		return this.getUser(stub)
	} else if function == "getTransaction" {
		return this.getTransaction(stub)
	}

	return nil, errors.New("Received unknown function invocation")
}

func main() {
	primitives.SetSecurityLevel("SHA3", 256)
	err := shim.Start(new(IntegrationChaincode))
	if err != nil {
		fmt.Printf("Error starting IntegrationChaincode: %s", err)
	}
}
