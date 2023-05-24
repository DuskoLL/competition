/*
Copyright IBM Corp. 2016 All Rights Reserved.

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

/*
 *本程序是实现账本中的实体A和实体B之间转账、查询余额、删除实体的操作。
*/

package main


/* Imports 是用来引入
 * 2个实体包用来格式化代码("fmt")和进行字符串处理("strconv")
 * 2个跟Hyperledger Fabric智能合约相关的包shim,peer
 */
import (
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// 定义智能合约结构体
type SimpleChaincode struct {
}
/*
 * Init方法是在智能合约初始化时由区块链网络调用的
 */
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Init")
	_, args := stub.GetFunctionAndParameters()  // 下划线忽略返回函数的值，args变量记录其他参数
	var A, B string  // 定义两个实体A和B
	var Aval, Bval int  // 定义两个实体A和B的初始余额Aval和Bval
	var err error

	// 检查args参数变量，数组长度必须为4，否则提示错误消息
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	// 初始化链码
	A = args[0]  // 获取实体A的值
	Aval, err = strconv.Atoi(args[1])  // 获取A的初始余额，如果不是整数，提示错误信息
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	B = args[2]  // 获取实体B的值
	Bval, err = strconv.Atoi(args[3])  // 获取B的初始余额，如果不是整数，提示错误信息
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)

	// 将状态值记录到分布式账本中，向账本中存入了两对键值，分别记录了A和B的余额
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))  // 向账本中存入A数据的键值对
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))  // 向账本中存入B数据的键值对
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)  // 返回状态为OK的消息
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function == "invoke" {
		// 从A向B转账
		return t.invoke(stub, args)
	} else if function == "delete" {
		// 从状态数据库中删除一个实体
		return t.delete(stub, args)
	} else if function == "query" {
		// 查询某实体余额
		return t.query(stub, args)
	}
     // 如果返回函数名称不对，提示只能调用invoke   delete  query这三个函数
	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}

// invoke实现从A向B转账的交易
func (t *SimpleChaincode) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var A, B string     // 定义两个实体A和B
	var Aval, Bval int  // 定义两个实体A和B的初始余额Aval和Bval
	var X int           // 转账交易金额
	var err error

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	A = args[0]
	B = args[1]

	// 得到实体的状态值
	Avalbytes, err := stub.GetState(A)  // 得到A的状态值
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Avalbytes == nil {
		return shim.Error("Entity not found")
	}
	Aval, _ = strconv.Atoi(string(Avalbytes))  // 将字符串转换为整型 

	Bvalbytes, err := stub.GetState(B)  // 得到B的状态值
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Bvalbytes == nil {
		return shim.Error("Entity not found")
	}
	Bval, _ = strconv.Atoi(string(Bvalbytes))  // 将字符串转换为整型

	// 执行转账交易
	X, err = strconv.Atoi(args[2])  // 将转账金额的字符串转换为整型
	if err != nil {
		return shim.Error("Invalid transaction amount, expecting a integer value")
	}
	Aval = Aval - X  // A从余额中减去转账金额
	Bval = Bval + X  // B从余额中加上转账金额
	fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)

	// 将A的更新值写入账本
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return shim.Error(err.Error())
	}
    // 将B的更新值写入账本
	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// delete实现从状态数据库中删除一个实体
func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	A := args[0]  // 得到想删除的键key

	// 从状态数据库中删除键key对应的键值,删除操作将作为交易存储在区块链上
	err := stub.DelState(A)
	if err != nil {
		return shim.Error("Failed to delete state")
	}

	return shim.Success(nil)
}

// query实现一个实体的余额查询
func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var A string  //定义实体A
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	A = args[0]

	// 得到实体A的状态值,即余额
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)  //输出查询余额的结果
	return shim.Success(Avalbytes)  // 返回状态为OK的消息,并将余额Avalbytes写入Response的Payload字段中
}
// 主函数,需要调用shim.Start()方法，启动链码必须通过调用shim包中的Start函数实现
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
