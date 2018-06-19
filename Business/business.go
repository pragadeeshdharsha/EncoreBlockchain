package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type program struct {
}

type businessInfo struct {
	repaymentAcNos []string `json:"Repayment A/c No"`
}

/*type repaymentAcNo struct {
	var
}*/

func (p *program) Init(stub shim.ChaincodeStubInterface) pb.Response {
	args := stub.GetStringArgs()
	fmt.Println("Converting into a list")
	repaymentAcNosList := businessInfo{repaymentAcNos: strings.Split(args[1], ",")}
	fmt.Printf("%+v", repaymentAcNosList)
	acntNosBytes, _ := json.Marshal(args[1])
	stub.PutState(args[0], acntNosBytes)
	fmt.Println("Successfully stored")
	return shim.Success(nil)
}

func (p *program) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()

	if function == "addRepaymentAcNos" {
		return p.addRepaymentAcNos(stub, args)
	}
	/*if function == "view" {
		return p.view(stub, args[0])
	}*/

	return shim.Success(nil)
}

func (p *program) addRepaymentAcNos(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	value, _ := stub.GetState(args[0])
	repaymentAcNosList := businessInfo{}
	json.Unmarshal(value, &repaymentAcNosList)
	list := strings.Split(args[1], ",")
	repaymentAcNosList.repaymentAcNos = append(repaymentAcNosList.repaymentAcNos, list...)
	a, _ := json.Marshal(repaymentAcNosList)
	stub.PutState(args[0], a)

	return shim.Success(nil)
}

/*func (p *program) view(stub shim.ChaincodeStubInterface, args string) pb.Response {

	value, _ := stub.GetState(args)

	accountList := businessInfo{}
	json.Unmarshal(value, &accountList)

	return shim.Success(value)
}*/

func main() {
	err := shim.Start(new(program))
	if err != nil {
		println("Unable to start the code")
	}
}
