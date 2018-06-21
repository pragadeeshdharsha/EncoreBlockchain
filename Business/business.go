package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type chainCode struct {
}

type businessInfo struct {
	businessName         string
	businessAcNo         string
	businessLimit        string
	businessWalletID     string //Hash
	businessLoanWalletID string
	businessLiabilityID  string
	maxROI               float32
	minROI               float32
	numberOfPrograms     int
	businessExposure     string
}

func (c *chainCode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (c *chainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()

	if len(args) != 11 {
		return shim.Error("Invalid number of arguments")
	}

	if function == "putNewBusinessInfo" { //Inserting a New Business information
		return putNewBusinessInfo(stub, args)
	} else if function == "getBusinessInfo" { // To view a Business information
		return getBusinessInfo(stub, args)
	}
	return shim.Success(nil)
}

func putNewBusinessInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	// CONVERTING STRING INTO 64 BIT FLOAT CONVERTION
	maxROIconvertion, err := strconv.ParseFloat(args[7], 32)
	if err != nil {
		fmt.Printf("Invalid Maximum ROI: %s\n", args[7])
		return shim.Error(err.Error())
	}
	maxROIconvertion32 := float32(maxROIconvertion) //32 bit convertion, as ParseFloat returns 64 bit

	minROIconvertion, err := strconv.ParseFloat(args[8], 32)
	if err != nil {
		fmt.Printf("Invalid Minimum ROI: %s\n", args[8])
		return shim.Error(err.Error())
	}
	minROIconvertion32 := float32(minROIconvertion)

	numOfPrograms, err := strconv.Atoi(args[9])
	if err != nil {
		fmt.Printf("Number of programs should be integer: %s\n", args[9])
	}
	newInfo := businessInfo{args[1], args[2], args[3], args[4], args[5], args[6], maxROIconvertion32, minROIconvertion32, numOfPrograms, args[10]}
	newInfoBytes, _ := json.Marshal(newInfo)
	err = stub.PutState(args[0], newInfoBytes) // businessID = args[0]
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func getBusinessInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		//fmt.Println("Required only one argument")
		return shim.Error("Required only one argument")
	}

	parsedBusinessInfo := businessInfo{}
	businessIDvalue, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("Failed to get the business information: " + err.Error())
	} else if businessIDvalue == nil {
		return shim.Error("No information is avalilable on this businessID " + args[0])
	}

	err = json.Unmarshal(businessIDvalue, &parsedBusinessInfo)
	if err != nil {
		return shim.Error("Unable to parse into the structure " + err.Error())
	}
	return shim.Success(nil)
}

func main() {
	err := shim.Start(new(chainCode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}

}
