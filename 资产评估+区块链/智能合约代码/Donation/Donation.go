/*
 * 本合约是捐赠存证合约教学案例
 */

 package main

 /* Imports 是用来引入
  *2个实体包：格式化代码("fmt")、读写JSON("encoding/json")
  *2个跟Hyperledger Fabric智能合约相关的包shim,peer
  */
 import (
	 "encoding/json"
	 "fmt"
 
	 "github.com/hyperledger/fabric/core/chaincode/shim"
	 sc "github.com/hyperledger/fabric/protos/peer"
 )
 
 // 定义捐赠存证智能合约结构体
 type DonationContract struct {
 }
 
 // 定义结构体Donation， 包括4个属性：捐赠项目、捐赠人、捐赠内容、接收机构，结构标记用JSON表示
 type Donation struct {
	 Project  string `json:"project"`
	 Donator  string `json:"donator"`
	 Content  string `json:"content"`
	 Receiver string `json:"receiver"`
 }
 
 /*
  * Init方法是在智能合约 "Donation"初始化时由区块链网络调用的
  * 本合约无初始化数据
  */
 func (s *DonationContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	 return shim.Success(nil)
 }
 
 /*
  * Invoke是具体运行应用的请求并返回结果的函数
  * 调用时要提供参数
  */
 func (s *DonationContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
 
	 // 获得调用智能合约的函数名和参数
	 function, args := APIstub.GetFunctionAndParameters()
	 // 按函数名分别指向与其对应的处理函数，并关联处理相应的账本
	 if function == "queryDonation" {
		 return s.queryDonation(APIstub, args)
	 } else if function == "createDonation" {
		 return s.createDonation(APIstub, args)
	 }
 
	 return shim.Error("Invalid Smart Contract function name.")
 }
 
 /*
  * queryDonation 是用主键KEY去查询对应的捐赠信息
  */
 func (s *DonationContract) queryDonation(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 if len(args) != 1 {
		 return shim.Error("Incorrect number of arguments. Expecting 1")
	 }
 
	 donationAsBytes, _ := APIstub.GetState(args[0])
	 return shim.Success(donationAsBytes)
 }
 
 /*
  * createDonation 是创建一条捐赠数据，并写入账本，每条捐赠数据包括5个数据内容：主键KEY、捐赠项目、捐赠人、捐赠内容、接收机构
  */
 func (s *DonationContract) createDonation(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 if len(args) != 5 {
		 return shim.Error("Incorrect number of arguments. Expecting 5")
	 }
 
	 var donation = Donation{Project: args[1], Donator: args[2], Content: args[3], Receiver: args[4]}
 
	 donationAsBytes, _ := json.Marshal(donation)
	 APIstub.PutState(args[0], donationAsBytes)
 
	 return shim.Success(nil)
 }
 
 /*
  * 主函数，需要调用shim.Start()方法，启动链码必须通过调用shim包中的Start函数实现
  */
 func main() {
 
	 // 创建一个新的Donation智能合约
	 err := shim.Start(new(DonationContract))
	 if err != nil {
		 fmt.Printf("Error creating new Donation Contract: %s", err)
	 }
 }
 