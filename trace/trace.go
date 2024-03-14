// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package trace

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// EcoTraceEcoInfo is an auto generated low-level Go binding around an user-defined struct.
type EcoTraceEcoInfo struct {
	Ecode     string
	Codata    string
	Operator  string
	Waterdata string
}

// TraceMetaData contains all meta data concerning the Trace contract.
var TraceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ecode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"codata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"operator\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"waterdata\",\"type\":\"string\"}],\"name\":\"addOrupdateProdData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"tracedata\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"ecode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"codata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"operator\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"waterdata\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"tracedataById\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"ecode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"codata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"operator\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"waterdata\",\"type\":\"string\"}],\"internalType\":\"structEcoTrace.EcoInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50335f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610ec88061005c5f395ff3fe608060405234801561000f575f80fd5b506004361061004a575f3560e01c80634e9fc3981461004e5780635c7e94951461007e5780638da5cb5b146100b1578063faeff189146100cf575b5f80fd5b61006860048036038101906100639190610818565b6100eb565b604051610075919061094e565b60405180910390f35b61009860048036038101906100939190610818565b610362565b6040516100a894939291906109b6565b60405180910390f35b6100b96105be565b6040516100c69190610a54565b60405180910390f35b6100e960048036038101906100e49190610a6d565b6105e1565b005b6100f36106a3565b6001826040516101039190610baa565b90815260200160405180910390206040518060800160405290815f8201805461012b90610bed565b80601f016020809104026020016040519081016040528092919081815260200182805461015790610bed565b80156101a25780601f10610179576101008083540402835291602001916101a2565b820191905f5260205f20905b81548152906001019060200180831161018557829003601f168201915b505050505081526020016001820180546101bb90610bed565b80601f01602080910402602001604051908101604052809291908181526020018280546101e790610bed565b80156102325780601f1061020957610100808354040283529160200191610232565b820191905f5260205f20905b81548152906001019060200180831161021557829003601f168201915b5050505050815260200160028201805461024b90610bed565b80601f016020809104026020016040519081016040528092919081815260200182805461027790610bed565b80156102c25780601f10610299576101008083540402835291602001916102c2565b820191905f5260205f20905b8154815290600101906020018083116102a557829003601f168201915b505050505081526020016003820180546102db90610bed565b80601f016020809104026020016040519081016040528092919081815260200182805461030790610bed565b80156103525780601f1061032957610100808354040283529160200191610352565b820191905f5260205f20905b81548152906001019060200180831161033557829003601f168201915b5050505050815250509050919050565b6001818051602081018201805184825260208301602085012081835280955050505050505f91509050805f01805461039990610bed565b80601f01602080910402602001604051908101604052809291908181526020018280546103c590610bed565b80156104105780601f106103e757610100808354040283529160200191610410565b820191905f5260205f20905b8154815290600101906020018083116103f357829003601f168201915b50505050509080600101805461042590610bed565b80601f016020809104026020016040519081016040528092919081815260200182805461045190610bed565b801561049c5780601f106104735761010080835404028352916020019161049c565b820191905f5260205f20905b81548152906001019060200180831161047f57829003601f168201915b5050505050908060020180546104b190610bed565b80601f01602080910402602001604051908101604052809291908181526020018280546104dd90610bed565b80156105285780601f106104ff57610100808354040283529160200191610528565b820191905f5260205f20905b81548152906001019060200180831161050b57829003601f168201915b50505050509080600301805461053d90610bed565b80601f016020809104026020016040519081016040528092919081815260200182805461056990610bed565b80156105b45780601f1061058b576101008083540402835291602001916105b4565b820191905f5260205f20905b81548152906001019060200180831161059757829003601f168201915b5050505050905084565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b836001866040516105f29190610baa565b90815260200160405180910390205f01908161060e9190610dc3565b50826001866040516106209190610baa565b9081526020016040518091039020600101908161063d9190610dc3565b508160018660405161064f9190610baa565b9081526020016040518091039020600201908161066c9190610dc3565b508060018660405161067e9190610baa565b9081526020016040518091039020600301908161069b9190610dc3565b505050505050565b6040518060800160405280606081526020016060815260200160608152602001606081525090565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61072a826106e4565b810181811067ffffffffffffffff82111715610749576107486106f4565b5b80604052505050565b5f61075b6106cb565b90506107678282610721565b919050565b5f67ffffffffffffffff821115610786576107856106f4565b5b61078f826106e4565b9050602081019050919050565b828183375f83830152505050565b5f6107bc6107b78461076c565b610752565b9050828152602081018484840111156107d8576107d76106e0565b5b6107e384828561079c565b509392505050565b5f82601f8301126107ff576107fe6106dc565b5b813561080f8482602086016107aa565b91505092915050565b5f6020828403121561082d5761082c6106d4565b5b5f82013567ffffffffffffffff81111561084a576108496106d8565b5b610856848285016107eb565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561089657808201518184015260208101905061087b565b5f8484015250505050565b5f6108ab8261085f565b6108b58185610869565b93506108c5818560208601610879565b6108ce816106e4565b840191505092915050565b5f608083015f8301518482035f8601526108f382826108a1565b9150506020830151848203602086015261090d82826108a1565b9150506040830151848203604086015261092782826108a1565b9150506060830151848203606086015261094182826108a1565b9150508091505092915050565b5f6020820190508181035f83015261096681846108d9565b905092915050565b5f82825260208201905092915050565b5f6109888261085f565b610992818561096e565b93506109a2818560208601610879565b6109ab816106e4565b840191505092915050565b5f6080820190508181035f8301526109ce818761097e565b905081810360208301526109e2818661097e565b905081810360408301526109f6818561097e565b90508181036060830152610a0a818461097e565b905095945050505050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610a3e82610a15565b9050919050565b610a4e81610a34565b82525050565b5f602082019050610a675f830184610a45565b92915050565b5f805f805f60a08688031215610a8657610a856106d4565b5b5f86013567ffffffffffffffff811115610aa357610aa26106d8565b5b610aaf888289016107eb565b955050602086013567ffffffffffffffff811115610ad057610acf6106d8565b5b610adc888289016107eb565b945050604086013567ffffffffffffffff811115610afd57610afc6106d8565b5b610b09888289016107eb565b935050606086013567ffffffffffffffff811115610b2a57610b296106d8565b5b610b36888289016107eb565b925050608086013567ffffffffffffffff811115610b5757610b566106d8565b5b610b63888289016107eb565b9150509295509295909350565b5f81905092915050565b5f610b848261085f565b610b8e8185610b70565b9350610b9e818560208601610879565b80840191505092915050565b5f610bb58284610b7a565b915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610c0457607f821691505b602082108103610c1757610c16610bc0565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610c797fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610c3e565b610c838683610c3e565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f610cc7610cc2610cbd84610c9b565b610ca4565b610c9b565b9050919050565b5f819050919050565b610ce083610cad565b610cf4610cec82610cce565b848454610c4a565b825550505050565b5f90565b610d08610cfc565b610d13818484610cd7565b505050565b5b81811015610d3657610d2b5f82610d00565b600181019050610d19565b5050565b601f821115610d7b57610d4c81610c1d565b610d5584610c2f565b81016020851015610d64578190505b610d78610d7085610c2f565b830182610d18565b50505b505050565b5f82821c905092915050565b5f610d9b5f1984600802610d80565b1980831691505092915050565b5f610db38383610d8c565b9150826002028217905092915050565b610dcc8261085f565b67ffffffffffffffff811115610de557610de46106f4565b5b610def8254610bed565b610dfa828285610d3a565b5f60209050601f831160018114610e2b575f8415610e19578287015190505b610e238582610da8565b865550610e8a565b601f198416610e3986610c1d565b5f5b82811015610e6057848901518255600182019150602085019450602081019050610e3b565b86831015610e7d5784890151610e79601f891682610d8c565b8355505b6001600288020188555050505b50505050505056fea2646970667358221220e6a1c5a566b6cd9477296254741776c67b02a0821acfce8aa92c49e069a52d6e64736f6c63430008180033",
}

// TraceABI is the input ABI used to generate the binding from.
// Deprecated: Use TraceMetaData.ABI instead.
var TraceABI = TraceMetaData.ABI

// TraceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TraceMetaData.Bin instead.
var TraceBin = TraceMetaData.Bin

// DeployTrace deploys a new Ethereum contract, binding an instance of Trace to it.
func DeployTrace(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Trace, error) {
	parsed, err := TraceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TraceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Trace{TraceCaller: TraceCaller{contract: contract}, TraceTransactor: TraceTransactor{contract: contract}, TraceFilterer: TraceFilterer{contract: contract}}, nil
}

// Trace is an auto generated Go binding around an Ethereum contract.
type Trace struct {
	TraceCaller     // Read-only binding to the contract
	TraceTransactor // Write-only binding to the contract
	TraceFilterer   // Log filterer for contract events
}

// TraceCaller is an auto generated read-only Go binding around an Ethereum contract.
type TraceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TraceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TraceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TraceSession struct {
	Contract     *Trace            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TraceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TraceCallerSession struct {
	Contract *TraceCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TraceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TraceTransactorSession struct {
	Contract     *TraceTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TraceRaw is an auto generated low-level Go binding around an Ethereum contract.
type TraceRaw struct {
	Contract *Trace // Generic contract binding to access the raw methods on
}

// TraceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TraceCallerRaw struct {
	Contract *TraceCaller // Generic read-only contract binding to access the raw methods on
}

// TraceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TraceTransactorRaw struct {
	Contract *TraceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrace creates a new instance of Trace, bound to a specific deployed contract.
func NewTrace(address common.Address, backend bind.ContractBackend) (*Trace, error) {
	contract, err := bindTrace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Trace{TraceCaller: TraceCaller{contract: contract}, TraceTransactor: TraceTransactor{contract: contract}, TraceFilterer: TraceFilterer{contract: contract}}, nil
}

// NewTraceCaller creates a new read-only instance of Trace, bound to a specific deployed contract.
func NewTraceCaller(address common.Address, caller bind.ContractCaller) (*TraceCaller, error) {
	contract, err := bindTrace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TraceCaller{contract: contract}, nil
}

// NewTraceTransactor creates a new write-only instance of Trace, bound to a specific deployed contract.
func NewTraceTransactor(address common.Address, transactor bind.ContractTransactor) (*TraceTransactor, error) {
	contract, err := bindTrace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TraceTransactor{contract: contract}, nil
}

// NewTraceFilterer creates a new log filterer instance of Trace, bound to a specific deployed contract.
func NewTraceFilterer(address common.Address, filterer bind.ContractFilterer) (*TraceFilterer, error) {
	contract, err := bindTrace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TraceFilterer{contract: contract}, nil
}

// bindTrace binds a generic wrapper to an already deployed contract.
func bindTrace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TraceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trace *TraceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trace.Contract.TraceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trace *TraceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trace.Contract.TraceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trace *TraceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trace.Contract.TraceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trace *TraceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trace *TraceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trace *TraceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trace.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Trace *TraceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Trace.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Trace *TraceSession) Owner() (common.Address, error) {
	return _Trace.Contract.Owner(&_Trace.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Trace *TraceCallerSession) Owner() (common.Address, error) {
	return _Trace.Contract.Owner(&_Trace.CallOpts)
}

// Tracedata is a free data retrieval call binding the contract method 0x5c7e9495.
//
// Solidity: function tracedata(string ) view returns(string ecode, string codata, string operator, string waterdata)
func (_Trace *TraceCaller) Tracedata(opts *bind.CallOpts, arg0 string) (struct {
	Ecode     string
	Codata    string
	Operator  string
	Waterdata string
}, error) {
	var out []interface{}
	err := _Trace.contract.Call(opts, &out, "tracedata", arg0)

	outstruct := new(struct {
		Ecode     string
		Codata    string
		Operator  string
		Waterdata string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ecode = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Codata = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Operator = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Waterdata = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// Tracedata is a free data retrieval call binding the contract method 0x5c7e9495.
//
// Solidity: function tracedata(string ) view returns(string ecode, string codata, string operator, string waterdata)
func (_Trace *TraceSession) Tracedata(arg0 string) (struct {
	Ecode     string
	Codata    string
	Operator  string
	Waterdata string
}, error) {
	return _Trace.Contract.Tracedata(&_Trace.CallOpts, arg0)
}

// Tracedata is a free data retrieval call binding the contract method 0x5c7e9495.
//
// Solidity: function tracedata(string ) view returns(string ecode, string codata, string operator, string waterdata)
func (_Trace *TraceCallerSession) Tracedata(arg0 string) (struct {
	Ecode     string
	Codata    string
	Operator  string
	Waterdata string
}, error) {
	return _Trace.Contract.Tracedata(&_Trace.CallOpts, arg0)
}

// TracedataById is a free data retrieval call binding the contract method 0x4e9fc398.
//
// Solidity: function tracedataById(string id) view returns((string,string,string,string))
func (_Trace *TraceCaller) TracedataById(opts *bind.CallOpts, id string) (EcoTraceEcoInfo, error) {
	var out []interface{}
	err := _Trace.contract.Call(opts, &out, "tracedataById", id)

	if err != nil {
		return *new(EcoTraceEcoInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(EcoTraceEcoInfo)).(*EcoTraceEcoInfo)

	return out0, err

}

// TracedataById is a free data retrieval call binding the contract method 0x4e9fc398.
//
// Solidity: function tracedataById(string id) view returns((string,string,string,string))
func (_Trace *TraceSession) TracedataById(id string) (EcoTraceEcoInfo, error) {
	return _Trace.Contract.TracedataById(&_Trace.CallOpts, id)
}

// TracedataById is a free data retrieval call binding the contract method 0x4e9fc398.
//
// Solidity: function tracedataById(string id) view returns((string,string,string,string))
func (_Trace *TraceCallerSession) TracedataById(id string) (EcoTraceEcoInfo, error) {
	return _Trace.Contract.TracedataById(&_Trace.CallOpts, id)
}

// AddOrupdateProdData is a paid mutator transaction binding the contract method 0xfaeff189.
//
// Solidity: function addOrupdateProdData(string id, string ecode, string codata, string operator, string waterdata) returns()
func (_Trace *TraceTransactor) AddOrupdateProdData(opts *bind.TransactOpts, id string, ecode string, codata string, operator string, waterdata string) (*types.Transaction, error) {
	return _Trace.contract.Transact(opts, "addOrupdateProdData", id, ecode, codata, operator, waterdata)
}

// AddOrupdateProdData is a paid mutator transaction binding the contract method 0xfaeff189.
//
// Solidity: function addOrupdateProdData(string id, string ecode, string codata, string operator, string waterdata) returns()
func (_Trace *TraceSession) AddOrupdateProdData(id string, ecode string, codata string, operator string, waterdata string) (*types.Transaction, error) {
	return _Trace.Contract.AddOrupdateProdData(&_Trace.TransactOpts, id, ecode, codata, operator, waterdata)
}

// AddOrupdateProdData is a paid mutator transaction binding the contract method 0xfaeff189.
//
// Solidity: function addOrupdateProdData(string id, string ecode, string codata, string operator, string waterdata) returns()
func (_Trace *TraceTransactorSession) AddOrupdateProdData(id string, ecode string, codata string, operator string, waterdata string) (*types.Transaction, error) {
	return _Trace.Contract.AddOrupdateProdData(&_Trace.TransactOpts, id, ecode, codata, operator, waterdata)
}
