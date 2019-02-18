// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package constant

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ConstantABI is the input ABI used to generate the binding from.
const ConstantABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_monetary\",\"type\":\"address\"},{\"name\":\"_oracle\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"__transferByAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"__purchase\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"__redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"lid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"__borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"__payoff\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"__liquidate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"transferByAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"purchaser\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"purchase\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"redeemer\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"redeem\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"borrow\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"lid\",\"type\":\"uint256\"},{\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"payoff\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"lid\",\"type\":\"uint256\"},{\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"liquidate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ConstantBin is the compiled bytecode used for deploying new contracts.
const ConstantBin = `0x608060405234801561001057600080fd5b50604051604080612295833981018060405281019080805190602001909291908051906020019092919050505033600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050612184806101116000396000f3006080604052600436106100f1576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306fdde03146100f6578063095ea7b31461018657806318160ddd146101eb57806323b872dd14610216578063313ce5671461029b57806339509351146102c65780634c5552821461032b57806370a082311461036657806395d89b41146103bd578063992c3e4b1461044d578063a3749f28146104a8578063a457c2d714610503578063a9059cbb14610568578063d5eed868146105cd578063dd62ed3e146105f1578063e137f31b14610668578063f9ba884f146106a3575b600080fd5b34801561010257600080fd5b5061010b61071e565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561014b578082015181840152602081019050610130565b50505050905090810190601f1680156101785780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561019257600080fd5b506101d1600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610757565b604051808215151515815260200191505060405180910390f35b3480156101f757600080fd5b50610200610884565b6040518082815260200191505060405180910390f35b34801561022257600080fd5b50610281600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061088e565b604051808215151515815260200191505060405180910390f35b3480156102a757600080fd5b506102b0610a40565b6040518082815260200191505060405180910390f35b3480156102d257600080fd5b50610311600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610a45565b604051808215151515815260200191505060405180910390f35b34801561033757600080fd5b50610364600480360381019080803590602001909291908035600019169060200190929190505050610c7c565b005b34801561037257600080fd5b506103a7600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f12565b6040518082815260200191505060405180910390f35b3480156103c957600080fd5b506103d2610f5a565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156104125780820151818401526020810190506103f7565b50505050905090810190601f16801561043f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561045957600080fd5b506104a6600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035600019169060200190929190505050610f93565b005b3480156104b457600080fd5b50610501600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803560001916906020019092919050505061103d565b005b34801561050f57600080fd5b5061054e600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506110e7565b604051808215151515815260200191505060405180910390f35b34801561057457600080fd5b506105b3600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061131e565b604051808215151515815260200191505060405180910390f35b6105ef6004803603810190808035600019169060200190929190505050611335565b005b3480156105fd57600080fd5b50610652600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506118e5565b6040518082815260200191505060405180910390f35b34801561067457600080fd5b506106a160048036038101908080359060200190929190803560001916906020019092919050505061196c565b005b3480156106af57600080fd5b5061071c600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035600019169060200190929190505050611b07565b005b6040805190810160405280601381526020017f436f6e7374616e7420537461626c65636f696e0000000000000000000000000081525081565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415151561079457600080fd5b81600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040518082815260200191505060405180910390a36001905092915050565b6000600254905090565b6000600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054821115151561091b57600080fd5b6109aa82600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611bb390919063ffffffff16565b600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610a35848484611bd4565b600190509392505050565b600281565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614151515610a8257600080fd5b610b1182600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611ded90919063ffffffff16565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546040518082815260200191505060405180910390a36001905092915050565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610cda57600080fd5b600483815481101515610ce957fe5b9060005260206000209060070201905060006002811115610d0657fe5b8160060160149054906101000a900460ff166002811115610d2357fe5b141515610d2f57600080fd5b8060030160020154816003016001015402600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637bfb09346040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180807f65746850726963650000000000000000000000000000000000000000000000008152506020019050602060405180830381600087803b158015610df157600080fd5b505af1158015610e05573d6000803e3d6000fd5b505050506040513d6020811015610e1b57600080fd5b8101908080519060200190929190505050101515610e3857600080fd5b60028160060160146101000a81548160ff02191690836002811115610e5957fe5b0217905550600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc82600301600001549081150290604051600060405180830381858888f19350505050158015610ecd573d6000803e3d6000fd5b507f19ab7221fec3242f7442ff80888f33c250afdf8d4cf62583cec2748acaad0ec48260405180826000191660001916815260200191505060405180910390a1505050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6040805190810160405280600581526020017f434f4e535400000000000000000000000000000000000000000000000000000081525081565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610fef57600080fd5b610ff98383611e0e565b7fb0de879351469d2741406aafc9ba1f44eb957cf44ee3391e59a7a9097050c9278160405180826000191660001916815260200191505060405180910390a1505050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561109957600080fd5b6110a38383611f99565b7fef72b9890ab0fc46404e72534dc1bbc275de9d7efd8b3657ad91f82e1a3d39c48160405180826000191660001916815260200191505060405180910390a1505050565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415151561112457600080fd5b6111b382600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611bb390919063ffffffff16565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546040518082815260200191505060405180910390a36001905092915050565b600061132b338484611bd4565b6001905092915050565b61133d6120d7565b6000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637bfb09346040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180807f65746850726963650000000000000000000000000000000000000000000000008152506020019050602060405180830381600087803b1580156113f057600080fd5b505af1158015611404573d6000803e3d6000fd5b505050506040513d602081101561141a57600080fd5b810190808051906020019092919050505090506064600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632adfef746040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180807f657468526174696f0000000000000000000000000000000000000000000000008152506020019050602060405180830381600087803b1580156114e057600080fd5b505af11580156114f4573d6000803e3d6000fd5b505050506040513d602081101561150a57600080fd5b81019080805190602001909291905050508234020281151561152857fe5b04826000018181525050600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632adfef746040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180807f657468496e7465726573740000000000000000000000000000000000000000008152506020019050602060405180830381600087803b1580156115e357600080fd5b505af11580156115f7573d6000803e3d6000fd5b505050506040513d602081101561160d57600080fd5b8101908080519060200190929190505050826020018181525050606060405190810160405280348152602001828152602001600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632adfef746040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180807f6574685468726573686f6c6400000000000000000000000000000000000000008152506020019050602060405180830381600087803b1580156116f057600080fd5b505af1158015611704573d6000803e3d6000fd5b505050506040513d602081101561171a57600080fd5b810190808051906020019092919050505081525082606001819052504282604001818152505033826080019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505060008260a00190600281111561178a57fe5b9081600281111561179757fe5b8152505060048290806001815401808255809150509060018203906000526020600020906007020160009091929091909150600082015181600001556020820151816001015560408201518160020155606082015181600301600082015181600001556020820151816001015560408201518160020155505060808201518160060160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a08201518160060160146101000a81548160ff0219169083600281111561187b57fe5b0217905550505050611891338360000151611f99565b7f036872352a35530308027377cd0aef0e9489e4339550fb7fcc3ae38bc648a134600160048054905003846040518083815260200182600019166000191681526020019250505060405180910390a1505050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600060048381548110151561197d57fe5b906000526020600020906007020190506000600281111561199a57fe5b8160060160149054906101000a900460ff1660028111156119b757fe5b1415156119c357600080fd5b8060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611a2157600080fd5b60018160060160146101000a81548160ff02191690836002811115611a4257fe5b0217905550611a67333083600201544203846001015402600101846000015402611bd4565b611a75308260000154611e0e565b3373ffffffffffffffffffffffffffffffffffffffff166108fc82600301600001549081150290604051600060405180830381858888f19350505050158015611ac2573d6000803e3d6000fd5b507fd68d5698fae2dc1d31bf2bf1ea674da9ebcc3989a5c8b6892414890a1de99a808260405180826000191660001916815260200191505060405180910390a1505050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611b6357600080fd5b611b6e848484611bd4565b7f3b36ee6b35325f38e95938557be92853c842b7a9a19fd7ac4931a6d24db526828160405180826000191660001916815260200191505060405180910390a150505050565b600080838311151515611bc557600080fd5b82840390508091505092915050565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548111151515611c2157600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614151515611c5d57600080fd5b611cae816000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611bb390919063ffffffff16565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550611d41816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611ded90919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a3505050565b6000808284019050838110151515611e0457600080fd5b8091505092915050565b60008273ffffffffffffffffffffffffffffffffffffffff1614151515611e3457600080fd5b6000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548111151515611e8157600080fd5b611e9681600254611bb390919063ffffffff16565b600281905550611eed816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611bb390919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b60008273ffffffffffffffffffffffffffffffffffffffff1614151515611fbf57600080fd5b611fd481600254611ded90919063ffffffff16565b60028190555061202b816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611ded90919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b61010060405190810160405280600081526020016000815260200160008152602001612101612136565b8152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016000600281111561213057fe5b81525090565b60606040519081016040528060008152602001600081526020016000815250905600a165627a7a723058200bb491f66438637ba2d597eaf3aa7714b5a14150f869d14484173ec3031653260029`

// DeployConstant deploys a new Ethereum contract, binding an instance of Constant to it.
func DeployConstant(auth *bind.TransactOpts, backend bind.ContractBackend, _monetary common.Address, _oracle common.Address) (common.Address, *types.Transaction, *Constant, error) {
	parsed, err := abi.JSON(strings.NewReader(ConstantABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ConstantBin), backend, _monetary, _oracle)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Constant{ConstantCaller: ConstantCaller{contract: contract}, ConstantTransactor: ConstantTransactor{contract: contract}, ConstantFilterer: ConstantFilterer{contract: contract}}, nil
}

// Constant is an auto generated Go binding around an Ethereum contract.
type Constant struct {
	ConstantCaller     // Read-only binding to the contract
	ConstantTransactor // Write-only binding to the contract
	ConstantFilterer   // Log filterer for contract events
}

// ConstantCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConstantCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstantTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConstantTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstantFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConstantFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstantSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConstantSession struct {
	Contract     *Constant         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConstantCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConstantCallerSession struct {
	Contract *ConstantCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ConstantTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConstantTransactorSession struct {
	Contract     *ConstantTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ConstantRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConstantRaw struct {
	Contract *Constant // Generic contract binding to access the raw methods on
}

// ConstantCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConstantCallerRaw struct {
	Contract *ConstantCaller // Generic read-only contract binding to access the raw methods on
}

// ConstantTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConstantTransactorRaw struct {
	Contract *ConstantTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConstant creates a new instance of Constant, bound to a specific deployed contract.
func NewConstant(address common.Address, backend bind.ContractBackend) (*Constant, error) {
	contract, err := bindConstant(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Constant{ConstantCaller: ConstantCaller{contract: contract}, ConstantTransactor: ConstantTransactor{contract: contract}, ConstantFilterer: ConstantFilterer{contract: contract}}, nil
}

// NewConstantCaller creates a new read-only instance of Constant, bound to a specific deployed contract.
func NewConstantCaller(address common.Address, caller bind.ContractCaller) (*ConstantCaller, error) {
	contract, err := bindConstant(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConstantCaller{contract: contract}, nil
}

// NewConstantTransactor creates a new write-only instance of Constant, bound to a specific deployed contract.
func NewConstantTransactor(address common.Address, transactor bind.ContractTransactor) (*ConstantTransactor, error) {
	contract, err := bindConstant(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConstantTransactor{contract: contract}, nil
}

// NewConstantFilterer creates a new log filterer instance of Constant, bound to a specific deployed contract.
func NewConstantFilterer(address common.Address, filterer bind.ContractFilterer) (*ConstantFilterer, error) {
	contract, err := bindConstant(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConstantFilterer{contract: contract}, nil
}

// bindConstant binds a generic wrapper to an already deployed contract.
func bindConstant(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConstantABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Constant *ConstantRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Constant.Contract.ConstantCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Constant *ConstantRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Constant.Contract.ConstantTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Constant *ConstantRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Constant.Contract.ConstantTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Constant *ConstantCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Constant.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Constant *ConstantTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Constant.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Constant *ConstantTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Constant.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_Constant *ConstantCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Constant.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_Constant *ConstantSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Constant.Contract.Allowance(&_Constant.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_Constant *ConstantCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Constant.Contract.Allowance(&_Constant.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_Constant *ConstantCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Constant.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_Constant *ConstantSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Constant.Contract.BalanceOf(&_Constant.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_Constant *ConstantCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Constant.Contract.BalanceOf(&_Constant.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_Constant *ConstantCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Constant.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_Constant *ConstantSession) Decimals() (*big.Int, error) {
	return _Constant.Contract.Decimals(&_Constant.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_Constant *ConstantCallerSession) Decimals() (*big.Int, error) {
	return _Constant.Contract.Decimals(&_Constant.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Constant *ConstantCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Constant.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Constant *ConstantSession) Name() (string, error) {
	return _Constant.Contract.Name(&_Constant.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Constant *ConstantCallerSession) Name() (string, error) {
	return _Constant.Contract.Name(&_Constant.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Constant *ConstantCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Constant.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Constant *ConstantSession) Symbol() (string, error) {
	return _Constant.Contract.Symbol(&_Constant.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Constant *ConstantCallerSession) Symbol() (string, error) {
	return _Constant.Contract.Symbol(&_Constant.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Constant *ConstantCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Constant.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Constant *ConstantSession) TotalSupply() (*big.Int, error) {
	return _Constant.Contract.TotalSupply(&_Constant.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Constant *ConstantCallerSession) TotalSupply() (*big.Int, error) {
	return _Constant.Contract.TotalSupply(&_Constant.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Constant *ConstantTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Constant.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Constant *ConstantSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Constant.Contract.Approve(&_Constant.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Constant *ConstantTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Constant.Contract.Approve(&_Constant.TransactOpts, spender, value)
}

// Borrow is a paid mutator transaction binding the contract method 0xd5eed868.
//
// Solidity: function borrow(bytes32 offchain) returns()
func (_Constant *ConstantTransactor) Borrow(opts *bind.TransactOpts, offchain [32]byte) (*types.Transaction, error) {
	return _Constant.contract.Transact(opts, "borrow", offchain)
}

// Borrow is a paid mutator transaction binding the contract method 0xd5eed868.
//
// Solidity: function borrow(bytes32 offchain) returns()
func (_Constant *ConstantSession) Borrow(offchain [32]byte) (*types.Transaction, error) {
	return _Constant.Contract.Borrow(&_Constant.TransactOpts, offchain)
}

// Borrow is a paid mutator transaction binding the contract method 0xd5eed868.
//
// Solidity: function borrow(bytes32 offchain) returns()
func (_Constant *ConstantTransactorSession) Borrow(offchain [32]byte) (*types.Transaction, error) {
	return _Constant.Contract.Borrow(&_Constant.TransactOpts, offchain)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Constant *ConstantTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Constant.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Constant *ConstantSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Constant.Contract.DecreaseAllowance(&_Constant.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Constant *ConstantTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Constant.Contract.DecreaseAllowance(&_Constant.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Constant *ConstantTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Constant.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Constant *ConstantSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Constant.Contract.IncreaseAllowance(&_Constant.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Constant *ConstantTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Constant.Contract.IncreaseAllowance(&_Constant.TransactOpts, spender, addedValue)
}

// Liquidate is a paid mutator transaction binding the contract method 0x4c555282.
//
// Solidity: function liquidate(uint256 lid, bytes32 offchain) returns()
func (_Constant *ConstantTransactor) Liquidate(opts *bind.TransactOpts, lid *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _Constant.contract.Transact(opts, "liquidate", lid, offchain)
}

// Liquidate is a paid mutator transaction binding the contract method 0x4c555282.
//
// Solidity: function liquidate(uint256 lid, bytes32 offchain) returns()
func (_Constant *ConstantSession) Liquidate(lid *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _Constant.Contract.Liquidate(&_Constant.TransactOpts, lid, offchain)
}

// Liquidate is a paid mutator transaction binding the contract method 0x4c555282.
//
// Solidity: function liquidate(uint256 lid, bytes32 offchain) returns()
func (_Constant *ConstantTransactorSession) Liquidate(lid *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _Constant.Contract.Liquidate(&_Constant.TransactOpts, lid, offchain)
}

// Payoff is a paid mutator transaction binding the contract method 0xe137f31b.
//
// Solidity: function payoff(uint256 lid, bytes32 offchain) returns()
func (_Constant *ConstantTransactor) Payoff(opts *bind.TransactOpts, lid *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _Constant.contract.Transact(opts, "payoff", lid, offchain)
}

// Payoff is a paid mutator transaction binding the contract method 0xe137f31b.
//
// Solidity: function payoff(uint256 lid, bytes32 offchain) returns()
func (_Constant *ConstantSession) Payoff(lid *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _Constant.Contract.Payoff(&_Constant.TransactOpts, lid, offchain)
}

// Payoff is a paid mutator transaction binding the contract method 0xe137f31b.
//
// Solidity: function payoff(uint256 lid, bytes32 offchain) returns()
func (_Constant *ConstantTransactorSession) Payoff(lid *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _Constant.Contract.Payoff(&_Constant.TransactOpts, lid, offchain)
}

// Purchase is a paid mutator transaction binding the contract method 0xa3749f28.
//
// Solidity: function purchase(address purchaser, uint256 value, bytes32 offchain) returns()
func (_Constant *ConstantTransactor) Purchase(opts *bind.TransactOpts, purchaser common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _Constant.contract.Transact(opts, "purchase", purchaser, value, offchain)
}

// Purchase is a paid mutator transaction binding the contract method 0xa3749f28.
//
// Solidity: function purchase(address purchaser, uint256 value, bytes32 offchain) returns()
func (_Constant *ConstantSession) Purchase(purchaser common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _Constant.Contract.Purchase(&_Constant.TransactOpts, purchaser, value, offchain)
}

// Purchase is a paid mutator transaction binding the contract method 0xa3749f28.
//
// Solidity: function purchase(address purchaser, uint256 value, bytes32 offchain) returns()
func (_Constant *ConstantTransactorSession) Purchase(purchaser common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _Constant.Contract.Purchase(&_Constant.TransactOpts, purchaser, value, offchain)
}

// Redeem is a paid mutator transaction binding the contract method 0x992c3e4b.
//
// Solidity: function redeem(address redeemer, uint256 value, bytes32 offchain) returns()
func (_Constant *ConstantTransactor) Redeem(opts *bind.TransactOpts, redeemer common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _Constant.contract.Transact(opts, "redeem", redeemer, value, offchain)
}

// Redeem is a paid mutator transaction binding the contract method 0x992c3e4b.
//
// Solidity: function redeem(address redeemer, uint256 value, bytes32 offchain) returns()
func (_Constant *ConstantSession) Redeem(redeemer common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _Constant.Contract.Redeem(&_Constant.TransactOpts, redeemer, value, offchain)
}

// Redeem is a paid mutator transaction binding the contract method 0x992c3e4b.
//
// Solidity: function redeem(address redeemer, uint256 value, bytes32 offchain) returns()
func (_Constant *ConstantTransactorSession) Redeem(redeemer common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _Constant.Contract.Redeem(&_Constant.TransactOpts, redeemer, value, offchain)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Constant *ConstantTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Constant.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Constant *ConstantSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Constant.Contract.Transfer(&_Constant.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Constant *ConstantTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Constant.Contract.Transfer(&_Constant.TransactOpts, to, value)
}

// TransferByAdmin is a paid mutator transaction binding the contract method 0xf9ba884f.
//
// Solidity: function transferByAdmin(address from, address to, uint256 value, bytes32 offchain) returns()
func (_Constant *ConstantTransactor) TransferByAdmin(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _Constant.contract.Transact(opts, "transferByAdmin", from, to, value, offchain)
}

// TransferByAdmin is a paid mutator transaction binding the contract method 0xf9ba884f.
//
// Solidity: function transferByAdmin(address from, address to, uint256 value, bytes32 offchain) returns()
func (_Constant *ConstantSession) TransferByAdmin(from common.Address, to common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _Constant.Contract.TransferByAdmin(&_Constant.TransactOpts, from, to, value, offchain)
}

// TransferByAdmin is a paid mutator transaction binding the contract method 0xf9ba884f.
//
// Solidity: function transferByAdmin(address from, address to, uint256 value, bytes32 offchain) returns()
func (_Constant *ConstantTransactorSession) TransferByAdmin(from common.Address, to common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _Constant.Contract.TransferByAdmin(&_Constant.TransactOpts, from, to, value, offchain)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Constant *ConstantTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Constant.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Constant *ConstantSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Constant.Contract.TransferFrom(&_Constant.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Constant *ConstantTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Constant.Contract.TransferFrom(&_Constant.TransactOpts, from, to, value)
}

// ConstantApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Constant contract.
type ConstantApprovalIterator struct {
	Event *ConstantApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConstantApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConstantApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConstantApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConstantApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConstantApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConstantApproval represents a Approval event raised by the Constant contract.
type ConstantApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Constant *ConstantFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ConstantApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Constant.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ConstantApprovalIterator{contract: _Constant.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Constant *ConstantFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ConstantApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Constant.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConstantApproval)
				if err := _Constant.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ConstantTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Constant contract.
type ConstantTransferIterator struct {
	Event *ConstantTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConstantTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConstantTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConstantTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConstantTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConstantTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConstantTransfer represents a Transfer event raised by the Constant contract.
type ConstantTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Constant *ConstantFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ConstantTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Constant.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ConstantTransferIterator{contract: _Constant.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Constant *ConstantFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ConstantTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Constant.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConstantTransfer)
				if err := _Constant.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ConstantBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the Constant contract.
type ConstantBorrowIterator struct {
	Event *ConstantBorrow // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConstantBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConstantBorrow)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConstantBorrow)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConstantBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConstantBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConstantBorrow represents a Borrow event raised by the Constant contract.
type ConstantBorrow struct {
	Lid      *big.Int
	Offchain [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x036872352a35530308027377cd0aef0e9489e4339550fb7fcc3ae38bc648a134.
//
// Solidity: event __borrow(uint256 lid, bytes32 offchain)
func (_Constant *ConstantFilterer) FilterBorrow(opts *bind.FilterOpts) (*ConstantBorrowIterator, error) {

	logs, sub, err := _Constant.contract.FilterLogs(opts, "__borrow")
	if err != nil {
		return nil, err
	}
	return &ConstantBorrowIterator{contract: _Constant.contract, event: "__borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x036872352a35530308027377cd0aef0e9489e4339550fb7fcc3ae38bc648a134.
//
// Solidity: event __borrow(uint256 lid, bytes32 offchain)
func (_Constant *ConstantFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *ConstantBorrow) (event.Subscription, error) {

	logs, sub, err := _Constant.contract.WatchLogs(opts, "__borrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConstantBorrow)
				if err := _Constant.contract.UnpackLog(event, "__borrow", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ConstantLiquidateIterator is returned from FilterLiquidate and is used to iterate over the raw logs and unpacked data for Liquidate events raised by the Constant contract.
type ConstantLiquidateIterator struct {
	Event *ConstantLiquidate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConstantLiquidateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConstantLiquidate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConstantLiquidate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConstantLiquidateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConstantLiquidateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConstantLiquidate represents a Liquidate event raised by the Constant contract.
type ConstantLiquidate struct {
	Offchain [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLiquidate is a free log retrieval operation binding the contract event 0x19ab7221fec3242f7442ff80888f33c250afdf8d4cf62583cec2748acaad0ec4.
//
// Solidity: event __liquidate(bytes32 offchain)
func (_Constant *ConstantFilterer) FilterLiquidate(opts *bind.FilterOpts) (*ConstantLiquidateIterator, error) {

	logs, sub, err := _Constant.contract.FilterLogs(opts, "__liquidate")
	if err != nil {
		return nil, err
	}
	return &ConstantLiquidateIterator{contract: _Constant.contract, event: "__liquidate", logs: logs, sub: sub}, nil
}

// WatchLiquidate is a free log subscription operation binding the contract event 0x19ab7221fec3242f7442ff80888f33c250afdf8d4cf62583cec2748acaad0ec4.
//
// Solidity: event __liquidate(bytes32 offchain)
func (_Constant *ConstantFilterer) WatchLiquidate(opts *bind.WatchOpts, sink chan<- *ConstantLiquidate) (event.Subscription, error) {

	logs, sub, err := _Constant.contract.WatchLogs(opts, "__liquidate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConstantLiquidate)
				if err := _Constant.contract.UnpackLog(event, "__liquidate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ConstantPayoffIterator is returned from FilterPayoff and is used to iterate over the raw logs and unpacked data for Payoff events raised by the Constant contract.
type ConstantPayoffIterator struct {
	Event *ConstantPayoff // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConstantPayoffIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConstantPayoff)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConstantPayoff)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConstantPayoffIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConstantPayoffIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConstantPayoff represents a Payoff event raised by the Constant contract.
type ConstantPayoff struct {
	Offchain [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPayoff is a free log retrieval operation binding the contract event 0xd68d5698fae2dc1d31bf2bf1ea674da9ebcc3989a5c8b6892414890a1de99a80.
//
// Solidity: event __payoff(bytes32 offchain)
func (_Constant *ConstantFilterer) FilterPayoff(opts *bind.FilterOpts) (*ConstantPayoffIterator, error) {

	logs, sub, err := _Constant.contract.FilterLogs(opts, "__payoff")
	if err != nil {
		return nil, err
	}
	return &ConstantPayoffIterator{contract: _Constant.contract, event: "__payoff", logs: logs, sub: sub}, nil
}

// WatchPayoff is a free log subscription operation binding the contract event 0xd68d5698fae2dc1d31bf2bf1ea674da9ebcc3989a5c8b6892414890a1de99a80.
//
// Solidity: event __payoff(bytes32 offchain)
func (_Constant *ConstantFilterer) WatchPayoff(opts *bind.WatchOpts, sink chan<- *ConstantPayoff) (event.Subscription, error) {

	logs, sub, err := _Constant.contract.WatchLogs(opts, "__payoff")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConstantPayoff)
				if err := _Constant.contract.UnpackLog(event, "__payoff", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ConstantPurchaseIterator is returned from FilterPurchase and is used to iterate over the raw logs and unpacked data for Purchase events raised by the Constant contract.
type ConstantPurchaseIterator struct {
	Event *ConstantPurchase // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConstantPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConstantPurchase)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConstantPurchase)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConstantPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConstantPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConstantPurchase represents a Purchase event raised by the Constant contract.
type ConstantPurchase struct {
	Offchain [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPurchase is a free log retrieval operation binding the contract event 0xef72b9890ab0fc46404e72534dc1bbc275de9d7efd8b3657ad91f82e1a3d39c4.
//
// Solidity: event __purchase(bytes32 offchain)
func (_Constant *ConstantFilterer) FilterPurchase(opts *bind.FilterOpts) (*ConstantPurchaseIterator, error) {

	logs, sub, err := _Constant.contract.FilterLogs(opts, "__purchase")
	if err != nil {
		return nil, err
	}
	return &ConstantPurchaseIterator{contract: _Constant.contract, event: "__purchase", logs: logs, sub: sub}, nil
}

// WatchPurchase is a free log subscription operation binding the contract event 0xef72b9890ab0fc46404e72534dc1bbc275de9d7efd8b3657ad91f82e1a3d39c4.
//
// Solidity: event __purchase(bytes32 offchain)
func (_Constant *ConstantFilterer) WatchPurchase(opts *bind.WatchOpts, sink chan<- *ConstantPurchase) (event.Subscription, error) {

	logs, sub, err := _Constant.contract.WatchLogs(opts, "__purchase")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConstantPurchase)
				if err := _Constant.contract.UnpackLog(event, "__purchase", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ConstantRedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the Constant contract.
type ConstantRedeemIterator struct {
	Event *ConstantRedeem // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConstantRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConstantRedeem)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConstantRedeem)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConstantRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConstantRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConstantRedeem represents a Redeem event raised by the Constant contract.
type ConstantRedeem struct {
	Offchain [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0xb0de879351469d2741406aafc9ba1f44eb957cf44ee3391e59a7a9097050c927.
//
// Solidity: event __redeem(bytes32 offchain)
func (_Constant *ConstantFilterer) FilterRedeem(opts *bind.FilterOpts) (*ConstantRedeemIterator, error) {

	logs, sub, err := _Constant.contract.FilterLogs(opts, "__redeem")
	if err != nil {
		return nil, err
	}
	return &ConstantRedeemIterator{contract: _Constant.contract, event: "__redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0xb0de879351469d2741406aafc9ba1f44eb957cf44ee3391e59a7a9097050c927.
//
// Solidity: event __redeem(bytes32 offchain)
func (_Constant *ConstantFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *ConstantRedeem) (event.Subscription, error) {

	logs, sub, err := _Constant.contract.WatchLogs(opts, "__redeem")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConstantRedeem)
				if err := _Constant.contract.UnpackLog(event, "__redeem", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ConstantTransferByAdminIterator is returned from FilterTransferByAdmin and is used to iterate over the raw logs and unpacked data for TransferByAdmin events raised by the Constant contract.
type ConstantTransferByAdminIterator struct {
	Event *ConstantTransferByAdmin // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConstantTransferByAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConstantTransferByAdmin)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConstantTransferByAdmin)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConstantTransferByAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConstantTransferByAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConstantTransferByAdmin represents a TransferByAdmin event raised by the Constant contract.
type ConstantTransferByAdmin struct {
	Offchain [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferByAdmin is a free log retrieval operation binding the contract event 0x3b36ee6b35325f38e95938557be92853c842b7a9a19fd7ac4931a6d24db52682.
//
// Solidity: event __transferByAdmin(bytes32 offchain)
func (_Constant *ConstantFilterer) FilterTransferByAdmin(opts *bind.FilterOpts) (*ConstantTransferByAdminIterator, error) {

	logs, sub, err := _Constant.contract.FilterLogs(opts, "__transferByAdmin")
	if err != nil {
		return nil, err
	}
	return &ConstantTransferByAdminIterator{contract: _Constant.contract, event: "__transferByAdmin", logs: logs, sub: sub}, nil
}

// WatchTransferByAdmin is a free log subscription operation binding the contract event 0x3b36ee6b35325f38e95938557be92853c842b7a9a19fd7ac4931a6d24db52682.
//
// Solidity: event __transferByAdmin(bytes32 offchain)
func (_Constant *ConstantFilterer) WatchTransferByAdmin(opts *bind.WatchOpts, sink chan<- *ConstantTransferByAdmin) (event.Subscription, error) {

	logs, sub, err := _Constant.contract.WatchLogs(opts, "__transferByAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConstantTransferByAdmin)
				if err := _Constant.contract.UnpackLog(event, "__transferByAdmin", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
