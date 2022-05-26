package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

type FoodContract struct{}

type food struct {
	OrderId                string `json:"orderId"`
	FoodId                 string `json:"foodId"`
	ConsumerId             string `json:"consumerId"`
	ManufactureId          string `json:"manufactureId"`
	WholesalerId           string `json:"wholesalerId"`
	RetailerId             string `json:"retailerId"`
	LogisticsId            string `json:"logisticsId"`
	Status                 string `json:"status"`
	RawFoodProcessDate     string `json:"rawProcessDate"`
	ManufactureProcessDate string `json:"manufactureProcessDate"`
	WholesaleProcessDate   string `json:"wholesaleProcessDate"`
	ShippingProcessDate    string `json:"shippingProcessDate"`
	RetailProcessDate      string `json:"retailProcessDate"`
	ProduceName            string `json:"name"`
	Grade                  string `json:"grade"`
	OrderPrice             int    `json:"orderPrice"`
	ShippingPrice          int    `json:"shippingPrice"`
	DeliveryDate           string `json:"deliveryDate"`
	Latitude 			   string `json:"latitude"`
	Longitude			   string `json:"longitude"`
}

func (t *FoodContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return setupFoodSupplyChainOrder(stub)
	// return shim.Success(nil)
}

func setupFoodSupplyChainOrder(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("_______________________________________________")
	fmt.Println("------ Chaincode Initiated Successfully -------")
	fmt.Println("_______________________________________________")

	_, args := stub.GetFunctionAndParameters()
	fmt.Println(args)
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
	if len(args) != 8 {
		fmt.Println(args)
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	fmt.Println(args)
	orderId := args[0] // THIS WILL BE THE SHA256 QR hash
	consumerId := args[1]
	orderPrice, _ := strconv.Atoi(args[2])
	shippingPrice, _ := strconv.Atoi(args[3])
	produceName := args[4]
	grade := args[5]
	lat:=args[6]
	long:=args[7]
	foodContract := food{
		OrderId: orderId, ConsumerId: consumerId, OrderPrice: orderPrice, ShippingPrice: shippingPrice, Status: "raw food created", ProduceName: produceName, Grade: grade,Latitude:lat,Longitude:long}
	foodBytes, _ := json.Marshal(foodContract)
	stub.PutState(foodContract.OrderId, foodBytes)
	
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
// func wrapWithStack() error {
// 	err := displayError()
// 	return errors.Wrap(err, "wrapping an application error with stack trace")
// }
// func displayError() error {
// 	return errors.New("example error message")
// }
func main() {
	// err := displayError()
	// fmt.Printf("print error without stack trace: %s\n\n", err)
	// fmt.Printf("print error with stack trace: %+v\n\n", err)
	// err = wrapWithStack()
	// fmt.Printf("%+v\n\n", err)
	err := shim.Start(new(FoodContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
