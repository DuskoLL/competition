/*
 * 本合约是捐赠溯源合约教学案例
 */

 package main

 /* Imports 是用来引入
  *5个实体包：处理bytes("bytes")、读写JSON("encoding/json")、格式化代码("fmt")、字符串处理("strconv")、时间处理（"time"）
  *2个跟Hyperledger Fabric智能合约相关的包shim,peer
  */
 import (
	 "bytes"
	 "encoding/json"
	 "fmt"
	 "strconv"
	 "time"
 
	 "github.com/hyperledger/fabric/core/chaincode/shim"
	 sc "github.com/hyperledger/fabric/protos/peer"
 )
 
 // 定义捐赠溯源智能合约结构体
 type DonationTraceContract struct {
 }
 
 // 定义结构体DonationTrace， 包括2个属性：捐赠溯源内容、捐赠过程状态，结构标记用JSON表示
 type DonationTrace struct {
	 Content string `json:"content"`
	 State   string `json:"state"`
 }
 
 /*
  * Init方法是在智能合约初始化时由区块链网络调用的
  * 本合约无初始化数据
  */
 func (s *DonationTraceContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	 return shim.Success(nil)
 }
 
 /*
  * Invoke是具体运行应用的请求并返回结果的函数
  * 调用时要提供参数
  */
 func (s *DonationTraceContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
 
	 // 获得调用智能合约的函数名和参数
	 function, args := APIstub.GetFunctionAndParameters()
	 // 按函数名分别指向与其对应的处理函数，并关联处理相应的账本
	 if function == "createDonationTrace" {
		 return s.createDonationTrace(APIstub, args)
	 } else if function == "getHistoryForNumber" {
		 return s.getHistoryForNumber(APIstub, args)
	 }
 
	 return shim.Error("Invalid Smart Contract function name.")
 }
 
 /*
  * createDonationTrace是创建捐赠过程中的一条溯源数据，并写入账本，每条捐赠溯源数据包括3个数据内容：主键KEY、捐赠溯源内容、捐赠过程状态
  */
 func (s *DonationTraceContract) createDonationTrace(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 if len(args) != 3 {
		 return shim.Error("Incorrect number of arguments. Expecting 3")
	 }
 
	 var donationTrace = DonationTrace{Content: args[1], State: args[2]}
 
	 donationTraceAsBytes, _ := json.Marshal(donationTrace)
	 APIstub.PutState(args[0], donationTraceAsBytes)
 
	 return shim.Success(nil)
 }
 
 /*
  * getHistoryForNumber 是用主键KEY去查询捐赠过程中的所有历史过程内容信息
  */
 func (s *DonationTraceContract) getHistoryForNumber(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	
	 if len(args) < 1 {
		 return shim.Error("Incorrect number of arguments. Expecting 1")
	 }
 
	 donationNumber := args[0]
	 fmt.Printf("- start getHistoryForDonation: %s\n", donationNumber)
 
	 resultsIterator, err := APIstub.GetHistoryForKey(donationNumber)
	 if err != nil {
		 return shim.Error(err.Error())
	 }
	 defer resultsIterator.Close()
 
	 // buffer是一个JSON数组，用来存储DonationTrace的历史值
	 var buffer bytes.Buffer
	 buffer.WriteString("[")
 
	 bArrayMemberAlreadyWritten := false
	 for resultsIterator.HasNext() {
		 response, err := resultsIterator.Next()
		 if err != nil {
			 return shim.Error(err.Error())
		 }
		 // 在数组成员前增加一个逗号
		 if bArrayMemberAlreadyWritten == true {
			 buffer.WriteString(",")
		 }
 
		 buffer.WriteString("{\"TxId\":")
		 buffer.WriteString("\"")
		 buffer.WriteString(response.TxId)
		 buffer.WriteString("\"")
 
		 buffer.WriteString(", \"Value\":")
		 
		 if response.IsDelete {
			 buffer.WriteString("null")
		 } else {
			 buffer.WriteString(string(response.Value))
		 }
 
		 buffer.WriteString(", \"Timestamp\":")
		 buffer.WriteString("\"")
 
		 buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		 buffer.WriteString("\"")
 
		 buffer.WriteString(", \"IsDelete\":")
		 buffer.WriteString("\"")
		 buffer.WriteString(strconv.FormatBool(response.IsDelete))
		 buffer.WriteString("\"")
		 buffer.WriteString("}")
		 bArrayMemberAlreadyWritten = true
	 }
	 buffer.WriteString("]")
 
	 fmt.Printf("- getHistoryForDonation returning:\n%s\n", buffer.String())
 
	 return shim.Success(buffer.Bytes())
 
 }
 
 /*
  * 主函数，需要调用shim.Start()方法，启动链码必须通过调用shim包中的Start函数实现
  */
 func main() {
 
	 // 创建一个新的DonationTrace智能合约
	 err := shim.Start(new(DonationTraceContract))
	 if err != nil {
		 fmt.Printf("Error creating new Donation Contract Trace: %s", err)
	 }
 }
 