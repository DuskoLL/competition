// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package insurances

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
)

// InsurancesABI is the input ABI used to generate the binding from.
const InsurancesABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"insuranceIds\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"index\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_province\",\"type\":\"string\"},{\"name\":\"_city\",\"type\":\"string\"},{\"name\":\"_weather\",\"type\":\"string\"},{\"name\":\"_temperature\",\"type\":\"int8\"},{\"name\":\"_winddirection\",\"type\":\"string\"},{\"name\":\"_windpower\",\"type\":\"string\"},{\"name\":\"_humidity\",\"type\":\"int8\"}],\"name\":\"uploadWeathers\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_insuranceId\",\"type\":\"uint256\"}],\"name\":\"censorInsurance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"fallback\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"insurances\",\"outputs\":[{\"name\":\"insuranceId\",\"type\":\"uint256\"},{\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"name\":\"compensation\",\"type\":\"uint256\"},{\"name\":\"insuranceStatus\",\"type\":\"bool\"},{\"name\":\"censorStatues\",\"type\":\"bool\"},{\"name\":\"endTime\",\"type\":\"uint256\"},{\"name\":\"isExpired\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"receive\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"},{\"name\":\"_compenstation\",\"type\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"createInsurance\",\"outputs\":[{\"name\":\"insuranceId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_insuranceId\",\"type\":\"uint256\"}],\"name\":\"findInsurance\",\"outputs\":[{\"name\":\"insuranceId\",\"type\":\"uint256\"},{\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"name\":\"compensation\",\"type\":\"uint256\"},{\"name\":\"insuranceStatus\",\"type\":\"bool\"},{\"name\":\"censorStatues\",\"type\":\"bool\"},{\"name\":\"endTime\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_insuranceId\",\"type\":\"uint256\"}],\"name\":\"executeInsurance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isExist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"weatherInfos\",\"outputs\":[{\"name\":\"province\",\"type\":\"string\"},{\"name\":\"city\",\"type\":\"string\"},{\"name\":\"weather\",\"type\":\"string\"},{\"name\":\"temperature\",\"type\":\"int8\"},{\"name\":\"winddirection\",\"type\":\"string\"},{\"name\":\"windpower\",\"type\":\"string\"},{\"name\":\"humidity\",\"type\":\"int8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalCompensation\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_insuranceId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"NewInsurance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_insuranceId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"result\",\"type\":\"bool\"}],\"name\":\"CensorInsurance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_insuranceId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ExecuteInsurance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_insuranceId\",\"type\":\"uint256\"}],\"name\":\"InsuranceExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_insuranceId\",\"type\":\"uint256\"}],\"name\":\"InsuranceNotExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_adminAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"}]"

// InsurancesBin is the compiled bytecode used for deploying new contracts.
var InsurancesBin = "0x608060405233600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506125ac806100546000396000f3006080604052600436106100db576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806312065fe0146100e057806316b021e81461010b578063335932fc14610162578063356f36ca146101a3578063365eb19d146102585780633ccfd60b14610285578063552079dc1461029c57806359fb49f9146102a6578063a3e76c0f14610349578063a802bbf814610353578063b6ee48db146103be578063c72b678714610456578063ca8f8ff314610483578063ce9f29d9146104c8578063f3e9bc2814610738575b600080fd5b3480156100ec57600080fd5b506100f5610763565b6040518082815260200191505060405180910390f35b34801561011757600080fd5b5061014c600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610782565b6040518082815260200191505060405180910390f35b34801561016e57600080fd5b5061018d6004803603810190808035906020019092919050505061079a565b6040518082815260200191505060405180910390f35b3480156101af57600080fd5b50610256600480360381019080803590602001908201803590602001919091929391929390803590602001908201803590602001919091929391929390803590602001908201803590602001919091929391929390803560000b9060200190929190803590602001908201803590602001919091929391929390803590602001908201803590602001919091929391929390803560000b90602001909291905050506107b2565b005b34801561026457600080fd5b5061028360048036038101908080359060200190929190505050610a09565b005b34801561029157600080fd5b5061029a610de5565b005b6102a4611043565b005b3480156102b257600080fd5b506102d160048036038101908080359060200190929190505050611045565b604051808881526020018773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200186815260200185151515158152602001841515151581526020018381526020018215151515815260200197505050505050505060405180910390f35b6103516110dd565b005b34801561035f57600080fd5b506103a8600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001909291905050506110df565b6040518082815260200191505060405180910390f35b3480156103ca57600080fd5b506103e960048036038101908080359060200190929190505050611603565b604051808781526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018415151515815260200183151515158152602001828152602001965050505050505060405180910390f35b34801561046257600080fd5b50610481600480360381019080803590602001909291905050506117dc565b005b34801561048f57600080fd5b506104ae60048036038101908080359060200190929190505050611eca565b604051808215151515815260200191505060405180910390f35b3480156104d457600080fd5b506104f360048036038101908080359060200190929190505050611eea565b604051808060200180602001806020018860000b60000b815260200180602001806020018760000b60000b815260200186810386528d818151815260200191508051906020019080838360005b8381101561055b578082015181840152602081019050610540565b50505050905090810190601f1680156105885780820380516001836020036101000a031916815260200191505b5086810385528c818151815260200191508051906020019080838360005b838110156105c15780820151818401526020810190506105a6565b50505050905090810190601f1680156105ee5780820380516001836020036101000a031916815260200191505b5086810384528b818151815260200191508051906020019080838360005b8381101561062757808201518184015260208101905061060c565b50505050905090810190601f1680156106545780820380516001836020036101000a031916815260200191505b50868103835289818151815260200191508051906020019080838360005b8381101561068d578082015181840152602081019050610672565b50505050905090810190601f1680156106ba5780820380516001836020036101000a031916815260200191505b50868103825288818151815260200191508051906020019080838360005b838110156106f35780820151818401526020810190506106d8565b50505050905090810190601f1680156107205780820380516001836020036101000a031916815260200191505b509c5050505050505050505050505060405180910390f35b34801561074457600080fd5b5061074d61224d565b6040518082815260200191505060405180910390f35b60003073ffffffffffffffffffffffffffffffffffffffff1631905090565b60066020528060005260406000206000915090505481565b60046020528060005260406000206000915090505481565b6107ba61243f565b60e0604051908101604052808e8e8080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505081526020018c8c8080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505081526020018a8a8080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505081526020018860000b815260200187878080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050815260200185858080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505081526020018360000b815250905060018190806001815401808255809150509060018203906000526020600020906007020160009091929091909150600082015181600001908051906020019061093a929190612481565b506020820151816001019080519060200190610957929190612481565b506040820151816002019080519060200190610974929190612481565b5060608201518160030160006101000a81548160ff021916908360000b60ff16021790555060808201518160040190805190602001906109b5929190612481565b5060a08201518160050190805190602001906109d2929190612481565b5060c08201518160060160006101000a81548160ff021916908360000b60ff16021790555050505050505050505050505050505050565b60008082610a15612501565b6005600083815260200190815260200160002060009054906101000a900460ff161515610aaa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f696e737572616e6365206e6f742065786973740000000000000000000000000081525060200191505060405180910390fd5b60006004600084815260200190815260200160002054815481101515610acc57fe5b906000526020600020906006020160e06040519081016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015481526020016003820160009054906101000a900460ff161515151581526020016003820160019054906101000a900460ff16151515158152602001600482015481526020016005820160009054906101000a900460ff1615151515815250509050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480610c385750806020015173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b1515610cac576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f6e6f742041646d696e206f722062656e6566696369617279000000000000000081525060200191505060405180910390fd5b610cb585612253565b1515610cc057610dde565b60006004600087815260200190815260200160002054815481101515610ce257fe5b906000526020600020906006020193508360030160009054906101000a900460ff16151515610d79576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f696e737572616e6365206973206578637465640000000000000000000000000081525060200191505060405180910390fd5b600192508215610da15760018460030160016101000a81548160ff0219169083151502179055505b847f6eb9ead4eb17be8377128ca9817d26b1caaee0d335e819f4303c2b6766468aa284604051808215151515815260200191505060405180910390a25b5050505050565b600080600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610ead576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260098152602001807f6e6f742061646d696e000000000000000000000000000000000000000000000081525060200191505060405180910390fd5b3073ffffffffffffffffffffffffffffffffffffffff1631915081600354101515610f40576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f63616e206e6f742077697468647261770000000000000000000000000000000081525060200191505060405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff16600354830360405180602001905060006040518083038185875af1925050509050801515610fed576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600b8152602001807f6661696c65642063616c6c00000000000000000000000000000000000000000081525060200191505060405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a942436460035484036040518082815260200191505060405180910390a25050565b565b60008181548110151561105457fe5b90600052602060002090600602016000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020154908060030160009054906101000a900460ff16908060030160019054906101000a900460ff16908060040154908060050160009054906101000a900460ff16905087565b565b60008060006110ec612501565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156111b1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260098152602001807f6e6f742061646d696e000000000000000000000000000000000000000000000081525060200191505060405180910390fd5b3073ffffffffffffffffffffffffffffffffffffffff163160035411151515611268576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001807f6163636f756e7442616c616e6365203c20746f74616c436f6d70656e7361746981526020017f6f6e00000000000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b8442019250868684604051602001808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200183815260200182815260200193505050506040516020818303038152906040526040518082805190602001908083835b6020831015156112fe57805182526020820191506020810190506020830392506112d9565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390206001900491506005600083815260200190815260200160002060009054906101000a900460ff161515156113c7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f696e737572616e6365206578697374000000000000000000000000000000000081525060200191505060405180910390fd5b60e0604051908101604052808381526020018873ffffffffffffffffffffffffffffffffffffffff168152602001878152602001600015158152602001600015158152602001848152602001600015158152509050600081908060018154018082558091505090600182039060005260206000209060060201600090919290919091506000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002015560608201518160030160006101000a81548160ff02191690831515021790555060808201518160030160016101000a81548160ff02191690831515021790555060a0820151816004015560c08201518160050160006101000a81548160ff02191690831515021790555050505060016005600084815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600080549050036004600084815260200190815260200160002081905550856003600082825401925050819055508673ffffffffffffffffffffffffffffffffffffffff16827f547962ec7fac5a8cc9580e652dc73617291177603313596daef886df1a27f8bc60405160405180910390a381600660008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508193505050509392505050565b600080600080600080611614612501565b6005600089815260200190815260200160002060009054906101000a900460ff1615156116a9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f696e737572616e636520646f206e6f742065786973740000000000000000000081525060200191505060405180910390fd5b6000600460008a8152602001908152602001600020548154811015156116cb57fe5b906000526020600020906006020160e06040519081016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015481526020016003820160009054906101000a900460ff161515151581526020016003820160019054906101000a900460ff16151515158152602001600482015481526020016005820160009054906101000a900460ff1615151515815250509050879650806020015195508060400151945080606001519350806080015192508060a0015191505091939550919395565b600080600080846117eb612501565b6005600083815260200190815260200160002060009054906101000a900460ff161515611880576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f696e737572616e6365206e6f742065786973740000000000000000000000000081525060200191505060405180910390fd5b600060046000848152602001908152602001600020548154811015156118a257fe5b906000526020600020906006020160e06040519081016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015481526020016003820160009054906101000a900460ff161515151581526020016003820160019054906101000a900460ff16151515158152602001600482015481526020016005820160009054906101000a900460ff1615151515815250509050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480611a0e5750806020015173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b1515611a82576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f6e6f742041646d696e206f722062656e6566696369617279000000000000000081525060200191505060405180910390fd5b611a8b87612253565b1515611a9657611ec1565b60006004600089815260200190815260200160002054815481101515611ab857fe5b90600052602060002090600602019550600260149054906101000a900460ff16151515611b4d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f6c6f636b206973206c6f636b656400000000000000000000000000000000000081525060200191505060405180910390fd5b6001600260146101000a81548160ff0219169083151502179055508560030160009054906101000a900460ff1615611c68578660405160200180828152602001807f2069732065786375746564000000000000000000000000000000000000000000815250600b019150506040516020818303038152906040526040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611c2d578082015181840152602081019050611c12565b50505050905090810190601f168015611c5a5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b8560030160019054906101000a900460ff161515611d69578660405160200180828152602001807f206973206e6f742063656e736f726564000000000000000000000000000000008152506010019150506040516020818303038152906040526040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611d2e578082015181840152602081019050611d13565b50505050905090810190601f168015611d5b5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b856002015494508560010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1693508373ffffffffffffffffffffffffffffffffffffffff168560405180602001905060006040518083038185875af1925050509250821515611e40576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600b8152602001807f6661696c65642063616c6c00000000000000000000000000000000000000000081525060200191505060405180910390fd5b60018660030160006101000a81548160ff021916908315150217905550846003600082825403925050819055506000600260146101000a81548160ff021916908315150217905550867f6f7e18f0d8960445e4ff55b876099f6d73f888232d5e5c91848247d3beadb50f866040518082815260200191505060405180910390a25b50505050505050565b60056020528060005260406000206000915054906101000a900460ff1681565b600181815481101515611ef957fe5b9060005260206000209060070201600091509050806000018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611fa55780601f10611f7a57610100808354040283529160200191611fa5565b820191906000526020600020905b815481529060010190602001808311611f8857829003601f168201915b505050505090806001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156120435780601f1061201857610100808354040283529160200191612043565b820191906000526020600020905b81548152906001019060200180831161202657829003601f168201915b505050505090806002018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156120e15780601f106120b6576101008083540402835291602001916120e1565b820191906000526020600020905b8154815290600101906020018083116120c457829003601f168201915b5050505050908060030160009054906101000a900460000b90806004018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156121925780601f1061216757610100808354040283529160200191612192565b820191906000526020600020905b81548152906001019060200180831161217557829003601f168201915b505050505090806005018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156122305780601f1061220557610100808354040283529160200191612230565b820191906000526020600020905b81548152906001019060200180831161221357829003601f168201915b5050505050908060060160009054906101000a900460000b905087565b60035481565b600061225d612501565b6000600460008581526020019081526020016000205481548110151561227f57fe5b906000526020600020906006020160e06040519081016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015481526020016003820160009054906101000a900460ff161515151581526020016003820160019054906101000a900460ff16151515158152602001600482015481526020016005820160009054906101000a900460ff16151515158152505090508060c00151156123a057827fe18b4e336f36a5b2e805a39d09aaaf50084b53f5035295d2918900cf75155e6960405160405180910390a260009150612439565b428160a0015110156124075760018160c00190151590811515815250508060400151600360008282540392505081905550827fe18b4e336f36a5b2e805a39d09aaaf50084b53f5035295d2918900cf75155e6960405160405180910390a260009150612439565b827f8142336de75476de6d07bfe34897b339001e8a87a41b3e09ff071a9a8be7079b60405160405180910390a2600191505b50919050565b60e0604051908101604052806060815260200160608152602001606081526020016000800b815260200160608152602001606081526020016000800b81525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106124c257805160ff19168380011785556124f0565b828001600101855582156124f0579182015b828111156124ef5782518255916020019190600101906124d4565b5b5090506124fd919061255b565b5090565b60e06040519081016040528060008152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600015158152602001600015158152602001600081526020016000151581525090565b61257d91905b80821115612579576000816000905550600101612561565b5090565b905600a165627a7a7230582012a12787ae1e849489c4f135c436fed084f9cd89300a38c43c9a4c4621d145960029"

// DeployInsurances deploys a new contract, binding an instance of Insurances to it.
func DeployInsurances(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Insurances, error) {
	parsed, err := abi.JSON(strings.NewReader(InsurancesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(InsurancesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Insurances{InsurancesCaller: InsurancesCaller{contract: contract}, InsurancesTransactor: InsurancesTransactor{contract: contract}, InsurancesFilterer: InsurancesFilterer{contract: contract}}, nil
}

func AsyncDeployInsurances(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(InsurancesABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(InsurancesBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Insurances is an auto generated Go binding around a Solidity contract.
type Insurances struct {
	InsurancesCaller     // Read-only binding to the contract
	InsurancesTransactor // Write-only binding to the contract
	InsurancesFilterer   // Log filterer for contract events
}

// InsurancesCaller is an auto generated read-only Go binding around a Solidity contract.
type InsurancesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InsurancesTransactor is an auto generated write-only Go binding around a Solidity contract.
type InsurancesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InsurancesFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type InsurancesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InsurancesSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type InsurancesSession struct {
	Contract     *Insurances       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InsurancesCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type InsurancesCallerSession struct {
	Contract *InsurancesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// InsurancesTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type InsurancesTransactorSession struct {
	Contract     *InsurancesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// InsurancesRaw is an auto generated low-level Go binding around a Solidity contract.
type InsurancesRaw struct {
	Contract *Insurances // Generic contract binding to access the raw methods on
}

// InsurancesCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type InsurancesCallerRaw struct {
	Contract *InsurancesCaller // Generic read-only contract binding to access the raw methods on
}

// InsurancesTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type InsurancesTransactorRaw struct {
	Contract *InsurancesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInsurances creates a new instance of Insurances, bound to a specific deployed contract.
func NewInsurances(address common.Address, backend bind.ContractBackend) (*Insurances, error) {
	contract, err := bindInsurances(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Insurances{InsurancesCaller: InsurancesCaller{contract: contract}, InsurancesTransactor: InsurancesTransactor{contract: contract}, InsurancesFilterer: InsurancesFilterer{contract: contract}}, nil
}

// NewInsurancesCaller creates a new read-only instance of Insurances, bound to a specific deployed contract.
func NewInsurancesCaller(address common.Address, caller bind.ContractCaller) (*InsurancesCaller, error) {
	contract, err := bindInsurances(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InsurancesCaller{contract: contract}, nil
}

// NewInsurancesTransactor creates a new write-only instance of Insurances, bound to a specific deployed contract.
func NewInsurancesTransactor(address common.Address, transactor bind.ContractTransactor) (*InsurancesTransactor, error) {
	contract, err := bindInsurances(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InsurancesTransactor{contract: contract}, nil
}

// NewInsurancesFilterer creates a new log filterer instance of Insurances, bound to a specific deployed contract.
func NewInsurancesFilterer(address common.Address, filterer bind.ContractFilterer) (*InsurancesFilterer, error) {
	contract, err := bindInsurances(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InsurancesFilterer{contract: contract}, nil
}

// bindInsurances binds a generic wrapper to an already deployed contract.
func bindInsurances(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InsurancesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Insurances *InsurancesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Insurances.Contract.InsurancesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Insurances *InsurancesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Insurances.Contract.InsurancesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Insurances *InsurancesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Insurances.Contract.InsurancesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Insurances *InsurancesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Insurances.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Insurances *InsurancesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Insurances.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Insurances *InsurancesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Insurances.Contract.contract.Transact(opts, method, params...)
}

// FindInsurance is a free data retrieval call binding the contract method 0xb6ee48db.
//
// Solidity: function findInsurance(uint256 _insuranceId) constant returns(uint256 insuranceId, address beneficiaryAddress, uint256 compensation, bool insuranceStatus, bool censorStatues, uint256 endTime)
func (_Insurances *InsurancesCaller) FindInsurance(opts *bind.CallOpts, _insuranceId *big.Int) (struct {
	InsuranceId        *big.Int
	BeneficiaryAddress common.Address
	Compensation       *big.Int
	InsuranceStatus    bool
	CensorStatues      bool
	EndTime            *big.Int
}, error) {
	ret := new(struct {
		InsuranceId        *big.Int
		BeneficiaryAddress common.Address
		Compensation       *big.Int
		InsuranceStatus    bool
		CensorStatues      bool
		EndTime            *big.Int
	})
	out := ret
	err := _Insurances.contract.Call(opts, out, "findInsurance", _insuranceId)
	return *ret, err
}

// FindInsurance is a free data retrieval call binding the contract method 0xb6ee48db.
//
// Solidity: function findInsurance(uint256 _insuranceId) constant returns(uint256 insuranceId, address beneficiaryAddress, uint256 compensation, bool insuranceStatus, bool censorStatues, uint256 endTime)
func (_Insurances *InsurancesSession) FindInsurance(_insuranceId *big.Int) (struct {
	InsuranceId        *big.Int
	BeneficiaryAddress common.Address
	Compensation       *big.Int
	InsuranceStatus    bool
	CensorStatues      bool
	EndTime            *big.Int
}, error) {
	return _Insurances.Contract.FindInsurance(&_Insurances.CallOpts, _insuranceId)
}

// FindInsurance is a free data retrieval call binding the contract method 0xb6ee48db.
//
// Solidity: function findInsurance(uint256 _insuranceId) constant returns(uint256 insuranceId, address beneficiaryAddress, uint256 compensation, bool insuranceStatus, bool censorStatues, uint256 endTime)
func (_Insurances *InsurancesCallerSession) FindInsurance(_insuranceId *big.Int) (struct {
	InsuranceId        *big.Int
	BeneficiaryAddress common.Address
	Compensation       *big.Int
	InsuranceStatus    bool
	CensorStatues      bool
	EndTime            *big.Int
}, error) {
	return _Insurances.Contract.FindInsurance(&_Insurances.CallOpts, _insuranceId)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() constant returns(uint256)
func (_Insurances *InsurancesCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Insurances.contract.Call(opts, out, "getBalance")
	return *ret0, err
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() constant returns(uint256)
func (_Insurances *InsurancesSession) GetBalance() (*big.Int, error) {
	return _Insurances.Contract.GetBalance(&_Insurances.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() constant returns(uint256)
func (_Insurances *InsurancesCallerSession) GetBalance() (*big.Int, error) {
	return _Insurances.Contract.GetBalance(&_Insurances.CallOpts)
}

// Index is a free data retrieval call binding the contract method 0x335932fc.
//
// Solidity: function index(uint256 ) constant returns(uint256)
func (_Insurances *InsurancesCaller) Index(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Insurances.contract.Call(opts, out, "index", arg0)
	return *ret0, err
}

// Index is a free data retrieval call binding the contract method 0x335932fc.
//
// Solidity: function index(uint256 ) constant returns(uint256)
func (_Insurances *InsurancesSession) Index(arg0 *big.Int) (*big.Int, error) {
	return _Insurances.Contract.Index(&_Insurances.CallOpts, arg0)
}

// Index is a free data retrieval call binding the contract method 0x335932fc.
//
// Solidity: function index(uint256 ) constant returns(uint256)
func (_Insurances *InsurancesCallerSession) Index(arg0 *big.Int) (*big.Int, error) {
	return _Insurances.Contract.Index(&_Insurances.CallOpts, arg0)
}

// InsuranceIds is a free data retrieval call binding the contract method 0x16b021e8.
//
// Solidity: function insuranceIds(address ) constant returns(uint256)
func (_Insurances *InsurancesCaller) InsuranceIds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Insurances.contract.Call(opts, out, "insuranceIds", arg0)
	return *ret0, err
}

// InsuranceIds is a free data retrieval call binding the contract method 0x16b021e8.
//
// Solidity: function insuranceIds(address ) constant returns(uint256)
func (_Insurances *InsurancesSession) InsuranceIds(arg0 common.Address) (*big.Int, error) {
	return _Insurances.Contract.InsuranceIds(&_Insurances.CallOpts, arg0)
}

// InsuranceIds is a free data retrieval call binding the contract method 0x16b021e8.
//
// Solidity: function insuranceIds(address ) constant returns(uint256)
func (_Insurances *InsurancesCallerSession) InsuranceIds(arg0 common.Address) (*big.Int, error) {
	return _Insurances.Contract.InsuranceIds(&_Insurances.CallOpts, arg0)
}

// Insurances is a free data retrieval call binding the contract method 0x59fb49f9.
//
// Solidity: function insurances(uint256 ) constant returns(uint256 insuranceId, address beneficiaryAddress, uint256 compensation, bool insuranceStatus, bool censorStatues, uint256 endTime, bool isExpired)
func (_Insurances *InsurancesCaller) Insurances(opts *bind.CallOpts, arg0 *big.Int) (struct {
	InsuranceId        *big.Int
	BeneficiaryAddress common.Address
	Compensation       *big.Int
	InsuranceStatus    bool
	CensorStatues      bool
	EndTime            *big.Int
	IsExpired          bool
}, error) {
	ret := new(struct {
		InsuranceId        *big.Int
		BeneficiaryAddress common.Address
		Compensation       *big.Int
		InsuranceStatus    bool
		CensorStatues      bool
		EndTime            *big.Int
		IsExpired          bool
	})
	out := ret
	err := _Insurances.contract.Call(opts, out, "insurances", arg0)
	return *ret, err
}

// Insurances is a free data retrieval call binding the contract method 0x59fb49f9.
//
// Solidity: function insurances(uint256 ) constant returns(uint256 insuranceId, address beneficiaryAddress, uint256 compensation, bool insuranceStatus, bool censorStatues, uint256 endTime, bool isExpired)
func (_Insurances *InsurancesSession) Insurances(arg0 *big.Int) (struct {
	InsuranceId        *big.Int
	BeneficiaryAddress common.Address
	Compensation       *big.Int
	InsuranceStatus    bool
	CensorStatues      bool
	EndTime            *big.Int
	IsExpired          bool
}, error) {
	return _Insurances.Contract.Insurances(&_Insurances.CallOpts, arg0)
}

// Insurances is a free data retrieval call binding the contract method 0x59fb49f9.
//
// Solidity: function insurances(uint256 ) constant returns(uint256 insuranceId, address beneficiaryAddress, uint256 compensation, bool insuranceStatus, bool censorStatues, uint256 endTime, bool isExpired)
func (_Insurances *InsurancesCallerSession) Insurances(arg0 *big.Int) (struct {
	InsuranceId        *big.Int
	BeneficiaryAddress common.Address
	Compensation       *big.Int
	InsuranceStatus    bool
	CensorStatues      bool
	EndTime            *big.Int
	IsExpired          bool
}, error) {
	return _Insurances.Contract.Insurances(&_Insurances.CallOpts, arg0)
}

// IsExist is a free data retrieval call binding the contract method 0xca8f8ff3.
//
// Solidity: function isExist(uint256 ) constant returns(bool)
func (_Insurances *InsurancesCaller) IsExist(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Insurances.contract.Call(opts, out, "isExist", arg0)
	return *ret0, err
}

// IsExist is a free data retrieval call binding the contract method 0xca8f8ff3.
//
// Solidity: function isExist(uint256 ) constant returns(bool)
func (_Insurances *InsurancesSession) IsExist(arg0 *big.Int) (bool, error) {
	return _Insurances.Contract.IsExist(&_Insurances.CallOpts, arg0)
}

// IsExist is a free data retrieval call binding the contract method 0xca8f8ff3.
//
// Solidity: function isExist(uint256 ) constant returns(bool)
func (_Insurances *InsurancesCallerSession) IsExist(arg0 *big.Int) (bool, error) {
	return _Insurances.Contract.IsExist(&_Insurances.CallOpts, arg0)
}

// TotalCompensation is a free data retrieval call binding the contract method 0xf3e9bc28.
//
// Solidity: function totalCompensation() constant returns(uint256)
func (_Insurances *InsurancesCaller) TotalCompensation(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Insurances.contract.Call(opts, out, "totalCompensation")
	return *ret0, err
}

// TotalCompensation is a free data retrieval call binding the contract method 0xf3e9bc28.
//
// Solidity: function totalCompensation() constant returns(uint256)
func (_Insurances *InsurancesSession) TotalCompensation() (*big.Int, error) {
	return _Insurances.Contract.TotalCompensation(&_Insurances.CallOpts)
}

// TotalCompensation is a free data retrieval call binding the contract method 0xf3e9bc28.
//
// Solidity: function totalCompensation() constant returns(uint256)
func (_Insurances *InsurancesCallerSession) TotalCompensation() (*big.Int, error) {
	return _Insurances.Contract.TotalCompensation(&_Insurances.CallOpts)
}

// WeatherInfos is a free data retrieval call binding the contract method 0xce9f29d9.
//
// Solidity: function weatherInfos(uint256 ) constant returns(string province, string city, string weather, int8 temperature, string winddirection, string windpower, int8 humidity)
func (_Insurances *InsurancesCaller) WeatherInfos(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Province      string
	City          string
	Weather       string
	Temperature   int8
	Winddirection string
	Windpower     string
	Humidity      int8
}, error) {
	ret := new(struct {
		Province      string
		City          string
		Weather       string
		Temperature   int8
		Winddirection string
		Windpower     string
		Humidity      int8
	})
	out := ret
	err := _Insurances.contract.Call(opts, out, "weatherInfos", arg0)
	return *ret, err
}

// WeatherInfos is a free data retrieval call binding the contract method 0xce9f29d9.
//
// Solidity: function weatherInfos(uint256 ) constant returns(string province, string city, string weather, int8 temperature, string winddirection, string windpower, int8 humidity)
func (_Insurances *InsurancesSession) WeatherInfos(arg0 *big.Int) (struct {
	Province      string
	City          string
	Weather       string
	Temperature   int8
	Winddirection string
	Windpower     string
	Humidity      int8
}, error) {
	return _Insurances.Contract.WeatherInfos(&_Insurances.CallOpts, arg0)
}

// WeatherInfos is a free data retrieval call binding the contract method 0xce9f29d9.
//
// Solidity: function weatherInfos(uint256 ) constant returns(string province, string city, string weather, int8 temperature, string winddirection, string windpower, int8 humidity)
func (_Insurances *InsurancesCallerSession) WeatherInfos(arg0 *big.Int) (struct {
	Province      string
	City          string
	Weather       string
	Temperature   int8
	Winddirection string
	Windpower     string
	Humidity      int8
}, error) {
	return _Insurances.Contract.WeatherInfos(&_Insurances.CallOpts, arg0)
}

// CensorInsurance is a paid mutator transaction binding the contract method 0x365eb19d.
//
// Solidity: function censorInsurance(uint256 _insuranceId) returns()
func (_Insurances *InsurancesTransactor) CensorInsurance(opts *bind.TransactOpts, _insuranceId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Insurances.contract.Transact(opts, "censorInsurance", _insuranceId)
}

func (_Insurances *InsurancesTransactor) AsyncCensorInsurance(handler func(*types.Receipt, error), opts *bind.TransactOpts, _insuranceId *big.Int) (*types.Transaction, error) {
	return _Insurances.contract.AsyncTransact(opts, handler, "censorInsurance", _insuranceId)
}

// CensorInsurance is a paid mutator transaction binding the contract method 0x365eb19d.
//
// Solidity: function censorInsurance(uint256 _insuranceId) returns()
func (_Insurances *InsurancesSession) CensorInsurance(_insuranceId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Insurances.Contract.CensorInsurance(&_Insurances.TransactOpts, _insuranceId)
}

func (_Insurances *InsurancesSession) AsyncCensorInsurance(handler func(*types.Receipt, error), _insuranceId *big.Int) (*types.Transaction, error) {
	return _Insurances.Contract.AsyncCensorInsurance(handler, &_Insurances.TransactOpts, _insuranceId)
}

// CensorInsurance is a paid mutator transaction binding the contract method 0x365eb19d.
//
// Solidity: function censorInsurance(uint256 _insuranceId) returns()
func (_Insurances *InsurancesTransactorSession) CensorInsurance(_insuranceId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Insurances.Contract.CensorInsurance(&_Insurances.TransactOpts, _insuranceId)
}

func (_Insurances *InsurancesTransactorSession) AsyncCensorInsurance(handler func(*types.Receipt, error), _insuranceId *big.Int) (*types.Transaction, error) {
	return _Insurances.Contract.AsyncCensorInsurance(handler, &_Insurances.TransactOpts, _insuranceId)
}

// CreateInsurance is a paid mutator transaction binding the contract method 0xa802bbf8.
//
// Solidity: function createInsurance(address _beneficiary, uint256 _compenstation, uint256 duration) returns(uint256 insuranceId)
func (_Insurances *InsurancesTransactor) CreateInsurance(opts *bind.TransactOpts, _beneficiary common.Address, _compenstation *big.Int, duration *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Insurances.contract.Transact(opts, "createInsurance", _beneficiary, _compenstation, duration)
}

func (_Insurances *InsurancesTransactor) AsyncCreateInsurance(handler func(*types.Receipt, error), opts *bind.TransactOpts, _beneficiary common.Address, _compenstation *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Insurances.contract.AsyncTransact(opts, handler, "createInsurance", _beneficiary, _compenstation, duration)
}

// CreateInsurance is a paid mutator transaction binding the contract method 0xa802bbf8.
//
// Solidity: function createInsurance(address _beneficiary, uint256 _compenstation, uint256 duration) returns(uint256 insuranceId)
func (_Insurances *InsurancesSession) CreateInsurance(_beneficiary common.Address, _compenstation *big.Int, duration *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Insurances.Contract.CreateInsurance(&_Insurances.TransactOpts, _beneficiary, _compenstation, duration)
}

func (_Insurances *InsurancesSession) AsyncCreateInsurance(handler func(*types.Receipt, error), _beneficiary common.Address, _compenstation *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Insurances.Contract.AsyncCreateInsurance(handler, &_Insurances.TransactOpts, _beneficiary, _compenstation, duration)
}

// CreateInsurance is a paid mutator transaction binding the contract method 0xa802bbf8.
//
// Solidity: function createInsurance(address _beneficiary, uint256 _compenstation, uint256 duration) returns(uint256 insuranceId)
func (_Insurances *InsurancesTransactorSession) CreateInsurance(_beneficiary common.Address, _compenstation *big.Int, duration *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Insurances.Contract.CreateInsurance(&_Insurances.TransactOpts, _beneficiary, _compenstation, duration)
}

func (_Insurances *InsurancesTransactorSession) AsyncCreateInsurance(handler func(*types.Receipt, error), _beneficiary common.Address, _compenstation *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Insurances.Contract.AsyncCreateInsurance(handler, &_Insurances.TransactOpts, _beneficiary, _compenstation, duration)
}

// ExecuteInsurance is a paid mutator transaction binding the contract method 0xc72b6787.
//
// Solidity: function executeInsurance(uint256 _insuranceId) returns()
func (_Insurances *InsurancesTransactor) ExecuteInsurance(opts *bind.TransactOpts, _insuranceId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Insurances.contract.Transact(opts, "executeInsurance", _insuranceId)
}

func (_Insurances *InsurancesTransactor) AsyncExecuteInsurance(handler func(*types.Receipt, error), opts *bind.TransactOpts, _insuranceId *big.Int) (*types.Transaction, error) {
	return _Insurances.contract.AsyncTransact(opts, handler, "executeInsurance", _insuranceId)
}

// ExecuteInsurance is a paid mutator transaction binding the contract method 0xc72b6787.
//
// Solidity: function executeInsurance(uint256 _insuranceId) returns()
func (_Insurances *InsurancesSession) ExecuteInsurance(_insuranceId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Insurances.Contract.ExecuteInsurance(&_Insurances.TransactOpts, _insuranceId)
}

func (_Insurances *InsurancesSession) AsyncExecuteInsurance(handler func(*types.Receipt, error), _insuranceId *big.Int) (*types.Transaction, error) {
	return _Insurances.Contract.AsyncExecuteInsurance(handler, &_Insurances.TransactOpts, _insuranceId)
}

// ExecuteInsurance is a paid mutator transaction binding the contract method 0xc72b6787.
//
// Solidity: function executeInsurance(uint256 _insuranceId) returns()
func (_Insurances *InsurancesTransactorSession) ExecuteInsurance(_insuranceId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Insurances.Contract.ExecuteInsurance(&_Insurances.TransactOpts, _insuranceId)
}

func (_Insurances *InsurancesTransactorSession) AsyncExecuteInsurance(handler func(*types.Receipt, error), _insuranceId *big.Int) (*types.Transaction, error) {
	return _Insurances.Contract.AsyncExecuteInsurance(handler, &_Insurances.TransactOpts, _insuranceId)
}

// Fallback is a paid mutator transaction binding the contract method 0x552079dc.
//
// Solidity: function fallback() returns()
func (_Insurances *InsurancesTransactor) Fallback(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Insurances.contract.Transact(opts, "fallback")
}

func (_Insurances *InsurancesTransactor) AsyncFallback(handler func(*types.Receipt, error), opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Insurances.contract.AsyncTransact(opts, handler, "fallback")
}

// Fallback is a paid mutator transaction binding the contract method 0x552079dc.
//
// Solidity: function fallback() returns()
func (_Insurances *InsurancesSession) Fallback() (*types.Transaction, *types.Receipt, error) {
	return _Insurances.Contract.Fallback(&_Insurances.TransactOpts)
}

func (_Insurances *InsurancesSession) AsyncFallback(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _Insurances.Contract.AsyncFallback(handler, &_Insurances.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract method 0x552079dc.
//
// Solidity: function fallback() returns()
func (_Insurances *InsurancesTransactorSession) Fallback() (*types.Transaction, *types.Receipt, error) {
	return _Insurances.Contract.Fallback(&_Insurances.TransactOpts)
}

func (_Insurances *InsurancesTransactorSession) AsyncFallback(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _Insurances.Contract.AsyncFallback(handler, &_Insurances.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract method 0xa3e76c0f.
//
// Solidity: function receive() returns()
func (_Insurances *InsurancesTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Insurances.contract.Transact(opts, "receive")
}

func (_Insurances *InsurancesTransactor) AsyncReceive(handler func(*types.Receipt, error), opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Insurances.contract.AsyncTransact(opts, handler, "receive")
}

// Receive is a paid mutator transaction binding the contract method 0xa3e76c0f.
//
// Solidity: function receive() returns()
func (_Insurances *InsurancesSession) Receive() (*types.Transaction, *types.Receipt, error) {
	return _Insurances.Contract.Receive(&_Insurances.TransactOpts)
}

func (_Insurances *InsurancesSession) AsyncReceive(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _Insurances.Contract.AsyncReceive(handler, &_Insurances.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract method 0xa3e76c0f.
//
// Solidity: function receive() returns()
func (_Insurances *InsurancesTransactorSession) Receive() (*types.Transaction, *types.Receipt, error) {
	return _Insurances.Contract.Receive(&_Insurances.TransactOpts)
}

func (_Insurances *InsurancesTransactorSession) AsyncReceive(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _Insurances.Contract.AsyncReceive(handler, &_Insurances.TransactOpts)
}

// UploadWeathers is a paid mutator transaction binding the contract method 0x356f36ca.
//
// Solidity: function uploadWeathers(string _province, string _city, string _weather, int8 _temperature, string _winddirection, string _windpower, int8 _humidity) returns()
func (_Insurances *InsurancesTransactor) UploadWeathers(opts *bind.TransactOpts, _province string, _city string, _weather string, _temperature int8, _winddirection string, _windpower string, _humidity int8) (*types.Transaction, *types.Receipt, error) {
	return _Insurances.contract.Transact(opts, "uploadWeathers", _province, _city, _weather, _temperature, _winddirection, _windpower, _humidity)
}

func (_Insurances *InsurancesTransactor) AsyncUploadWeathers(handler func(*types.Receipt, error), opts *bind.TransactOpts, _province string, _city string, _weather string, _temperature int8, _winddirection string, _windpower string, _humidity int8) (*types.Transaction, error) {
	return _Insurances.contract.AsyncTransact(opts, handler, "uploadWeathers", _province, _city, _weather, _temperature, _winddirection, _windpower, _humidity)
}

// UploadWeathers is a paid mutator transaction binding the contract method 0x356f36ca.
//
// Solidity: function uploadWeathers(string _province, string _city, string _weather, int8 _temperature, string _winddirection, string _windpower, int8 _humidity) returns()
func (_Insurances *InsurancesSession) UploadWeathers(_province string, _city string, _weather string, _temperature int8, _winddirection string, _windpower string, _humidity int8) (*types.Transaction, *types.Receipt, error) {
	return _Insurances.Contract.UploadWeathers(&_Insurances.TransactOpts, _province, _city, _weather, _temperature, _winddirection, _windpower, _humidity)
}

func (_Insurances *InsurancesSession) AsyncUploadWeathers(handler func(*types.Receipt, error), _province string, _city string, _weather string, _temperature int8, _winddirection string, _windpower string, _humidity int8) (*types.Transaction, error) {
	return _Insurances.Contract.AsyncUploadWeathers(handler, &_Insurances.TransactOpts, _province, _city, _weather, _temperature, _winddirection, _windpower, _humidity)
}

// UploadWeathers is a paid mutator transaction binding the contract method 0x356f36ca.
//
// Solidity: function uploadWeathers(string _province, string _city, string _weather, int8 _temperature, string _winddirection, string _windpower, int8 _humidity) returns()
func (_Insurances *InsurancesTransactorSession) UploadWeathers(_province string, _city string, _weather string, _temperature int8, _winddirection string, _windpower string, _humidity int8) (*types.Transaction, *types.Receipt, error) {
	return _Insurances.Contract.UploadWeathers(&_Insurances.TransactOpts, _province, _city, _weather, _temperature, _winddirection, _windpower, _humidity)
}

func (_Insurances *InsurancesTransactorSession) AsyncUploadWeathers(handler func(*types.Receipt, error), _province string, _city string, _weather string, _temperature int8, _winddirection string, _windpower string, _humidity int8) (*types.Transaction, error) {
	return _Insurances.Contract.AsyncUploadWeathers(handler, &_Insurances.TransactOpts, _province, _city, _weather, _temperature, _winddirection, _windpower, _humidity)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Insurances *InsurancesTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Insurances.contract.Transact(opts, "withdraw")
}

func (_Insurances *InsurancesTransactor) AsyncWithdraw(handler func(*types.Receipt, error), opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Insurances.contract.AsyncTransact(opts, handler, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Insurances *InsurancesSession) Withdraw() (*types.Transaction, *types.Receipt, error) {
	return _Insurances.Contract.Withdraw(&_Insurances.TransactOpts)
}

func (_Insurances *InsurancesSession) AsyncWithdraw(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _Insurances.Contract.AsyncWithdraw(handler, &_Insurances.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Insurances *InsurancesTransactorSession) Withdraw() (*types.Transaction, *types.Receipt, error) {
	return _Insurances.Contract.Withdraw(&_Insurances.TransactOpts)
}

func (_Insurances *InsurancesTransactorSession) AsyncWithdraw(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _Insurances.Contract.AsyncWithdraw(handler, &_Insurances.TransactOpts)
}

// InsurancesCensorInsurance represents a CensorInsurance event raised by the Insurances contract.
type InsurancesCensorInsurance struct {
	InsuranceId *big.Int
	Result      bool
	Raw         types.Log // Blockchain specific contextual infos
}

// WatchCensorInsurance is a free log subscription operation binding the contract event 0x6eb9ead4eb17be8377128ca9817d26b1caaee0d335e819f4303c2b6766468aa2.
//
// Solidity: event CensorInsurance(uint256 indexed _insuranceId, bool result)
func (_Insurances *InsurancesFilterer) WatchCensorInsurance(fromBlock *uint64, handler func(int, []types.Log), _insuranceId *big.Int) error {
	return _Insurances.contract.WatchLogs(fromBlock, handler, "CensorInsurance", _insuranceId)
}

func (_Insurances *InsurancesFilterer) WatchAllCensorInsurance(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Insurances.contract.WatchLogs(fromBlock, handler, "CensorInsurance")
}

// ParseCensorInsurance is a log parse operation binding the contract event 0x6eb9ead4eb17be8377128ca9817d26b1caaee0d335e819f4303c2b6766468aa2.
//
// Solidity: event CensorInsurance(uint256 indexed _insuranceId, bool result)
func (_Insurances *InsurancesFilterer) ParseCensorInsurance(log types.Log) (*InsurancesCensorInsurance, error) {
	event := new(InsurancesCensorInsurance)
	if err := _Insurances.contract.UnpackLog(event, "CensorInsurance", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchCensorInsurance is a free log subscription operation binding the contract event 0x6eb9ead4eb17be8377128ca9817d26b1caaee0d335e819f4303c2b6766468aa2.
//
// Solidity: event CensorInsurance(uint256 indexed _insuranceId, bool result)
func (_Insurances *InsurancesSession) WatchCensorInsurance(fromBlock *uint64, handler func(int, []types.Log), _insuranceId *big.Int) error {
	return _Insurances.Contract.WatchCensorInsurance(fromBlock, handler, _insuranceId)
}

func (_Insurances *InsurancesSession) WatchAllCensorInsurance(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Insurances.Contract.WatchAllCensorInsurance(fromBlock, handler)
}

// ParseCensorInsurance is a log parse operation binding the contract event 0x6eb9ead4eb17be8377128ca9817d26b1caaee0d335e819f4303c2b6766468aa2.
//
// Solidity: event CensorInsurance(uint256 indexed _insuranceId, bool result)
func (_Insurances *InsurancesSession) ParseCensorInsurance(log types.Log) (*InsurancesCensorInsurance, error) {
	return _Insurances.Contract.ParseCensorInsurance(log)
}

// InsurancesExecuteInsurance represents a ExecuteInsurance event raised by the Insurances contract.
type InsurancesExecuteInsurance struct {
	InsuranceId *big.Int
	Value       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// WatchExecuteInsurance is a free log subscription operation binding the contract event 0x6f7e18f0d8960445e4ff55b876099f6d73f888232d5e5c91848247d3beadb50f.
//
// Solidity: event ExecuteInsurance(uint256 indexed _insuranceId, uint256 value)
func (_Insurances *InsurancesFilterer) WatchExecuteInsurance(fromBlock *uint64, handler func(int, []types.Log), _insuranceId *big.Int) error {
	return _Insurances.contract.WatchLogs(fromBlock, handler, "ExecuteInsurance", _insuranceId)
}

func (_Insurances *InsurancesFilterer) WatchAllExecuteInsurance(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Insurances.contract.WatchLogs(fromBlock, handler, "ExecuteInsurance")
}

// ParseExecuteInsurance is a log parse operation binding the contract event 0x6f7e18f0d8960445e4ff55b876099f6d73f888232d5e5c91848247d3beadb50f.
//
// Solidity: event ExecuteInsurance(uint256 indexed _insuranceId, uint256 value)
func (_Insurances *InsurancesFilterer) ParseExecuteInsurance(log types.Log) (*InsurancesExecuteInsurance, error) {
	event := new(InsurancesExecuteInsurance)
	if err := _Insurances.contract.UnpackLog(event, "ExecuteInsurance", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchExecuteInsurance is a free log subscription operation binding the contract event 0x6f7e18f0d8960445e4ff55b876099f6d73f888232d5e5c91848247d3beadb50f.
//
// Solidity: event ExecuteInsurance(uint256 indexed _insuranceId, uint256 value)
func (_Insurances *InsurancesSession) WatchExecuteInsurance(fromBlock *uint64, handler func(int, []types.Log), _insuranceId *big.Int) error {
	return _Insurances.Contract.WatchExecuteInsurance(fromBlock, handler, _insuranceId)
}

func (_Insurances *InsurancesSession) WatchAllExecuteInsurance(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Insurances.Contract.WatchAllExecuteInsurance(fromBlock, handler)
}

// ParseExecuteInsurance is a log parse operation binding the contract event 0x6f7e18f0d8960445e4ff55b876099f6d73f888232d5e5c91848247d3beadb50f.
//
// Solidity: event ExecuteInsurance(uint256 indexed _insuranceId, uint256 value)
func (_Insurances *InsurancesSession) ParseExecuteInsurance(log types.Log) (*InsurancesExecuteInsurance, error) {
	return _Insurances.Contract.ParseExecuteInsurance(log)
}

// InsurancesInsuranceExpired represents a InsuranceExpired event raised by the Insurances contract.
type InsurancesInsuranceExpired struct {
	InsuranceId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// WatchInsuranceExpired is a free log subscription operation binding the contract event 0xe18b4e336f36a5b2e805a39d09aaaf50084b53f5035295d2918900cf75155e69.
//
// Solidity: event InsuranceExpired(uint256 indexed _insuranceId)
func (_Insurances *InsurancesFilterer) WatchInsuranceExpired(fromBlock *uint64, handler func(int, []types.Log), _insuranceId *big.Int) error {
	return _Insurances.contract.WatchLogs(fromBlock, handler, "InsuranceExpired", _insuranceId)
}

func (_Insurances *InsurancesFilterer) WatchAllInsuranceExpired(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Insurances.contract.WatchLogs(fromBlock, handler, "InsuranceExpired")
}

// ParseInsuranceExpired is a log parse operation binding the contract event 0xe18b4e336f36a5b2e805a39d09aaaf50084b53f5035295d2918900cf75155e69.
//
// Solidity: event InsuranceExpired(uint256 indexed _insuranceId)
func (_Insurances *InsurancesFilterer) ParseInsuranceExpired(log types.Log) (*InsurancesInsuranceExpired, error) {
	event := new(InsurancesInsuranceExpired)
	if err := _Insurances.contract.UnpackLog(event, "InsuranceExpired", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchInsuranceExpired is a free log subscription operation binding the contract event 0xe18b4e336f36a5b2e805a39d09aaaf50084b53f5035295d2918900cf75155e69.
//
// Solidity: event InsuranceExpired(uint256 indexed _insuranceId)
func (_Insurances *InsurancesSession) WatchInsuranceExpired(fromBlock *uint64, handler func(int, []types.Log), _insuranceId *big.Int) error {
	return _Insurances.Contract.WatchInsuranceExpired(fromBlock, handler, _insuranceId)
}

func (_Insurances *InsurancesSession) WatchAllInsuranceExpired(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Insurances.Contract.WatchAllInsuranceExpired(fromBlock, handler)
}

// ParseInsuranceExpired is a log parse operation binding the contract event 0xe18b4e336f36a5b2e805a39d09aaaf50084b53f5035295d2918900cf75155e69.
//
// Solidity: event InsuranceExpired(uint256 indexed _insuranceId)
func (_Insurances *InsurancesSession) ParseInsuranceExpired(log types.Log) (*InsurancesInsuranceExpired, error) {
	return _Insurances.Contract.ParseInsuranceExpired(log)
}

// InsurancesInsuranceNotExpired represents a InsuranceNotExpired event raised by the Insurances contract.
type InsurancesInsuranceNotExpired struct {
	InsuranceId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// WatchInsuranceNotExpired is a free log subscription operation binding the contract event 0x8142336de75476de6d07bfe34897b339001e8a87a41b3e09ff071a9a8be7079b.
//
// Solidity: event InsuranceNotExpired(uint256 indexed _insuranceId)
func (_Insurances *InsurancesFilterer) WatchInsuranceNotExpired(fromBlock *uint64, handler func(int, []types.Log), _insuranceId *big.Int) error {
	return _Insurances.contract.WatchLogs(fromBlock, handler, "InsuranceNotExpired", _insuranceId)
}

func (_Insurances *InsurancesFilterer) WatchAllInsuranceNotExpired(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Insurances.contract.WatchLogs(fromBlock, handler, "InsuranceNotExpired")
}

// ParseInsuranceNotExpired is a log parse operation binding the contract event 0x8142336de75476de6d07bfe34897b339001e8a87a41b3e09ff071a9a8be7079b.
//
// Solidity: event InsuranceNotExpired(uint256 indexed _insuranceId)
func (_Insurances *InsurancesFilterer) ParseInsuranceNotExpired(log types.Log) (*InsurancesInsuranceNotExpired, error) {
	event := new(InsurancesInsuranceNotExpired)
	if err := _Insurances.contract.UnpackLog(event, "InsuranceNotExpired", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchInsuranceNotExpired is a free log subscription operation binding the contract event 0x8142336de75476de6d07bfe34897b339001e8a87a41b3e09ff071a9a8be7079b.
//
// Solidity: event InsuranceNotExpired(uint256 indexed _insuranceId)
func (_Insurances *InsurancesSession) WatchInsuranceNotExpired(fromBlock *uint64, handler func(int, []types.Log), _insuranceId *big.Int) error {
	return _Insurances.Contract.WatchInsuranceNotExpired(fromBlock, handler, _insuranceId)
}

func (_Insurances *InsurancesSession) WatchAllInsuranceNotExpired(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Insurances.Contract.WatchAllInsuranceNotExpired(fromBlock, handler)
}

// ParseInsuranceNotExpired is a log parse operation binding the contract event 0x8142336de75476de6d07bfe34897b339001e8a87a41b3e09ff071a9a8be7079b.
//
// Solidity: event InsuranceNotExpired(uint256 indexed _insuranceId)
func (_Insurances *InsurancesSession) ParseInsuranceNotExpired(log types.Log) (*InsurancesInsuranceNotExpired, error) {
	return _Insurances.Contract.ParseInsuranceNotExpired(log)
}

// InsurancesNewInsurance represents a NewInsurance event raised by the Insurances contract.
type InsurancesNewInsurance struct {
	InsuranceId *big.Int
	Beneficiary common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// WatchNewInsurance is a free log subscription operation binding the contract event 0x547962ec7fac5a8cc9580e652dc73617291177603313596daef886df1a27f8bc.
//
// Solidity: event NewInsurance(uint256 indexed _insuranceId, address indexed _beneficiary)
func (_Insurances *InsurancesFilterer) WatchNewInsurance(fromBlock *uint64, handler func(int, []types.Log), _insuranceId *big.Int, _beneficiary common.Address) error {
	return _Insurances.contract.WatchLogs(fromBlock, handler, "NewInsurance", _insuranceId, _beneficiary)
}

func (_Insurances *InsurancesFilterer) WatchAllNewInsurance(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Insurances.contract.WatchLogs(fromBlock, handler, "NewInsurance")
}

// ParseNewInsurance is a log parse operation binding the contract event 0x547962ec7fac5a8cc9580e652dc73617291177603313596daef886df1a27f8bc.
//
// Solidity: event NewInsurance(uint256 indexed _insuranceId, address indexed _beneficiary)
func (_Insurances *InsurancesFilterer) ParseNewInsurance(log types.Log) (*InsurancesNewInsurance, error) {
	event := new(InsurancesNewInsurance)
	if err := _Insurances.contract.UnpackLog(event, "NewInsurance", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchNewInsurance is a free log subscription operation binding the contract event 0x547962ec7fac5a8cc9580e652dc73617291177603313596daef886df1a27f8bc.
//
// Solidity: event NewInsurance(uint256 indexed _insuranceId, address indexed _beneficiary)
func (_Insurances *InsurancesSession) WatchNewInsurance(fromBlock *uint64, handler func(int, []types.Log), _insuranceId *big.Int, _beneficiary common.Address) error {
	return _Insurances.Contract.WatchNewInsurance(fromBlock, handler, _insuranceId, _beneficiary)
}

func (_Insurances *InsurancesSession) WatchAllNewInsurance(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Insurances.Contract.WatchAllNewInsurance(fromBlock, handler)
}

// ParseNewInsurance is a log parse operation binding the contract event 0x547962ec7fac5a8cc9580e652dc73617291177603313596daef886df1a27f8bc.
//
// Solidity: event NewInsurance(uint256 indexed _insuranceId, address indexed _beneficiary)
func (_Insurances *InsurancesSession) ParseNewInsurance(log types.Log) (*InsurancesNewInsurance, error) {
	return _Insurances.Contract.ParseNewInsurance(log)
}

// InsurancesWithdraw represents a Withdraw event raised by the Insurances contract.
type InsurancesWithdraw struct {
	AdminAddress common.Address
	Value        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed _adminAddress, uint256 value)
func (_Insurances *InsurancesFilterer) WatchWithdraw(fromBlock *uint64, handler func(int, []types.Log), _adminAddress common.Address) error {
	return _Insurances.contract.WatchLogs(fromBlock, handler, "Withdraw", _adminAddress)
}

func (_Insurances *InsurancesFilterer) WatchAllWithdraw(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Insurances.contract.WatchLogs(fromBlock, handler, "Withdraw")
}

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed _adminAddress, uint256 value)
func (_Insurances *InsurancesFilterer) ParseWithdraw(log types.Log) (*InsurancesWithdraw, error) {
	event := new(InsurancesWithdraw)
	if err := _Insurances.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed _adminAddress, uint256 value)
func (_Insurances *InsurancesSession) WatchWithdraw(fromBlock *uint64, handler func(int, []types.Log), _adminAddress common.Address) error {
	return _Insurances.Contract.WatchWithdraw(fromBlock, handler, _adminAddress)
}

func (_Insurances *InsurancesSession) WatchAllWithdraw(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Insurances.Contract.WatchAllWithdraw(fromBlock, handler)
}

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed _adminAddress, uint256 value)
func (_Insurances *InsurancesSession) ParseWithdraw(log types.Log) (*InsurancesWithdraw, error) {
	return _Insurances.Contract.ParseWithdraw(log)
}
