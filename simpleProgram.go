package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type simpleProgram struct {
}

//Bank is a bank
type Bank struct {
	bankID     string `json:"ID"`
	bankName   string `json:"Name"`
	bankBranch string `json:"Branch"`
	bankCode   string `json:"Code"`

	//bankWalletID string `json:"walletID"`
	//bankAssetID  string `json:"assetID"`

	//bankChargesWallet   int `json:"chargesWallet"`
	//bankLiabilityWallet int `json:"liabilityWallet"`
	//TDSReceivableWallet int `json:"receivableWallet"`
}

//Business is a business
type Business struct {
	businessID   string `json:"ID"`
	businessName string `json:"Name"`
}

//Main
func main() {
	err := shim.Start(new(simpleProgram))
	if err != nil {
		fmt.Printf("error starting chaincode: %s", err)
	}
}

//Init initializes chaincode
func (t *simpleProgram) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

//Invoke invokes various chaincode functionalities
func (t *simpleProgram) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	//get the function and args
	function, args := stub.GetFunctionAndParameters()
	switch function {
	case "initBank":
		return t.initBank(stub, args)
	case "initBuissness":
		return t.initBuissness(stub, args)
	case "makePayment":
		return t.makePayment(stub, args)
	default:
		err := "invalid function\n"
		return shim.Error(err)
	}
}

func (t *simpleProgram) initBank(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	bankID := args[0]
	bankName := args[1]
	bankBranch := args[2]
	bankCode := args[3]

	bank := &Bank{bankID, bankName, bankBranch, bankCode}
	bankJSONasBytes, err := json.Marshal(bank)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(bankID, bankJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *simpleProgram) initBuissness(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	businessID := args[0]
	businessName := args[1]

	business := &Business{businessID, businessName}
	bankJSONasBytes, err := json.Marshal(business)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(businessID, bankJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)

}

func (t *simpleProgram) makePayment(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	return shim.Success(nil)

}

//create a bank with dummy variables
//initialize its wallet
//create a buissness with dummy variables
//initialize its wallet
//transfer money from buissness to bank and vice versa
