/*
Copyright IBM Corp 2016 All Rights Reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
		 http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	//"strconv"
	//s "strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var i int
//**************************************************************************
//Entity - Structure for an entity like ClaimingInsurer, AtFaultInsurer, oem, customer
//****************************************************************************

type Entity struct {
	Type    string  `json:"type"`
	Name    string  `json:"name"`
	Qty     string     `json:"qty"`
}


//Product - Structure for products used in create and transfer
type TxnApprove struct {

	toEntity			string `json:"toEntity"`
	fromEntity			string `json:"fromEntity"`
    claimNo          	string `json:"claimNo"`
    policyNo           	string `json:"policyNo"`
    insured            	string `json:"insured"`
    lossDate           	string `json:"lossDate"`
	lossType      		string `json:"lossType"`
	lossDesp     		string `json:"lossDesp"`
	amountPaid   		string `json:"amountPaid"`
	faultParty      	string `json:"faultParty"`
	faultInsurer    	string `json:"faultInsurer"`
	percentage      	string `json:"percentage"`
	doc      			string `json:"doc"`
	docType     		string `json:"docType"`
	subrogationAmount   string `json:"subrogationAmount"`
	remarks        	 	string `json:"remarks"`
	subrogationDate     string `json:"subrogationDate"`
	status      		string `json:"status"`
	documentHash		string `json:"documentHash"`
	fileName			string `json:"fileName"`
	uploadId			string `json:"uploadId"`
	settlementAmount	string `json:"settlementAmount"`
	settlementDate		string `json:"settlementDate"`
	settlementRemarks	string `json:"settlementRemarks"`
	ID       			string `json:"ID"`
  Time           string `json:"Time"`

}

type TxnClaim struct {

	toEntity			string `json:"toEntity"`
	fromEntity			string `json:"fromEntity"`
    claimNo          	string `json:"claimNo"`
    policyNo           	string `json:"policyNo"`
    insured            	string `json:"insured"`
    lossDate           	string `json:"lossDate"`
	lossType      		string `json:"lossType"`
	lossDesp     		string `json:"lossDesp"`
	amountPaid   		string `json:"amountPaid"`
	faultParty      	string `json:"faultParty"`
	faultInsurer    	string `json:"faultInsurer"`
	percentage      	string `json:"percentage"`
	doc      			string `json:"doc"`
	docType     		string `json:"docType"`
	subrogationAmount   string `json:"subrogationAmount"`
	remarks        	 	string `json:"remarks"`
	subrogationDate     string `json:"subrogationDate"`
	status      		string `json:"status"`
	documentHash		string `json:"documentHash"`
	fileName			string `json:"fileName"`
	uploadId			string `json:"uploadId"`
	ID       			string `json:"ID"`
  Time           string `json:"Time"`


}

//******************************************************************************
// sabrogationChaincode example simple Chaincode implementation
//*******************************************************************************

type sabrogationChaincode struct {
}

func main() {
	err := shim.Start(new(sabrogationChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}


//*******************************************************************************
// Init resets all the things
//*********************************************************************************


func (t *sabrogationChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4")
	}

	key1 := args[0] // ClaimingInsurer
	key2 := args[1] // AtFaultInsurer


	ClaimingInsurer := Entity{
		Type:    "ClaimingInsurer",
		Name:    key1,
		Qty: 	 "1000",
	}
	fmt.Println(ClaimingInsurer)
	bytes, err := json.Marshal(ClaimingInsurer)
	if err != nil {
		fmt.Println("Error marsalling")
		return nil, errors.New("Error marshalling")
	}
	fmt.Println(bytes)
	err = stub.PutState(key1, bytes)
	if err != nil {
		fmt.Println("Error writing state")
		return nil, err
	}

	AtFaultInsurer := Entity{
		Type:    "AtFaultInsurer",
    Name:    key2,
		Qty: 	 "1000",
	}

	fmt.Println(AtFaultInsurer)
	bytes, err = json.Marshal(AtFaultInsurer)
	if err != nil {
		fmt.Println("Error marsalling")
		return nil, errors.New("Error marshalling")
	}
	fmt.Println(bytes)
	err = stub.PutState(key2, bytes)
	if err != nil {
		fmt.Println("Error writing state")
		return nil, err
	}




    //****************************************************************************
	// Initialize the collection of  keys for products and various transactions
	//*****************************************************************************

	fmt.Println("Initializing keys collection")
	var blank []string
	blankBytes, _ := json.Marshal(&blank)
	err = stub.PutState("TxnClaim", blankBytes)
	if err != nil {
		fmt.Println("Failed to initialize Products key collection")
	}

	blankBytes, _ = json.Marshal(&blank)
	err = stub.PutState("TxnApprove", blankBytes)
	if err != nil {
		fmt.Println("Failed to initialize Products key collection")
	}

	return nil, nil
}

// *************************************************************************************************************
// Invoke isur entry point to invoke a chaincode function
//******************************************************************************************************************

func (t *sabrogationChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions/transactions
	if function == "init" {
		return t.Init(stub, "init", args)
	}  else if function == "createClaim" {
		return t.createClaim(stub, args)
	} else if function == "approveClaim" {
		return t.approveClaim(stub, args)
	}
	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *sabrogationChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)


	// Handle different functions
	if function == "read" {
		return t.read(stub, args)
	} else if function == "getAllTxnClaim" {
		return t.getAllTxnClaim(stub)
	}  else if function == "getAllTxnApprove" {
		return t.getAllTxnApprove(stub)
	}
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}

// write - invoke function to write key/value pair
// func (t *sabrogationChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

// 	fmt.Println("running write()")

// 	if len(args) != 4 {
// 		return nil, errors.New("Incorrect number of arguments. expecting 3")
// 	}

// 	//writing a new customer to blockchain
// 	typeOf := args[0]
// 	name := args[1]
// 	qty  := args[2]


// 	entity := Entity{
// 		Type:    typeOf,
// 		Name:    name,
// 		Qty:     qty,
// 	}
// 	fmt.Println(entity)
// 	bytes, err := json.Marshal(entity)
// 	if err != nil {
// 		fmt.Println("Error marsalling")
// 		return nil, errors.New("Error marshalling")
// 	}
// 	fmt.Println(bytes)
// 	err = stub.PutState(name, bytes)
// 	if err != nil {
// 		fmt.Println("Error writing state")
// 		return nil, err
// 	}

// 	return nil, nil
// }

// read - query function to read key/value pair
func (t *sabrogationChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("read() is running")

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. expecting 1")
	}

	key := args[0] // name of Entity

	bytes, err := stub.GetState(key)
	if err != nil {
		fmt.Println("Error retrieving " + key)
		return nil, errors.New("Error retrieving " + key)
	}
	customer := Entity{}
	err = json.Unmarshal(bytes, &customer)
	if err != nil {
		fmt.Println("Error Unmarshaling customerBytes")
		return nil, errors.New("Error Unmarshaling customerBytes")
	}
	bytes, err = json.Marshal(customer)
	if err != nil {
		fmt.Println("Error marshaling customer")
		return nil, errors.New("Error marshaling customer")
	}

	fmt.Println(bytes)
	return bytes, nil
}



// createClaim

func (t *sabrogationChaincode) createClaim(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	fmt.Println("add is running ")

	if len(args) != 21 {
		return nil, errors.New("Incorrect Number of arguments.Expecting 21 for add")
	}


	ID := stub.GetTxID()
blockTime, err := stub.GetTxTimestamp()
if err != nil {
	
	return nil, err
}
args = append(args, ID)
args = append(args, blockTime.String())
		t.putTxnClaim(stub, args)

fmt.Println("Error Unmarshaling entity Bytes")

	return nil, nil
}


func (t *sabrogationChaincode) putTxnClaim(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("putTxnClaim is running ")

fmt.Println("Error Unmarshaling entity Bytes")
	if len(args) != 23 {
		return nil, errors.New("Incorrect Number of arguments.Expecting 23 for putTxnClaim")
	}
	  txn := TxnClaim{
		  		toEntity:args[0],
				fromEntity:args[1],
				claimNo: args[2],
                policyNo: args[3],
                insured: args[4],
                lossDate: args[5],
                lossType: args[6],
                lossDesp: args[7],
                amountPaid: args[8],
                faultParty: args[9],
                faultInsurer: args[10],
                percentage: args[11],
                doc: args[12],
                docType: args[13],
                subrogationAmount: args[14],
                remarks: args[15],
                subrogationDate: args[16],
                status: args[17],
				documentHash:args[18],
				fileName:args[19],
				uploadId:args[20],
				ID:args[21],
                Time:args[22],

      }

	bytes, err := json.Marshal(txn)
	if err != nil {
		fmt.Println("Error marshaling putTxnClaim")
		return nil, errors.New("Error marshaling putTxnClaim")
	}

	err = stub.PutState(txn.ID, bytes)
	if err != nil {
		fmt.Println("$$$#############################$$$")
fmt.Println(err)
fmt.Println("$$@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@$")
		return nil, err
	}

fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
fmt.Println(bytes)
fmt.Println("$$$$$$$$$$$$$$$$$*****************$")
	return t.appendKey(stub, "TxnClaim", txn.ID)
}

// approve

func (t *sabrogationChaincode) approveClaim(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	fmt.Println("add is running ")

	if len(args) != 24 {
		return nil, errors.New("Incorrect Number of arguments.Expecting 24 for add")
	}

	key := args[0]   //Entity ex: customer

	// GET the state of entity from the ledger
	bytes, err := stub.GetState(key)
	if err != nil {
		return nil, errors.New("Failed to get state of " + key)
	}

	entity := Entity{}
	err = json.Unmarshal(bytes, &entity)
	if err != nil {
		fmt.Println("Error Unmarshaling entity Bytes")
		return nil, errors.New("Error Unmarshaling entity Bytes")
	}

             //type1 := entity.Type




	// Write the state back to the ledger
	bytes, err = json.Marshal(entity)
	if err != nil {
		fmt.Println("Error marshaling entity")
		return nil, errors.New("Error marshaling entity")
	}
	err = stub.PutState(key, bytes)
	if err != nil {
		return nil, err
	}

        args = append(args, stub.GetTxID())
		blockTime, err := stub.GetTxTimestamp()
		if err != nil {
			return nil, err
		}
		args = append(args, blockTime.String())
		t.putTxnApprove(stub, args)


	return nil, nil
}
func (t *sabrogationChaincode) putTxnApprove(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("putTxnApprove is running ")

	if len(args) != 26 {
		return nil, errors.New("Incorrect Number of arguments.Expecting 26 for putTxnApprove")
	}
	  txn := TxnApprove{
		  		toEntity:args[0],
				fromEntity:args[1],
				claimNo: args[2],
                policyNo: args[3],
                insured: args[4],
                lossDate: args[5],
                lossType: args[6],
                lossDesp: args[7],
                amountPaid: args[8],
                faultParty: args[9],
                faultInsurer: args[10],
                percentage: args[11],
                doc: args[12],
                docType: args[13],
                subrogationAmount: args[14],
                remarks: args[15],
                subrogationDate: args[16],
                status: args[17],
				documentHash:args[18],
				fileName:args[19],
				uploadId:args[20],
				settlementAmount:args[21],
				settlementDate:	args[22],
				settlementRemarks:args[23],
				ID:args[24],
                Time:args[25],

      }

	bytes, err := json.Marshal(txn)
	if err != nil {
		fmt.Println("Error marshaling putTxnApprove")
		return nil, errors.New("Error marshaling putTxnApprove")
	}

	err = stub.PutState(txn.ID, bytes)
	if err != nil {
		return nil, err
	}

	return t.appendKey(stub, "putTxnApprove", txn.ID)
}



func (t *sabrogationChaincode) getAllTxnClaim(stub shim.ChaincodeStubInterface) ([]byte, error) {
	fmt.Println("getAllTxnTopup is running ")

	var txns []TxnClaim
	fmt.Println("get%%%%%%%%%%%%%% ")

	// Get list of all the keys - TxnClaim
	keysBytes, err := stub.GetState("TxnClaim")
	if err != nil {
		fmt.Println("Error retrieving TxnClaim keys")
		return nil, errors.New("Error retrieving TxnClaim keys")
	}
	var keys []string
	err = json.Unmarshal(keysBytes, &keys)
	if err != nil {
		fmt.Println("Error unmarshalling TxnClaim key")
		return nil, errors.New("Error unmarshalling TxnClaim keys")
	}

	// Get each product txn "TxnClaim" keys
	for _, value := range keys {
		bytes, err := stub.GetState(value)

		var txn TxnClaim
		err = json.Unmarshal(bytes, &txn)
		if err != nil {
			fmt.Println("Error retrieving txn " + value)
			return nil, errors.New("Error retrieving txn " + value)
		}

		fmt.Println("Appending txn" + value)
		txns = append(txns, txn)

		fmt.Println("txn is...................")
		fmt.Println(txn)
	}
	fmt.Println("txnssssssss are.......................")
		fmt.Println(txns)
	fmt.Println("txnended............")

	bytes, err := json.Marshal(txns)
	fmt.Println("txnbytesssssssssssssss are.......................")
		fmt.Println(bytes)
	fmt.Println("tbytes ended..................................")
	if err != nil {
		fmt.Println("Error marshaling txns topup")
		return nil, errors.New("Error marshaling txns topup")
	}

	return bytes, nil
}
func (t *sabrogationChaincode) getAllTxnApprove(stub shim.ChaincodeStubInterface) ([]byte, error) {
	fmt.Println("getAllTxnTopup is running ")

	var txns []TxnApprove

	// Get list of all the keys - TxnApprove
	keysBytes, err := stub.GetState("TxnApprove")
	if err != nil {
		fmt.Println("Error retrieving TxnApprove keys")
		return nil, errors.New("Error retrieving TxnApprove keys")
	}
	var keys []string
	err = json.Unmarshal(keysBytes, &keys)
	if err != nil {
		fmt.Println("Error unmarshalling TxnApprove key")
		return nil, errors.New("Error unmarshalling TxnApprove keys")
	}

	// Get each product txn "TxnApprove" keys
	for _, value := range keys {
		bytes, err := stub.GetState(value)

		var txn TxnApprove
		err = json.Unmarshal(bytes, &txn)
		if err != nil {
			fmt.Println("Error retrieving txn " + value)
			return nil, errors.New("Error retrieving txn " + value)
		}

		fmt.Println("Appending txn" + value)
		txns = append(txns, txn)
	}

	bytes, err := json.Marshal(txns)
	if err != nil {
		fmt.Println("Error marshaling txns topup")
		return nil, errors.New("Error marshaling txns topup")
	}
	return bytes, nil
}

func (t *sabrogationChaincode) appendKey(stub shim.ChaincodeStubInterface, primeKey string, key string) ([]byte, error) {
	fmt.Println("appendKey is running " + primeKey + " " + key)

	bytes, err := stub.GetState(primeKey)
	if err != nil {
		return nil, err
	}
	var keys []string
	err = json.Unmarshal(bytes, &keys)
	if err != nil {
		return nil, err
	}
	keys = append(keys, key)
	bytes, err = json.Marshal(keys)
	if err != nil {
		fmt.Println("Error marshaling " + primeKey)
		return nil, errors.New("Error marshaling keys" + primeKey)
	}
	err = stub.PutState(primeKey, bytes)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
