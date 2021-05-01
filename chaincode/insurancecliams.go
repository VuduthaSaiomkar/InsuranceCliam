package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

// is smartcontract name
type newInsuranceClaim struct {
}

//IBClaim is structer to use track the claim process
type IBClaim struct {
	ClaimNumber  string `json:"claimnumber"`
	PolicyNumber string `json:"policynumber"`
	Description  string `json:"description"`
	PoliceStatus string `json:"policestatus"`
	ClaimStatus  string `json:"claimstatus"`
	ClaimAmount  string `json:"claimamount"`
}

// ===================================================================================
// Main
// ===================================================================================
func main() {
	err := shim.Start(new(newInsuranceClaim))
	if err != nil {
		fmt.Printf("Error starting newInsuranceClaim chaincode: %s", err)
	}
}

// Init initializes chaincode
// ===========================
func (t *newInsuranceClaim) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke - Our entry point for Invocations
// ========================================
func (t *newInsuranceClaim) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	mspid, err := cid.GetMSPID(stub)
	if err != nil {
		return shim.Error("something went wrong :" + err.Error())
	}
	if function == "fncreateclaim" && mspid == "workshopMSP" {
		return t.fncreateclaim(stub, args)
	} else if function == "fnupdatepolicestatus" && mspid == "policeMSP" {
		return t.fnupdatepolicestatus(stub, args)
	} else if function == "fnupdateinsurancestatus" && mspid == "insuranceMSP" {
		return t.fnupdateinsurancestatus(stub, args)
	} else if function == "fngetClaimbyID" {
		return t.fngetClaimbyID(stub, args)
	}
	return shim.Error("recevied unknown function name or user doesnot have access" + function)
}

//to add file into the network
func (t *newInsuranceClaim) fncreateclaim(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	var claimlog IBClaim
	err := json.Unmarshal([]byte(args[0]), &claimlog)
	claimlog.ClaimNumber = "Claim" + "" + fnGetTimsstampinSec(stub)
	claimlogBytes, _ := json.Marshal(claimlog)
	err = stub.PutState(claimlog.ClaimNumber, claimlogBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

/*
	getting timestamp for all transaction
*/

func fnGetTimsstampinSec(stub shim.ChaincodeStubInterface) string {

	sec, _ := stub.GetTxTimestamp()
	return string(sec.GetSeconds())
}

/*
	to get the file based on given hash
*/

func (t *newInsuranceClaim) fnupdatepolicestatus(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting filehash")
	}
	var claimlog IBClaim
	var d struct {
		ClaimNumber  string `json:"claimnumber"`
		PoliceStatus string `json:"policestatus"`
	}
	err := json.Unmarshal([]byte(args[0]), &d)
	assetAsBytes, err := stub.GetState(d.ClaimNumber)
	if err != nil {
		return shim.Error("Failed to get box:" + err.Error())
	} else if assetAsBytes == nil {
		return shim.Error("Claim number doesnot exists")
	}
	err = json.Unmarshal(assetAsBytes, &claimlog)
	claimlog.PoliceStatus = d.PoliceStatus
	claimlogBytes, _ := json.Marshal(claimlog)
	err = stub.PutState(claimlog.ClaimNumber, claimlogBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(claimlogBytes)
}

func (t *newInsuranceClaim) fnupdateinsurancestatus(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	var claimlog IBClaim
	var d struct {
		ClaimNumber     string `json:"claimnumber"`
		InsuranceStatus string `json:"InsuranceStatus"`
	}
	err := json.Unmarshal([]byte(args[0]), &d)
	assetAsBytes, err := stub.GetState(d.ClaimNumber)
	if err != nil {
		return shim.Error("Failed to get box:" + err.Error())
	} else if assetAsBytes == nil {
		return shim.Error("Claim number doesnot exists")
	}
	err = json.Unmarshal(assetAsBytes, &claimlog)
	claimlog.ClaimStatus = d.InsuranceStatus
	claimlogBytes, _ := json.Marshal(claimlog)
	err = stub.PutState(claimlog.ClaimNumber, claimlogBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(claimlogBytes)
}

func (t *newInsuranceClaim) fngetClaimbyID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	claimlogBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("Failed to get box:" + err.Error())
	} else if claimlogBytes == nil {
		return shim.Error("Claim number doesnot exists")
	}
	return shim.Success(claimlogBytes)
}
