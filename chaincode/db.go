package main

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func (this *IntegrationChaincode) issueCoinToUser(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	/*
		args[1]     ID
		args[2]     Name
		args[3]     Number
		args[4]     UserID
		args[5]     UserName
		args[6]     TxID
	*/
	if len(args) != 7 {
		logger.Error("Incorrect number of arguments,Expecting 7")
		return nil, errors.New("Incorrect number of arguments,Expecting 7")
	}

	bIsFlag, err := this.queryInstitution(stub, args[1])
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	if !bIsFlag {
		logger.Error("Can't find Institution")
		return nil, errors.New("Can't find Institution")
	}

	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: args[1]}}
	col2 := shim.Column{Value: &shim.Column_String_{String_: args[2]}}
	columns = append(columns, col1)
	columns = append(columns, col2)

	var columnTemps []shim.Column
	colTemp1 := shim.Column{Value: &shim.Column_String_{String_: args[4]}}
	colTemp2 := shim.Column{Value: &shim.Column_String_{String_: args[5]}}
	columnTemps = append(columnTemps, colTemp1)
	columnTemps = append(columnTemps, colTemp2)

	row, err := stub.GetRow("Institution", columns)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	strTotalNumber := row.Columns[2].GetString_()
	strRestNumber := row.Columns[3].GetString_()
	RestNumber, err := strconv.ParseFloat(strRestNumber, 64)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	blance, err := strconv.ParseFloat(args[3], 64)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	if RestNumber < blance {
		logger.Error("not sufficient funds")
		return nil, errors.New("not sufficient funds")
	}

	RestNumber = RestNumber - blance

	ok, err := stub.ReplaceRow(
		"Institution",
		shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: args[1]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[2]}},
				&shim.Column{Value: &shim.Column_String_{String_: strTotalNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: strconv.FormatFloat(RestNumber, 'f', 2, 64)}}},
		})
	if !ok || err != nil {
		logger.Error("Failed replace:", err)
		return nil, err
	}

	bIsFlag, err = this.queryUser(stub, args[4])
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	if !bIsFlag {
		ok, err := stub.InsertRow(
			"User",
			shim.Row{
				Columns: []*shim.Column{
					&shim.Column{Value: &shim.Column_String_{String_: args[4]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[5]}},
					&shim.Column{Value: &shim.Column_String_{strconv.FormatFloat(blance, 'f', 2, 64)}}},
			})
		if !ok || err != nil {
			logger.Error("Failed insert UserRow:", err)
			return nil, errors.New("Failed insert UserRow")
		}
	} else {
		rowTemp, err := stub.GetRow("User", columnTemps)
		if err != nil {
			logger.Error(err)
			return nil, err
		}

		strBlance := rowTemp.Columns[2].GetString_()
		blanceTemp, err := strconv.ParseFloat(strBlance, 64)
		if err != nil {
			logger.Error(err)
			return nil, err
		}

		blance = blance + blanceTemp

		ok, err = stub.ReplaceRow(
			"User",
			shim.Row{
				Columns: []*shim.Column{
					&shim.Column{Value: &shim.Column_String_{String_: args[4]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[5]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[3]}}},
			})
		if !ok || err != nil {
			logger.Error("Failed replace:", err)
			return nil, err
		}
	}

	//获取时间戳
	timestamp := time.Now().Unix()

	err = this.writeTransaction(stub, args[6], strconv.FormatInt(timestamp, 10), args[1], args[4], "0", "1", args[3], "1.00")
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return nil, nil
}

func (this *IntegrationChaincode) transfer(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	/*
		args[1]     fromID
		args[2]     fromName
		args[3]     Number
		args[4]     toID
		args[5]     toName
		args[6]     rate
		args[7]     TxID
	*/
	if len(args) != 8 {
		logger.Error("Incorrect number of arguments,Expecting 8")
		return nil, errors.New("Incorrect number of arguments,Expecting 8")
	}

	bIsFlag, err := this.queryUser(stub, args[1])
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	if !bIsFlag {
		logger.Error("Can't find User")
		return nil, errors.New("Can't find User")
	}

	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: args[1]}}
	col2 := shim.Column{Value: &shim.Column_String_{String_: args[2]}}
	columns = append(columns, col1)
	columns = append(columns, col2)

	var columnTemps []shim.Column
	colTemp1 := shim.Column{Value: &shim.Column_String_{String_: args[4]}}
	colTemp2 := shim.Column{Value: &shim.Column_String_{String_: args[5]}}
	columnTemps = append(columnTemps, colTemp1)
	columnTemps = append(columnTemps, colTemp2)

	row, err := stub.GetRow("User", columns)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	strFromNumber := row.Columns[2].GetString_()
	fromNumber, err := strconv.ParseFloat(strFromNumber, 64)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	number, err := strconv.ParseFloat(args[3], 64)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	rate, err := strconv.ParseFloat(args[6], 64)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	fromNumberTemp := number * rate

	if fromNumber < fromNumberTemp {
		logger.Error("not sufficient funds")
		return nil, errors.New("not sufficient funds")
	}

	fromNumber = fromNumber - fromNumberTemp

	ok, err := stub.ReplaceRow(
		"User",
		shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: args[1]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[2]}},
				&shim.Column{Value: &shim.Column_String_{String_: strconv.FormatFloat(fromNumber, 'f', 2, 64)}}},
		})
	if !ok || err != nil {
		logger.Error("Failed replace:", err)
		return nil, err
	}

	bIsFlag, err = this.queryUser(stub, args[4])
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	if !bIsFlag {
		ok, err := stub.InsertRow(
			"User",
			shim.Row{
				Columns: []*shim.Column{
					&shim.Column{Value: &shim.Column_String_{String_: args[4]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[5]}},
					&shim.Column{Value: &shim.Column_String_{args[3]}}},
			})
		if !ok || err != nil {
			logger.Error("Failed insert UserRow:", err)
			return nil, errors.New("Failed insert UserRow")
		}
	} else {
		rowTemp, err := stub.GetRow("User", columnTemps)
		if err != nil {
			logger.Error(err)
			return nil, err
		}

		strBlance := rowTemp.Columns[2].GetString_()
		blanceTemp, err := strconv.ParseFloat(strBlance, 64)
		if err != nil {
			logger.Error(err)
			return nil, err
		}

		number = number + blanceTemp

		ok, err = stub.ReplaceRow(
			"User",
			shim.Row{
				Columns: []*shim.Column{
					&shim.Column{Value: &shim.Column_String_{String_: args[4]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[5]}},
					&shim.Column{Value: &shim.Column_String_{String_: strconv.FormatFloat(number, 'f', 2, 64)}}},
			})
		if !ok || err != nil {
			logger.Error("Failed replace:", err)
			return nil, err
		}
	}

	//获取时间戳
	timestamp := time.Now().Unix()

	err = this.writeTransaction(stub, args[7], strconv.FormatInt(timestamp, 10), args[1], args[4], "1", "1", args[3], args[6])
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return nil, nil
}

func (this *IntegrationChaincode) issueCoin(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	/*
		args[1]     ID
		args[2]     Name
		args[3]     Number
		args[4]     TxID
	*/
	if len(args) != 5 {
		logger.Error("Incorrect number of arguments,Expecting 5")
		return nil, errors.New("Incorrect number of arguments,Expecting 5")
	}
	logger.Debug("issueCoin args", args)
	bIsFlag, err := this.queryInstitution(stub, args[1])
	if err != nil {
		logger.Error(err)
	}
	if !bIsFlag {
		logger.Error("Can't find Institution")
		return nil, errors.New("Can't find Institution")
	}

	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: args[1]}}
	col2 := shim.Column{Value: &shim.Column_String_{String_: args[2]}}
	columns = append(columns, col1)
	columns = append(columns, col2)
	row, err := stub.GetRow("Institution", columns)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	//replace
	strTotalNumber := row.Columns[2].GetString_()
	strRestNumber := row.Columns[3].GetString_()
	TotalNumber, err := strconv.ParseFloat(strTotalNumber, 64)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	RestNumber, err := strconv.ParseFloat(strRestNumber, 64)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	NumberTemp, err := strconv.ParseFloat(args[3], 64)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	TotalNumber = TotalNumber + NumberTemp
	RestNumber = RestNumber + NumberTemp

	ok, err := stub.ReplaceRow(
		"Institution",
		shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: args[1]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[2]}},
				&shim.Column{Value: &shim.Column_String_{String_: strconv.FormatFloat(TotalNumber, 'f', 2, 64)}},
				&shim.Column{Value: &shim.Column_String_{String_: strconv.FormatFloat(RestNumber, 'f', 2, 64)}}},
		})
	if !ok || err != nil {
		logger.Error("Failed replace:", err)
		return nil, err
	}

	//获取时间戳
	timestamp := time.Now().Unix()

	err = this.writeTransaction(stub, args[4], strconv.FormatInt(timestamp, 10), args[1], args[1], "0", "0", args[3], "1")
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return nil, nil
}

func (this *IntegrationChaincode) writeTransaction(stub shim.ChaincodeStubInterface, ID string, IDTime string, FromID string, ToID string, FromType string, ToType string, Number string, Rate string) error {
	ok, err := stub.InsertRow(
		"Transaction",
		shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: ID}},
				&shim.Column{Value: &shim.Column_String_{String_: IDTime}},
				&shim.Column{Value: &shim.Column_String_{String_: FromType}},
				&shim.Column{Value: &shim.Column_String_{String_: FromID}},
				&shim.Column{Value: &shim.Column_String_{String_: ToType}},
				&shim.Column{Value: &shim.Column_String_{String_: ToID}},
				&shim.Column{Value: &shim.Column_String_{String_: Number}},
				&shim.Column{Value: &shim.Column_String_{String_: Rate}}},
		})
	if !ok || err != nil {
		logger.Error("Failed insert TransactionRow:", err)
		return errors.New("Failed insert TransactionRow")
	}
	logger.Debug("Insert TransactionRow done...")
	return nil
}

func (this *IntegrationChaincode) createInstitution(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	/*
		args[1]     ID
		args[2]     Name
	*/
	if len(args) != 3 {
		logger.Error("Incorrect number of arguments. Expecting 3")
		return nil, errors.New("Incorrect number of arguments. Expecting 3")
	}
	bIsFlag, err := this.queryInstitution(stub, args[1])
	if err != nil {
		logger.Error(err)
	}
	if bIsFlag {
		logger.Error("Institution exist")
		return nil, errors.New("Institution exist")
	}
	ok, err := stub.InsertRow(
		"Institution",
		shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: args[1]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[2]}},
				&shim.Column{Value: &shim.Column_String_{String_: "100000000"}},
				&shim.Column{Value: &shim.Column_String_{String_: "100000000"}}},
		})
	if !ok || err != nil {
		return nil, errors.New("Failed insert assetRow")
	}
	return nil, nil

}

func (this *IntegrationChaincode) queryInstitution(stub shim.ChaincodeStubInterface, id string) (bool, error) {
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: id}}
	columns = append(columns, col1)
	rowChannel, err := stub.GetRows("Institution", columns)
	bIsFlag := true
	if err != nil {
		bIsFlag = false
		logger.Error("Failed query Institution row")
		return bIsFlag, errors.New("Failed query Institution row")
	}

	var rows []shim.Row
	for {
		select {
		case row, ok := <-rowChannel:
			if !ok {
				rowChannel = nil
			} else {
				rows = append(rows, row)
			}
		}
		if rowChannel == nil {
			break
		}
	}

	if len(rows) == 0 {
		bIsFlag = false
		logger.Error("Can't find Institution")
		return bIsFlag, nil
	}

	return bIsFlag, nil
}

func (this *IntegrationChaincode) queryUser(stub shim.ChaincodeStubInterface, id string) (bool, error) {
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: id}}
	columns = append(columns, col1)
	rowChannel, err := stub.GetRows("User", columns)
	bIsFlag := true
	if err != nil {
		bIsFlag = false
		logger.Error("Failed query User row")
		return bIsFlag, errors.New("Failed query User row")
	}

	var rows []shim.Row
	for {
		select {
		case row, ok := <-rowChannel:
			if !ok {
				rowChannel = nil
			} else {
				rows = append(rows, row)
			}
		}
		if rowChannel == nil {
			break
		}
	}

	if len(rows) == 0 {
		bIsFlag = false
		logger.Error("Can't find User")
		return bIsFlag, nil
	}

	return bIsFlag, nil
}

func (this *IntegrationChaincode) getInstitutionById(stub shim.ChaincodeStubInterface, id string) ([]byte, error) {
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: id}}
	columns = append(columns, col1)
	rowChannel, err := stub.GetRows("Institution", columns)
	if err != nil {
		logger.Error("Failed query Institution row")
		return nil, errors.New("Failed query Institution row")
	}

	var rows []shim.Row
	for {
		select {
		case row, ok := <-rowChannel:
			if !ok {
				rowChannel = nil
			} else {
				rows = append(rows, row)
			}
		}
		if rowChannel == nil {
			break
		}
	}

	if len(rows) == 0 {
		logger.Error("Can't find Institution")
		return nil, errors.New("Can't find Institution")
	}

	jsonRe := Institutions{}
	var jsonRow Institution
	for _, row := range rows {
		jsonRow.ID = row.Columns[0].GetString_()
		jsonRow.Name = row.Columns[1].GetString_()
		jsonRow.TotalNumber = row.Columns[2].GetString_()
		jsonRow.RestNumber = row.Columns[3].GetString_()
		jsonRe.InstitutionIndex = append(jsonRe.InstitutionIndex, jsonRow)
	}
	jsonResp, err := json.Marshal(jsonRe)
	if err != nil {
		return nil, errors.New("Failed marshal Institutions")
	}

	return jsonResp, nil
}

func (this *IntegrationChaincode) getUserById(stub shim.ChaincodeStubInterface, id string) ([]byte, error) {
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: id}}
	columns = append(columns, col1)
	rowChannel, err := stub.GetRows("User", columns)
	if err != nil {
		logger.Error("Failed query User row")
		return nil, errors.New("Failed query User row")
	}

	var rows []shim.Row
	for {
		select {
		case row, ok := <-rowChannel:
			if !ok {
				rowChannel = nil
			} else {
				rows = append(rows, row)
			}
		}
		if rowChannel == nil {
			break
		}
	}

	if len(rows) == 0 {
		logger.Error("Can't find User")
		return nil, errors.New("Can't find User")
	}

	jsonRe := Users{}
	var jsonRow User
	for _, row := range rows {
		jsonRow.ID = row.Columns[0].GetString_()
		jsonRow.Name = row.Columns[1].GetString_()
		jsonRow.Number = row.Columns[2].GetString_()
		jsonRe.UserIndex = append(jsonRe.UserIndex, jsonRow)
	}
	jsonResp, err := json.Marshal(jsonRe)
	if err != nil {
		return nil, errors.New("Failed marshal User")
	}

	return jsonResp, nil
}

func (this *IntegrationChaincode) getTransactionById(stub shim.ChaincodeStubInterface, id string) ([]byte, error) {
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: id}}
	columns = append(columns, col1)
	rowChannel, err := stub.GetRows("Transaction", columns)
	if err != nil {
		logger.Error("Failed query Transaction row")
		return nil, errors.New("Failed query Transaction row")
	}

	var rows []shim.Row
	for {
		select {
		case row, ok := <-rowChannel:
			if !ok {
				rowChannel = nil
			} else {
				rows = append(rows, row)
			}
		}
		if rowChannel == nil {
			break
		}
	}

	if len(rows) == 0 {
		logger.Error("Can't find Transaction")
		return nil, errors.New("Can't find Transaction")
	}

	jsonRe := Transactions{}
	var jsonRow Transaction
	for _, row := range rows {
		jsonRow.ID = row.Columns[0].GetString_()
		jsonRow.IDTime = row.Columns[1].GetString_()
		jsonRow.FromType = row.Columns[2].GetString_()
		jsonRow.FromID = row.Columns[3].GetString_()
		jsonRow.ToType = row.Columns[4].GetString_()
		jsonRow.ToID = row.Columns[5].GetString_()
		jsonRow.Number = row.Columns[6].GetString_()
		jsonRow.Rate = row.Columns[7].GetString_()
		jsonRe.TransactionIndex = append(jsonRe.TransactionIndex, jsonRow)
	}
	jsonResp, err := json.Marshal(jsonRe)
	if err != nil {
		return nil, errors.New("Failed marshal Transaction")
	}

	return jsonResp, nil
}

func (this *IntegrationChaincode) getInstitution(stub shim.ChaincodeStubInterface) ([]byte, error) {
	var columns []shim.Column
	rowChannel, err := stub.GetRows("Institution", columns)
	if err != nil {
		logger.Error("Failed query Institution row")
		return nil, errors.New("Failed query Institution row")
	}

	var rows []shim.Row
	for {
		select {
		case row, ok := <-rowChannel:
			if !ok {
				rowChannel = nil
			} else {
				rows = append(rows, row)
			}
		}
		if rowChannel == nil {
			break
		}
	}

	if len(rows) == 0 {
		logger.Error("Can't find Institution")
		return nil, errors.New("Can't find Institution")
	}

	jsonRe := Institutions{}
	var jsonRow Institution
	for _, row := range rows {
		jsonRow.ID = row.Columns[0].GetString_()
		jsonRow.Name = row.Columns[1].GetString_()
		jsonRow.TotalNumber = row.Columns[2].GetString_()
		jsonRow.RestNumber = row.Columns[3].GetString_()
		jsonRe.InstitutionIndex = append(jsonRe.InstitutionIndex, jsonRow)
	}
	jsonResp, err := json.Marshal(jsonRe)
	if err != nil {
		return nil, errors.New("Failed marshal Institutions")
	}

	return jsonResp, nil
}

func (this *IntegrationChaincode) getUser(stub shim.ChaincodeStubInterface) ([]byte, error) {
	var columns []shim.Column
	rowChannel, err := stub.GetRows("User", columns)
	if err != nil {
		logger.Error("Failed query User row")
		return nil, errors.New("Failed query User row")
	}

	var rows []shim.Row
	for {
		select {
		case row, ok := <-rowChannel:
			if !ok {
				rowChannel = nil
			} else {
				rows = append(rows, row)
			}
		}
		if rowChannel == nil {
			break
		}
	}

	if len(rows) == 0 {
		logger.Error("Can't find User")
		return nil, errors.New("Can't find User")
	}

	jsonRe := Users{}
	var jsonRow User
	for _, row := range rows {
		jsonRow.ID = row.Columns[0].GetString_()
		jsonRow.Name = row.Columns[1].GetString_()
		jsonRow.Number = row.Columns[2].GetString_()
		jsonRe.UserIndex = append(jsonRe.UserIndex, jsonRow)
	}
	jsonResp, err := json.Marshal(jsonRe)
	if err != nil {
		return nil, errors.New("Failed marshal User")
	}

	return jsonResp, nil
}

func (this *IntegrationChaincode) getTransaction(stub shim.ChaincodeStubInterface) ([]byte, error) {
	var columns []shim.Column
	rowChannel, err := stub.GetRows("Transaction", columns)
	if err != nil {
		logger.Error("Failed query Transaction row")
		return nil, errors.New("Failed query Transaction row")
	}

	var rows []shim.Row
	for {
		select {
		case row, ok := <-rowChannel:
			if !ok {
				rowChannel = nil
			} else {
				rows = append(rows, row)
			}
		}
		if rowChannel == nil {
			break
		}
	}

	if len(rows) == 0 {
		logger.Error("Can't find Transaction")
		return nil, errors.New("Can't find Transaction")
	}

	jsonRe := Transactions{}
	var jsonRow Transaction
	for _, row := range rows {
		jsonRow.ID = row.Columns[0].GetString_()
		jsonRow.IDTime = row.Columns[1].GetString_()
		jsonRow.FromType = row.Columns[2].GetString_()
		jsonRow.FromID = row.Columns[3].GetString_()
		jsonRow.ToType = row.Columns[4].GetString_()
		jsonRow.ToID = row.Columns[5].GetString_()
		jsonRow.Number = row.Columns[6].GetString_()
		jsonRow.Rate = row.Columns[7].GetString_()
		jsonRe.TransactionIndex = append(jsonRe.TransactionIndex, jsonRow)
	}
	jsonResp, err := json.Marshal(jsonRe)
	if err != nil {
		return nil, errors.New("Failed marshal Transaction")
	}

	return jsonResp, nil
}
