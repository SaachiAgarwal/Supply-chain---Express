/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Writing Your First Blockchain Application
 */

package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"

)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the car structure, with 4 properties.  Structure tags are used by encoding/json library
type Car struct {
	target interface{}
}

/*
 * The Init method is called when the Smart Contract "fabcar" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "fabcar"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "queryCar" {
		return s.queryCar(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "createCar" {
		return s.createCar(APIstub, args)
	} else if function == "queryAllCars" {
		return s.queryAllCars(APIstub)
	} else if function == "deleteCar" {
		return s.deleteCar(APIstub, args)
	}else if function == "changeCarOwner" {
		return s.changeCarOwner(APIstub, args)
	} else if function == "auditquery" {
		return s.auditquery(APIstub, args)
	}
	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) queryCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {


	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	carAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(carAsBytes)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	cars := []Car{
		
	}

	i := 0
	for i < len(cars) {
		fmt.Println("i is ", i)
		carAsBytes, _ := json.Marshal(cars[i])
		APIstub.PutState("CAR"+strconv.Itoa(i), carAsBytes)
		fmt.Println("Added", cars[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func dumpMap(space string, m map[string]interface{}) {
                for k, v := range m {
                                if mv, ok := v.(map[string]interface{}); ok {
                                                fmt.Printf("{ \"%v\": \n", k)
                                                dumpMap(space+"\t", mv)
                                                fmt.Printf("}\n")
 
                                } else {
                                                fmt.Printf("%v %v : %v\n", space, k, v)}
 
                }
 
}

func (s *SmartContract) createCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2{
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
 jsonMap := make(map[string]interface{})
 
                err := json.Unmarshal([]byte(args[1]), &jsonMap)
                if err != nil {
                                panic(err)
                }
                dumpMap("", jsonMap)
                // temp := jsonMap["Make"].([]interface{})
                // fmt.Println(temp[0] )
 
                /*var c = car{jsonMap}
                fmt.Println(c)*/
 
 
                jsonString, _ := json.Marshal(jsonMap)
                //var p = string(jsonString)
                //var object = eval(p);

	

        //var car = Car{object}      	
	//car.ObjectType = "Car"
        //car.idGenerated = args[0]
        

        
        //carAsBytes, _ := json.Marshal(car)

	
	//val,_:=json.Marshal(result)
  
	APIstub.PutState(args[0],jsonString)

	return shim.Success(nil)
}	

func (s *SmartContract) queryAllCars(APIstub shim.ChaincodeStubInterface) sc.Response {

queryString:= fmt.Sprintf("{\"selector\":{\"docType\":\"Car\"} }") 

	fmt.Println("queryString:= ", queryString)

	queryResults, err := getQueryResultForQueryString(APIstub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)

	}

func (s *SmartContract) changeCarOwner(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {


	if len(args) != 2 {
			return shim.Error("Incorrect number of arguments. Expecting 2")
		}

	carAsBytes, _ := APIstub.GetState(args[0])
	jsonMap1 := make(map[string]interface{})
        err1 := json.Unmarshal([]byte(carAsBytes), &jsonMap1)
                   if err1 != nil {
                                panic(err1)
                    }
        dumpMap("", jsonMap1)
        var newowner = args[1]

       jsonMap1["Owner"] = newowner
       jsonString1, _ := json.Marshal(jsonMap1)

        //car := Car{}

	//json.Unmarshal(carAsBytes, &car)
	//car.Owner = args[1]

	//carAsBytes, _ = json.Marshal(car)
	APIstub.PutState(args[0], jsonString1)

	return shim.Success(nil)
}

func (s *SmartContract) deleteCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
        

	carAsBytes, _ := APIstub.GetState(args[0])
         car := Car{}
	
	json.Unmarshal(carAsBytes, &car)


	carAsBytes, _ = json.Marshal(car)
	
	APIstub.DelState(args[0])                                                 //remove the key from chaincode state
	

	return shim.Success(nil)
}


func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

func (s *SmartContract) auditquery(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	type AuditHistory struct {
		TxId    string   `json:"txId"`
		Value   map[string]interface {}     `json:"value"`
	}
	var history []AuditHistory;
	//var car Car

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	carId := args[0]
	fmt.Printf("- start auditqueryForCar: %s\n", carId)

	// Get History
	resultsIterator, err := stub.GetHistoryForKey(carId)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		historyData, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

jsonMap2 := make(map[string]interface{})
 
                err2 := json.Unmarshal(historyData.Value, &jsonMap2)
                if err2 != nil {
                                panic(err2)
                }
                dumpMap("", jsonMap2)

		var tx AuditHistory
		tx.TxId = historyData.TxId                     //copy transaction id over
		//json.Unmarshal(historyData.Value, &car)     //un stringify it aka JSON.parse()
		if historyData.Value == nil {                  //car has been deleted
			json.Unmarshal(historyData.Value, &jsonMap2) //un stringify it aka JSON.parse()
//jsonString2, _ := json.Marshal(jsonMap2)			
tx.Value = jsonMap2                     //copy nil marble
		} else {
			json.Unmarshal(historyData.Value, &jsonMap2) //un stringify it aka JSON.parse()
//jsonString2, _ := json.Marshal(jsonMap2)			
tx.Value = jsonMap2                 //copy marble over
		}	
                    history = append(history, tx)              //add this tx to the list
	}
	fmt.Printf("- auditqueryForCar returning:\n%s", history)

	//change to array of bytes
	historyAsBytes, _ := json.Marshal(history)     //convert to array of bytes
	return shim.Success(historyAsBytes)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
