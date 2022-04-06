// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package unlocked

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
)

// UnlockedMetaData contains all meta data concerning the Unlocked contract.
var UnlockedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040518060400160405280600781526020017f556e6c6f636b64000000000000000000000000000000000000000000000000008152506040518060400160405280600381526020017f554c4b000000000000000000000000000000000000000000000000000000000081525081600390805190602001906200009692919062000370565b508060049080519060200190620000af92919062000370565b505050620000d2620000c66200011760201b60201c565b6200011f60201b60201c565b6200011133620000e7620001e560201b60201c565b600a620000f59190620005ba565b629896806200010591906200060b565b620001ee60201b60201c565b620007de565b600033905090565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b60006012905090565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160362000260576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200025790620006cd565b60405180910390fd5b62000274600083836200036660201b60201c565b8060026000828254620002889190620006ef565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254620002df9190620006ef565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516200034691906200075d565b60405180910390a362000362600083836200036b60201b60201c565b5050565b505050565b505050565b8280546200037e90620007a9565b90600052602060002090601f016020900481019282620003a25760008555620003ee565b82601f10620003bd57805160ff1916838001178555620003ee565b82800160010185558215620003ee579182015b82811115620003ed578251825591602001919060010190620003d0565b5b509050620003fd919062000401565b5090565b5b808211156200041c57600081600090555060010162000402565b5090565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008160011c9050919050565b6000808291508390505b6001851115620004ae5780860481111562000486576200048562000420565b5b6001851615620004965780820291505b8081029050620004a6856200044f565b945062000466565b94509492505050565b600082620004c957600190506200059c565b81620004d957600090506200059c565b8160018114620004f25760028114620004fd5762000533565b60019150506200059c565b60ff84111562000512576200051162000420565b5b8360020a9150848211156200052c576200052b62000420565b5b506200059c565b5060208310610133831016604e8410600b84101617156200056d5782820a90508381111562000567576200056662000420565b5b6200059c565b6200057c84848460016200045c565b9250905081840481111562000596576200059562000420565b5b81810290505b9392505050565b6000819050919050565b600060ff82169050919050565b6000620005c782620005a3565b9150620005d483620005ad565b9250620006037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484620004b7565b905092915050565b60006200061882620005a3565b91506200062583620005a3565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161562000661576200066062000420565b5b828202905092915050565b600082825260208201905092915050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b6000620006b5601f836200066c565b9150620006c2826200067d565b602082019050919050565b60006020820190508181036000830152620006e881620006a6565b9050919050565b6000620006fc82620005a3565b91506200070983620005a3565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111562000741576200074062000420565b5b828201905092915050565b6200075781620005a3565b82525050565b60006020820190506200077460008301846200074c565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620007c257607f821691505b602082108103620007d857620007d76200077a565b5b50919050565b6119e180620007ee6000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806370a0823111610097578063a457c2d711610066578063a457c2d714610276578063a9059cbb146102a6578063dd62ed3e146102d6578063f2fde38b14610306576100f5565b806370a0823114610200578063715018a6146102305780638da5cb5b1461023a57806395d89b4114610258576100f5565b806323b872dd116100d357806323b872dd14610166578063313ce5671461019657806339509351146101b457806340c10f19146101e4576100f5565b806306fdde03146100fa578063095ea7b31461011857806318160ddd14610148575b600080fd5b610102610322565b60405161010f9190611108565b60405180910390f35b610132600480360381019061012d91906111c3565b6103b4565b60405161013f919061121e565b60405180910390f35b6101506103d7565b60405161015d9190611248565b60405180910390f35b610180600480360381019061017b9190611263565b6103e1565b60405161018d919061121e565b60405180910390f35b61019e610410565b6040516101ab91906112d2565b60405180910390f35b6101ce60048036038101906101c991906111c3565b610419565b6040516101db919061121e565b60405180910390f35b6101fe60048036038101906101f991906111c3565b6104c3565b005b61021a600480360381019061021591906112ed565b61054d565b6040516102279190611248565b60405180910390f35b610238610595565b005b61024261061d565b60405161024f9190611329565b60405180910390f35b610260610647565b60405161026d9190611108565b60405180910390f35b610290600480360381019061028b91906111c3565b6106d9565b60405161029d919061121e565b60405180910390f35b6102c060048036038101906102bb91906111c3565b6107c3565b6040516102cd919061121e565b60405180910390f35b6102f060048036038101906102eb9190611344565b6107e6565b6040516102fd9190611248565b60405180910390f35b610320600480360381019061031b91906112ed565b61086d565b005b606060038054610331906113b3565b80601f016020809104026020016040519081016040528092919081815260200182805461035d906113b3565b80156103aa5780601f1061037f576101008083540402835291602001916103aa565b820191906000526020600020905b81548152906001019060200180831161038d57829003601f168201915b5050505050905090565b6000806103bf610964565b90506103cc81858561096c565b600191505092915050565b6000600254905090565b6000806103ec610964565b90506103f9858285610b35565b610404858585610bc1565b60019150509392505050565b60006012905090565b600080610424610964565b90506104b8818585600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546104b39190611413565b61096c565b600191505092915050565b6104cb610964565b73ffffffffffffffffffffffffffffffffffffffff166104e961061d565b73ffffffffffffffffffffffffffffffffffffffff161461053f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610536906114b5565b60405180910390fd5b6105498282610e40565b5050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b61059d610964565b73ffffffffffffffffffffffffffffffffffffffff166105bb61061d565b73ffffffffffffffffffffffffffffffffffffffff1614610611576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610608906114b5565b60405180910390fd5b61061b6000610f9f565b565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b606060048054610656906113b3565b80601f0160208091040260200160405190810160405280929190818152602001828054610682906113b3565b80156106cf5780601f106106a4576101008083540402835291602001916106cf565b820191906000526020600020905b8154815290600101906020018083116106b257829003601f168201915b5050505050905090565b6000806106e4610964565b90506000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050838110156107aa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107a190611547565b60405180910390fd5b6107b7828686840361096c565b60019250505092915050565b6000806107ce610964565b90506107db818585610bc1565b600191505092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b610875610964565b73ffffffffffffffffffffffffffffffffffffffff1661089361061d565b73ffffffffffffffffffffffffffffffffffffffff16146108e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108e0906114b5565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610958576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161094f906115d9565b60405180910390fd5b61096181610f9f565b50565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036109db576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109d29061166b565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610a4a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a41906116fd565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92583604051610b289190611248565b60405180910390a3505050565b6000610b4184846107e6565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610bbb5781811015610bad576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ba490611769565b60405180910390fd5b610bba848484840361096c565b5b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610c30576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c27906117fb565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610c9f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c969061188d565b60405180910390fd5b610caa838383611065565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610d30576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d279061191f565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610dc39190611413565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610e279190611248565b60405180910390a3610e3a84848461106a565b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610eaf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ea69061198b565b60405180910390fd5b610ebb60008383611065565b8060026000828254610ecd9190611413565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610f229190611413565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610f879190611248565b60405180910390a3610f9b6000838361106a565b5050565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b505050565b505050565b600081519050919050565b600082825260208201905092915050565b60005b838110156110a957808201518184015260208101905061108e565b838111156110b8576000848401525b50505050565b6000601f19601f8301169050919050565b60006110da8261106f565b6110e4818561107a565b93506110f481856020860161108b565b6110fd816110be565b840191505092915050565b6000602082019050818103600083015261112281846110cf565b905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061115a8261112f565b9050919050565b61116a8161114f565b811461117557600080fd5b50565b60008135905061118781611161565b92915050565b6000819050919050565b6111a08161118d565b81146111ab57600080fd5b50565b6000813590506111bd81611197565b92915050565b600080604083850312156111da576111d961112a565b5b60006111e885828601611178565b92505060206111f9858286016111ae565b9150509250929050565b60008115159050919050565b61121881611203565b82525050565b6000602082019050611233600083018461120f565b92915050565b6112428161118d565b82525050565b600060208201905061125d6000830184611239565b92915050565b60008060006060848603121561127c5761127b61112a565b5b600061128a86828701611178565b935050602061129b86828701611178565b92505060406112ac868287016111ae565b9150509250925092565b600060ff82169050919050565b6112cc816112b6565b82525050565b60006020820190506112e760008301846112c3565b92915050565b6000602082840312156113035761130261112a565b5b600061131184828501611178565b91505092915050565b6113238161114f565b82525050565b600060208201905061133e600083018461131a565b92915050565b6000806040838503121561135b5761135a61112a565b5b600061136985828601611178565b925050602061137a85828601611178565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806113cb57607f821691505b6020821081036113de576113dd611384565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061141e8261118d565b91506114298361118d565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561145e5761145d6113e4565b5b828201905092915050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b600061149f60208361107a565b91506114aa82611469565b602082019050919050565b600060208201905081810360008301526114ce81611492565b9050919050565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b600061153160258361107a565b915061153c826114d5565b604082019050919050565b6000602082019050818103600083015261156081611524565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b60006115c360268361107a565b91506115ce82611567565b604082019050919050565b600060208201905081810360008301526115f2816115b6565b9050919050565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b600061165560248361107a565b9150611660826115f9565b604082019050919050565b6000602082019050818103600083015261168481611648565b9050919050565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b60006116e760228361107a565b91506116f28261168b565b604082019050919050565b60006020820190508181036000830152611716816116da565b9050919050565b7f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000600082015250565b6000611753601d8361107a565b915061175e8261171d565b602082019050919050565b6000602082019050818103600083015261178281611746565b9050919050565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b60006117e560258361107a565b91506117f082611789565b604082019050919050565b60006020820190508181036000830152611814816117d8565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b600061187760238361107a565b91506118828261181b565b604082019050919050565b600060208201905081810360008301526118a68161186a565b9050919050565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b600061190960268361107a565b9150611914826118ad565b604082019050919050565b60006020820190508181036000830152611938816118fc565b9050919050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b6000611975601f8361107a565b91506119808261193f565b602082019050919050565b600060208201905081810360008301526119a481611968565b905091905056fea2646970667358221220f9a4e8775ec86d5af80df6d719cc58a26973ea1f9857033e990ffe6fd06dee0064736f6c634300080d0033",
}

// UnlockedABI is the input ABI used to generate the binding from.
// Deprecated: Use UnlockedMetaData.ABI instead.
var UnlockedABI = UnlockedMetaData.ABI

// UnlockedBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UnlockedMetaData.Bin instead.
var UnlockedBin = UnlockedMetaData.Bin

// DeployUnlocked deploys a new Ethereum contract, binding an instance of Unlocked to it.
func DeployUnlocked(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Unlocked, error) {
	parsed, err := UnlockedMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UnlockedBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Unlocked{UnlockedCaller: UnlockedCaller{contract: contract}, UnlockedTransactor: UnlockedTransactor{contract: contract}, UnlockedFilterer: UnlockedFilterer{contract: contract}}, nil
}

// Unlocked is an auto generated Go binding around an Ethereum contract.
type Unlocked struct {
	UnlockedCaller     // Read-only binding to the contract
	UnlockedTransactor // Write-only binding to the contract
	UnlockedFilterer   // Log filterer for contract events
}

// UnlockedCaller is an auto generated read-only Go binding around an Ethereum contract.
type UnlockedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnlockedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UnlockedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnlockedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UnlockedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnlockedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UnlockedSession struct {
	Contract     *Unlocked         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UnlockedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UnlockedCallerSession struct {
	Contract *UnlockedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// UnlockedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UnlockedTransactorSession struct {
	Contract     *UnlockedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// UnlockedRaw is an auto generated low-level Go binding around an Ethereum contract.
type UnlockedRaw struct {
	Contract *Unlocked // Generic contract binding to access the raw methods on
}

// UnlockedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UnlockedCallerRaw struct {
	Contract *UnlockedCaller // Generic read-only contract binding to access the raw methods on
}

// UnlockedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UnlockedTransactorRaw struct {
	Contract *UnlockedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUnlocked creates a new instance of Unlocked, bound to a specific deployed contract.
func NewUnlocked(address common.Address, backend bind.ContractBackend) (*Unlocked, error) {
	contract, err := bindUnlocked(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Unlocked{UnlockedCaller: UnlockedCaller{contract: contract}, UnlockedTransactor: UnlockedTransactor{contract: contract}, UnlockedFilterer: UnlockedFilterer{contract: contract}}, nil
}

// NewUnlockedCaller creates a new read-only instance of Unlocked, bound to a specific deployed contract.
func NewUnlockedCaller(address common.Address, caller bind.ContractCaller) (*UnlockedCaller, error) {
	contract, err := bindUnlocked(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UnlockedCaller{contract: contract}, nil
}

// NewUnlockedTransactor creates a new write-only instance of Unlocked, bound to a specific deployed contract.
func NewUnlockedTransactor(address common.Address, transactor bind.ContractTransactor) (*UnlockedTransactor, error) {
	contract, err := bindUnlocked(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UnlockedTransactor{contract: contract}, nil
}

// NewUnlockedFilterer creates a new log filterer instance of Unlocked, bound to a specific deployed contract.
func NewUnlockedFilterer(address common.Address, filterer bind.ContractFilterer) (*UnlockedFilterer, error) {
	contract, err := bindUnlocked(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UnlockedFilterer{contract: contract}, nil
}

// bindUnlocked binds a generic wrapper to an already deployed contract.
func bindUnlocked(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UnlockedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Unlocked *UnlockedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Unlocked.Contract.UnlockedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Unlocked *UnlockedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Unlocked.Contract.UnlockedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Unlocked *UnlockedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Unlocked.Contract.UnlockedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Unlocked *UnlockedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Unlocked.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Unlocked *UnlockedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Unlocked.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Unlocked *UnlockedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Unlocked.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Unlocked *UnlockedCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Unlocked.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Unlocked *UnlockedSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Unlocked.Contract.Allowance(&_Unlocked.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Unlocked *UnlockedCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Unlocked.Contract.Allowance(&_Unlocked.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Unlocked *UnlockedCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Unlocked.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Unlocked *UnlockedSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Unlocked.Contract.BalanceOf(&_Unlocked.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Unlocked *UnlockedCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Unlocked.Contract.BalanceOf(&_Unlocked.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Unlocked *UnlockedCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Unlocked.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Unlocked *UnlockedSession) Decimals() (uint8, error) {
	return _Unlocked.Contract.Decimals(&_Unlocked.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Unlocked *UnlockedCallerSession) Decimals() (uint8, error) {
	return _Unlocked.Contract.Decimals(&_Unlocked.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Unlocked *UnlockedCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Unlocked.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Unlocked *UnlockedSession) Name() (string, error) {
	return _Unlocked.Contract.Name(&_Unlocked.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Unlocked *UnlockedCallerSession) Name() (string, error) {
	return _Unlocked.Contract.Name(&_Unlocked.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Unlocked *UnlockedCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Unlocked.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Unlocked *UnlockedSession) Owner() (common.Address, error) {
	return _Unlocked.Contract.Owner(&_Unlocked.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Unlocked *UnlockedCallerSession) Owner() (common.Address, error) {
	return _Unlocked.Contract.Owner(&_Unlocked.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Unlocked *UnlockedCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Unlocked.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Unlocked *UnlockedSession) Symbol() (string, error) {
	return _Unlocked.Contract.Symbol(&_Unlocked.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Unlocked *UnlockedCallerSession) Symbol() (string, error) {
	return _Unlocked.Contract.Symbol(&_Unlocked.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Unlocked *UnlockedCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Unlocked.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Unlocked *UnlockedSession) TotalSupply() (*big.Int, error) {
	return _Unlocked.Contract.TotalSupply(&_Unlocked.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Unlocked *UnlockedCallerSession) TotalSupply() (*big.Int, error) {
	return _Unlocked.Contract.TotalSupply(&_Unlocked.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Unlocked *UnlockedTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Unlocked.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Unlocked *UnlockedSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Unlocked.Contract.Approve(&_Unlocked.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Unlocked *UnlockedTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Unlocked.Contract.Approve(&_Unlocked.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Unlocked *UnlockedTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Unlocked.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Unlocked *UnlockedSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Unlocked.Contract.DecreaseAllowance(&_Unlocked.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Unlocked *UnlockedTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Unlocked.Contract.DecreaseAllowance(&_Unlocked.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Unlocked *UnlockedTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Unlocked.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Unlocked *UnlockedSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Unlocked.Contract.IncreaseAllowance(&_Unlocked.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Unlocked *UnlockedTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Unlocked.Contract.IncreaseAllowance(&_Unlocked.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Unlocked *UnlockedTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Unlocked.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Unlocked *UnlockedSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Unlocked.Contract.Mint(&_Unlocked.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Unlocked *UnlockedTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Unlocked.Contract.Mint(&_Unlocked.TransactOpts, to, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Unlocked *UnlockedTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Unlocked.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Unlocked *UnlockedSession) RenounceOwnership() (*types.Transaction, error) {
	return _Unlocked.Contract.RenounceOwnership(&_Unlocked.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Unlocked *UnlockedTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Unlocked.Contract.RenounceOwnership(&_Unlocked.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Unlocked *UnlockedTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Unlocked.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Unlocked *UnlockedSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Unlocked.Contract.Transfer(&_Unlocked.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Unlocked *UnlockedTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Unlocked.Contract.Transfer(&_Unlocked.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Unlocked *UnlockedTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Unlocked.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Unlocked *UnlockedSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Unlocked.Contract.TransferFrom(&_Unlocked.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Unlocked *UnlockedTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Unlocked.Contract.TransferFrom(&_Unlocked.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Unlocked *UnlockedTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Unlocked.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Unlocked *UnlockedSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Unlocked.Contract.TransferOwnership(&_Unlocked.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Unlocked *UnlockedTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Unlocked.Contract.TransferOwnership(&_Unlocked.TransactOpts, newOwner)
}

// UnlockedApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Unlocked contract.
type UnlockedApprovalIterator struct {
	Event *UnlockedApproval // Event containing the contract specifics and raw log

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
func (it *UnlockedApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnlockedApproval)
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
		it.Event = new(UnlockedApproval)
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
func (it *UnlockedApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnlockedApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnlockedApproval represents a Approval event raised by the Unlocked contract.
type UnlockedApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Unlocked *UnlockedFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*UnlockedApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Unlocked.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &UnlockedApprovalIterator{contract: _Unlocked.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Unlocked *UnlockedFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *UnlockedApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Unlocked.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnlockedApproval)
				if err := _Unlocked.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Unlocked *UnlockedFilterer) ParseApproval(log types.Log) (*UnlockedApproval, error) {
	event := new(UnlockedApproval)
	if err := _Unlocked.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnlockedOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Unlocked contract.
type UnlockedOwnershipTransferredIterator struct {
	Event *UnlockedOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *UnlockedOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnlockedOwnershipTransferred)
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
		it.Event = new(UnlockedOwnershipTransferred)
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
func (it *UnlockedOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnlockedOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnlockedOwnershipTransferred represents a OwnershipTransferred event raised by the Unlocked contract.
type UnlockedOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Unlocked *UnlockedFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UnlockedOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Unlocked.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UnlockedOwnershipTransferredIterator{contract: _Unlocked.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Unlocked *UnlockedFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UnlockedOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Unlocked.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnlockedOwnershipTransferred)
				if err := _Unlocked.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Unlocked *UnlockedFilterer) ParseOwnershipTransferred(log types.Log) (*UnlockedOwnershipTransferred, error) {
	event := new(UnlockedOwnershipTransferred)
	if err := _Unlocked.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnlockedTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Unlocked contract.
type UnlockedTransferIterator struct {
	Event *UnlockedTransfer // Event containing the contract specifics and raw log

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
func (it *UnlockedTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnlockedTransfer)
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
		it.Event = new(UnlockedTransfer)
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
func (it *UnlockedTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnlockedTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnlockedTransfer represents a Transfer event raised by the Unlocked contract.
type UnlockedTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Unlocked *UnlockedFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*UnlockedTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Unlocked.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &UnlockedTransferIterator{contract: _Unlocked.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Unlocked *UnlockedFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *UnlockedTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Unlocked.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnlockedTransfer)
				if err := _Unlocked.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Unlocked *UnlockedFilterer) ParseTransfer(log types.Log) (*UnlockedTransfer, error) {
	event := new(UnlockedTransfer)
	if err := _Unlocked.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
