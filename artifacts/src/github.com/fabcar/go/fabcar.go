package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/pkg/errors"
)

type FoodContract struct{}

type food struct {
	OrderId                string
	FoodId                 string
	ConsumerId             string
	ManufactureId          string
	WholesalerId           string
	RetailerId             string
	LogisticsId            string
	Status                 string
	RawFoodProcessDate     string
	ManufactureProcessDate string
	WholesaleProcessDate   string
	ShippingProcessDate    string
	RetailProcessDate      string
	ProduceName            string
	Grade                  string
	OrderPrice             int
	ShippingPrice          int
	DeliveryDate           string
}

func (t *FoodContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return setupFoodSupplyChainOrder(stub)
}

func setupFoodSupplyChainOrder(stub shim.ChaincodeStubInterface) pb.Response {
	_, args := stub.GetFunctionAndParameters()
	orderId := args[0]
	consumerId := args[1]
	orderPrice, _ := strconv.Atoi(args[2])
	shippingPrice, _ := strconv.Atoi(args[3])
	produceName := args[4]
	grade := args[5]
	foodContract := food{
		OrderId: orderId, ConsumerId: consumerId, OrderPrice: orderPrice, ShippingPrice: shippingPrice, Status: "order initiated", ProduceName: produceName, Grade: grade}
	foodBytes, _ := json.Marshal(foodContract)
	stub.PutState(foodContract.OrderId, foodBytes)
	return shim.Success(nil)
}

func (t *FoodContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "createRawFood" {
		return t.createRawFood(stub, args)
	} else if function == "manufactureProcessing" {
		return t.manufactureProcessing(stub, args)
	} else if function == "wholesalerDistribute" {
		return t.wholesalerDistribute(stub, args)
	} else if function == "initiateShipment" {
		return t.initiateShipment(stub, args)
	} else if function == "deliverToRetail" {
		return t.deliverToRetail(stub, args)
	} else if function == "completeOrder" {
		return t.completeOrder(stub, args)
	} else if function == "query" {
		return t.query(stub, args)
	}
	return shim.Error("Invalid function name")
}

func (f *FoodContract) createRawFood(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	orderId := args[0] // THIS WILL BE THE SHA256 QR hash
	foodBytes, _ := stub.GetState(orderId)
	fd := food{}
	json.Unmarshal(foodBytes, &fd)
	if fd.Status == "order initiated" {
		fd.FoodId = args[1] // Take input of farmer ID
		currentts := time.Now()
		fd.RawFoodProcessDate = currentts.Format("2006-01-02 15:04:05")
		fd.Status = "raw food created"
	} else {
		fmt.Printf("Order not initiated yet")
	}
	foodBytes, _ = json.Marshal(fd)
	stub.PutState(orderId, foodBytes)
	return shim.Success(nil)
}

func (f *FoodContract) manufactureProcessing(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	orderId := args[0]
	foodBytes, err := stub.GetState(orderId)
	fd := food{}
	err = json.Unmarshal(foodBytes, &fd)
	if err != nil {
		return shim.Error(err.Error())
	}
	if fd.Status == "raw food created" {
		fd.ManufactureId = args[1] // Manufacturer ID
		currentts := time.Now()
		fd.ManufactureProcessDate = currentts.Format("2006-01-02 15:04:05")
		fd.Status = "manufactured food"
	} else {
		fmt.Printf("Raw food not generated yet")
	}
	foodBytes0, _ := json.Marshal(fd)
	err = stub.PutState(orderId, foodBytes0)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (f *FoodContract) wholesalerDistribute(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	orderId := args[0]
	foodBytes, err := stub.GetState(orderId)
	fd := food{}
	err = json.Unmarshal(foodBytes, &fd)
	if err != nil {
		return shim.Error(err.Error())
	}
	if fd.Status == "manufactured food" {
		fd.WholesalerId = args[1] // Wholsaler ID
		currentts := time.Now()
		fd.WholesaleProcessDate = currentts.Format("2006-01-02 15:04:05")
		fd.Status = "wholesaler distribute"
	} else {
		fmt.Printf("Food not yet manufactured")
	}
	foodBytes0, _ := json.Marshal(fd)
	err = stub.PutState(orderId, foodBytes0)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (f *FoodContract) initiateShipment(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	orderId := args[0]
	foodBytes, err := stub.GetState(orderId)
	fd := food{}
	err = json.Unmarshal(foodBytes, &fd)
	if err != nil {
		return shim.Error(err.Error())
	}
	if fd.Status == "wholesaler distribute" {
		fd.LogisticsId = args[1] // Logistic ID
		currentts := time.Now()
		fd.ShippingProcessDate = currentts.Format("2006-01-02 15:04:05")
		fd.Status = "initiated shipment"
	} else {
		fmt.Printf("Wholesaler not initiated yet")
	}
	foodBytes0, _ := json.Marshal(fd)
	err = stub.PutState(orderId, foodBytes0)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (f *FoodContract) deliverToRetail(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	orderId := args[0]
	foodBytes, err := stub.GetState(orderId)
	fd := food{}
	err = json.Unmarshal(foodBytes, &fd)
	if err != nil {
		return shim.Error(err.Error())
	}
	if fd.Status == "initiated shipment" {
		fd.RetailerId = args[1] // Retailer ID
		currentts := time.Now()
		fd.RetailProcessDate = currentts.Format("2006-01-02 15:04:05")
		fd.Status = "Retailer started"
	} else {
		fmt.Printf("Shipment not initiated yet")
	}
	foodBytes0, _ := json.Marshal(fd)
	err = stub.PutState(orderId, foodBytes0)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (f *FoodContract) completeOrder(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	orderId := args[0]

	foodBytes, err := stub.GetState(orderId)
	fd := food{}
	err = json.Unmarshal(foodBytes, &fd)
	if err != nil {
		return shim.Error(err.Error())
	}
	if fd.Status == "Retailer started" {
		currentts := time.Now()
		fd.DeliveryDate = currentts.Format("2006-01-02 15:04:05")
		fd.Status = "Consumer received order"
	} else {
		fmt.Printf("Retailer not initiated yet")
	}
	foodBytes0, _ := json.Marshal(fd)
	err = stub.PutState(orderId, foodBytes0)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

// QUERY //

func (f *FoodContract) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var ENIITY string
	var err error
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expected ENIITY Name")
	}
	ENIITY = args[0]
	Avalbytes, err := stub.GetState(ENIITY)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + ENIITY + "\"}"
		return shim.Error(jsonResp)
	}
	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil order for " + ENIITY + "\"}"

		return shim.Error(jsonResp)
	}
	return shim.Success(Avalbytes)
}

//wrapping error with stack
func wrapWithStack() error {
	err := displayError()
	return errors.Wrap(err, "wrapping an application error with stack trace")
}
func displayError() error {
	return errors.New("example error message")
}
func main() {
	err := displayError()
	fmt.Printf("print error without stack trace: %s\n\n", err)
	fmt.Printf("print error with stack trace: %+v\n\n", err)
	err = wrapWithStack()
	fmt.Printf("%+v\n\n", err)
}
