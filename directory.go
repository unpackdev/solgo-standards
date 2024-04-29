package standards

import "github.com/unpackdev/standards/shared"

// standards is a map that stores ContractStandard instances indexed by their Standard identifier.
var standards = map[shared.Standard]shared.ContractStandard{
	ERC20: {
		Name: "ERC-20 Token Standard",
		Url:  "https://eips.ethereum.org/EIPS/eip-20",
		Type: ERC20,
		ABI:  `[{"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_value","type":"uint256"}],"name":"approve","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_from","type":"address"},{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transferFrom","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"decimals","outputs":[{"name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"balance","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transfer","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"},{"name":"_spender","type":"address"}],"name":"allowance","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"payable":true,"stateMutability":"payable","type":"fallback"},{"anonymous":false,"inputs":[{"indexed":true,"name":"owner","type":"address"},{"indexed":true,"name":"spender","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Transfer","type":"event"}]`,
		Functions: []shared.Function{
			shared.NewFunction("totalSupply", nil, []shared.Output{{Type: shared.TypeUint256}}),
			shared.NewFunction("balanceOf", []shared.Input{{Type: shared.TypeAddress}}, []shared.Output{{Type: shared.TypeUint256}}),
			shared.NewFunction("transfer", []shared.Input{{Type: shared.TypeAddress}, {Type: shared.TypeUint256}}, []shared.Output{{Type: shared.TypeBool}}),
			shared.NewFunction("transferFrom", []shared.Input{{Type: shared.TypeAddress}, {Type: shared.TypeAddress}, {Type: shared.TypeUint256}}, []shared.Output{{Type: shared.TypeBool}}),
			shared.NewFunction("approve", []shared.Input{{Type: shared.TypeAddress}, {Type: shared.TypeUint256}}, []shared.Output{{Type: shared.TypeBool}}),
			shared.NewFunction("allowance", []shared.Input{{Type: shared.TypeAddress}, {Type: shared.TypeAddress}}, []shared.Output{{Type: shared.TypeUint256}}),
		},
		Events: []shared.Event{
			shared.NewEvent("Transfer", []shared.Input{{Type: shared.TypeAddress, Indexed: true}, {Type: shared.TypeAddress, Indexed: true}, {Type: shared.TypeUint256}}, nil),
			shared.NewEvent("Approval", []shared.Input{{Type: shared.TypeAddress, Indexed: true}, {Type: shared.TypeAddress, Indexed: true}, {Type: shared.TypeUint256}}, nil),
		},
	},
	ERC721: {
		Name: "ERC-721 Non-Fungible Token Standard",
		Url:  "https://eips.ethereum.org/EIPS/eip-721",
		Type: ERC721,
		ABI:  `[{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":true,"internalType":"address","name":"approved","type":"address"},{"indexed":true,"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":true,"internalType":"address","name":"operator","type":"address"},{"indexed":false,"internalType":"bool","name":"approved","type":"bool"}],"name":"ApprovalForAll","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":true,"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"Transfer","type":"event"},{"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"approve","outputs":[],"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"owner","type":"address"}],"name":"balanceOf","outputs":[{"internalType":"uint256","name":"balance","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"getApproved","outputs":[{"internalType":"address","name":"operator","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"address","name":"operator","type":"address"}],"name":"isApprovedForAll","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"name","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"ownerOf","outputs":[{"internalType":"address","name":"owner","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"safeTransferFrom","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"},{"internalType":"bytes","name":"data","type":"bytes"}],"name":"safeTransferFrom","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"operator","type":"address"},{"internalType":"bool","name":"_approved","type":"bool"}],"name":"setApprovalForAll","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bytes4","name":"interfaceId","type":"bytes4"}],"name":"supportsInterface","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"symbol","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"tokenURI","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"transferFrom","outputs":[],"stateMutability":"nonpayable","type":"function"}]`,
		Functions: []shared.Function{
			shared.NewFunction("name", nil, []shared.Output{{Type: shared.TypeString}}),
			shared.NewFunction("symbol", nil, []shared.Output{{Type: shared.TypeString}}),
			shared.NewFunction("totalSupply", nil, []shared.Output{{Type: shared.TypeUint256}}),
			shared.NewFunction("balanceOf", []shared.Input{{Type: shared.TypeAddress}}, []shared.Output{{Type: shared.TypeUint256}}),
			shared.NewFunction("ownerOf", []shared.Input{{Type: shared.TypeUint256}}, []shared.Output{{Type: shared.TypeAddress}}),
			shared.NewFunction("transferFrom", []shared.Input{{Type: shared.TypeAddress}, {Type: shared.TypeAddress}, {Type: shared.TypeUint256}}, nil),
			shared.NewFunction("approve", []shared.Input{{Type: shared.TypeAddress}, {Type: shared.TypeUint256}}, nil),
			shared.NewFunction("setApprovalForAll", []shared.Input{{Type: shared.TypeAddress}, {Type: shared.TypeBool}}, nil),
			shared.NewFunction("getApproved", []shared.Input{{Type: shared.TypeUint256}}, []shared.Output{{Type: shared.TypeAddress}}),
			shared.NewFunction("isApprovedForAll", []shared.Input{{Type: shared.TypeAddress}, {Type: shared.TypeAddress}}, []shared.Output{{Type: shared.TypeBool}}),
		},
		Events: []shared.Event{
			shared.NewEvent("Transfer", []shared.Input{{Type: shared.TypeAddress, Indexed: true}, {Type: shared.TypeAddress, Indexed: true}, {Type: shared.TypeUint256}}, nil),
			shared.NewEvent("Approval", []shared.Input{{Type: shared.TypeAddress, Indexed: true}, {Type: shared.TypeAddress, Indexed: true}, {Type: shared.TypeUint256}}, nil),
			shared.NewEvent("ApprovalForAll", []shared.Input{{Type: shared.TypeAddress, Indexed: true}, {Type: shared.TypeAddress, Indexed: true}, {Type: shared.TypeBool}}, nil),
		},
	},
	/*	ERC1155: {
			Name: "ERC-1155 Multi Token Standard",
			Url:  "https://eips.ethereum.org/EIPS/eip-1155",
			Type: ERC1155,
			ABI:  `[{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"account","type":"address"},{"indexed":true,"internalType":"address","name":"operator","type":"address"},{"indexed":false,"internalType":"bool","name":"approved","type":"bool"}],"name":"ApprovalForAll","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"operator","type":"address"},{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":false,"internalType":"uint256[]","name":"ids","type":"uint256[]"},{"indexed":false,"internalType":"uint256[]","name":"values","type":"uint256[]"}],"name":"TransferBatch","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"operator","type":"address"},{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":false,"internalType":"uint256","name":"id","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"TransferSingle","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"string","name":"value","type":"string"},{"indexed":true,"internalType":"uint256","name":"id","type":"uint256"}],"name":"URI","type":"event"},{"inputs":[{"internalType":"address","name":"account","type":"address"},{"internalType":"uint256","name":"id","type":"uint256"}],"name":"balanceOf","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address[]","name":"accounts","type":"address[]"},{"internalType":"uint256[]","name":"ids","type":"uint256[]"}],"name":"balanceOfBatch","outputs":[{"internalType":"uint256[]","name":"","type":"uint256[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"account","type":"address"},{"internalType":"address","name":"operator","type":"address"}],"name":"isApprovedForAll","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256[]","name":"ids","type":"uint256[]"},{"internalType":"uint256[]","name":"amounts","type":"uint256[]"},{"internalType":"bytes","name":"data","type":"bytes"}],"name":"safeBatchTransferFrom","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"id","type":"uint256"},{"internalType":"uint256","name":"amount","type":"uint256"},{"internalType":"bytes","name":"data","type":"bytes"}],"name":"safeTransferFrom","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"operator","type":"address"},{"internalType":"bool","name":"approved","type":"bool"}],"name":"setApprovalForAll","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bytes4","name":"interfaceId","type":"bytes4"}],"name":"supportsInterface","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"id","type":"uint256"}],"name":"uri","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`,
			Functions: []Function{
				newFunction("safeTransferFrom", []Input{{Type: TypeAddress}, {Type: TypeAddress}, {Type: TypeUint256}, {Type: TypeUint256}, {Type: TypeBytes}}, nil),
				newFunction("safeBatchTransferFrom", []Input{{Type: TypeAddress}, {Type: TypeAddress}, {Type: TypeUint256Array}, {Type: TypeUint256Array}, {Type: TypeBytes}}, nil),
				newFunction("balanceOf", []Input{{Type: TypeAddress}, {Type: TypeUint256}}, []Output{{Type: TypeUint256}}),
				newFunction("balanceOfBatch", []Input{{Type: TypeAddressArray}, {Type: TypeUint256Array}}, []Output{{Type: TypeUint256Array}}),
				newFunction("setApprovalForAll", []Input{{Type: TypeAddress}, {Type: TypeBool}}, nil),
				newFunction("isApprovedForAll", []Input{{Type: TypeAddress}, {Type: TypeAddress}}, []Output{{Type: TypeBool}}),
			},
			Events: []Event{
				newEvent("TransferSingle", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}, {Type: TypeUint256}, {Type: TypeUint256}}, nil),
				newEvent("TransferBatch", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}, {Type: TypeAddressArray, Indexed: true}, {Type: TypeUint256Array}, {Type: TypeUint256Array}}, nil),
				newEvent("ApprovalForAll", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}, {Type: TypeBool}}, nil),
				newEvent("URI", []Input{{Type: TypeString, Indexed: false}, {Type: TypeUint256, Indexed: true}}, nil),
			},
		},
		ERC1820: {
			Name: "ERC-1820 Pseudo-introspection Registry Contract",
			Url:  "https://eips.ethereum.org/EIPS/eip-1820",
			Type: ERC1820,
			ABI:  `[{"constant":true,"inputs":[{"name":"account","type":"address"},{"name":"interfaceHash","type":"bytes32"}],"name":"getInterfaceImplementer","outputs":[{"name":"implementer","type":"address"}],"type":"function"},{"constant":false,"inputs":[{"name":"interfaceHash","type":"bytes32"},{"name":"implementer","type":"address"}],"name":"setInterfaceImplementer","outputs":[],"type":"function"},{"constant":false,"inputs":[{"name":"newManager","type":"address"}],"name":"setManager","outputs":[],"type":"function"},{"constant":true,"inputs":[{"name":"account","type":"address"}],"name":"getManager","outputs":[{"name":"manager","type":"address"}],"type":"function"},{"constant":false,"inputs":[{"name":"account","type":"address"}],"name":"setInterfaceImplementer","outputs":[],"type":"function"},{"inputs":[],"type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"name":"account","type":"address"},{"indexed":true,"name":"interfaceHash","type":"bytes32"},{"indexed":true,"name":"implementer","type":"address"}],"name":"InterfaceImplementerSet","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"account","type":"address"},{"indexed":true,"name":"newManager","type":"address"}],"name":"ManagerChanged","type":"event"}]`,
			Functions: []Function{
				newFunction("setInterfaceImplementer", []Input{{Type: TypeAddress}, {Type: TypeBytes32}, {Type: TypeAddress}}, nil),
				newFunction("getInterfaceImplementer", []Input{{Type: TypeAddress}, {Type: TypeBytes32}}, []Output{{Type: TypeAddress}}),
				newFunction("interfaceHash", []Input{{Type: TypeString}}, []Output{{Type: TypeBytes32}}),
				newFunction("updateERC165Cache", []Input{{Type: TypeAddress}, {Type: TypeBytes32}}, nil),
				newFunction("implementsERC165InterfaceNoCache", []Input{{Type: TypeAddress}, {Type: TypeBytes32}}, []Output{{Type: TypeBool}}),
				newFunction("implementsERC165Interface", []Input{{Type: TypeAddress}, {Type: TypeBytes32}}, []Output{{Type: TypeBool}}),
			},
			Events: []Event{
				newEvent("InterfaceImplementerSet", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeBytes32, Indexed: true}, {Type: TypeAddress, Indexed: true}}, nil),
				newEvent("ManagerChanged", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}}, nil),
			},
		},
		ERC1822: {
			Name:     "ERC-1822 Universal Proxy Standard (UPS)",
			Url:      "https://eips.ethereum.org/EIPS/eip-1822",
			Stagnant: true,
			Type:     ERC1822,
			ABI:      `[{"constant":true,"inputs":[],"name":"getImplementation","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"","type":"address"}],"name":"upgradeTo","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"","type":"address"},{"name":"","type":"string"}],"name":"upgradeToAndCall","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"","type":"address"}],"name":"setProxyOwner","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"anonymous":false,"inputs":[{"indexed":true,"name":"","type":"address"}],"name":"Upgraded","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"","type":"address"},{"indexed":true,"name":"","type":"address"}],"name":"ProxyOwnershipTransferred","type":"event"}]`,
			Functions: []Function{
				newFunction("getImplementation", nil, []Output{{Type: TypeAddress}}),
				newFunction("upgradeTo", []Input{{Type: TypeAddress}}, nil),
				newFunction("upgradeToAndCall", []Input{{Type: TypeAddress, Indexed: false}, {Type: TypeString, Indexed: false}}, nil),
				newFunction("setProxyOwner", []Input{{Type: TypeAddress}}, nil),
			},
			Events: []Event{
				newEvent("Upgraded", []Input{{Type: TypeAddress, Indexed: true}}, nil),
				newEvent("ProxyOwnershipTransferred", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}}, nil),
			},
		},
		ERC1967: {
			Name: "ERC-1967 Proxy Storage Slots",
			Url:  "https://eips.ethereum.org/EIPS/eip-1967",
			Type: ERC1967,
			ABI:  `[{"constant":false,"inputs":[{"name":"","type":"address"},{"name":"","type":"bytes32"},{"name":"","type":"address"}],"name":"setInterfaceImplementer","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"},{"name":"","type":"bytes32"}],"name":"getInterfaceImplementer","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"string"}],"name":"interfaceHash","outputs":[{"name":"","type":"bytes32"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"","type":"address"},{"name":"","type":"bytes32"}],"name":"updateERC165Cache","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"},{"name":"","type":"bytes32"}],"name":"implementsERC165InterfaceNoCache","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"},{"name":"","type":"bytes32"}],"name":"implementsERC165Interface","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"anonymous":false,"inputs":[{"indexed":true,"name":"","type":"address"},{"indexed":true,"name":"","type":"bytes32"},{"indexed":true,"name":"","type":"address"}],"name":"InterfaceImplementerSet","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"","type":"address"},{"indexed":true,"name":"","type":"address"}],"name":"AdminChanged","type":"event"}]`,
			Functions: []Function{
				newFunction("setInterfaceImplementer", []Input{{Type: TypeAddress}, {Type: TypeBytes32}, {Type: TypeAddress}}, nil),
				newFunction("getInterfaceImplementer", []Input{{Type: TypeAddress}, {Type: TypeBytes32}}, []Output{{Type: TypeAddress}}),
				newFunction("interfaceHash", []Input{{Type: TypeString}}, []Output{{Type: TypeBytes32}}),
				newFunction("updateERC165Cache", []Input{{Type: TypeAddress}, {Type: TypeBytes32}}, nil),
				newFunction("implementsERC165InterfaceNoCache", []Input{{Type: TypeAddress}, {Type: TypeBytes32}}, []Output{{Type: TypeBool}}),
				newFunction("implementsERC165Interface", []Input{{Type: TypeAddress}, {Type: TypeBytes32}}, []Output{{Type: TypeBool}}),
			},
			Events: []Event{
				newEvent("InterfaceImplementerSet", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeBytes32, Indexed: true}, {Type: TypeAddress, Indexed: true}}, nil),
				newEvent("AdminChanged", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}}, nil),
			},
		},
		OZOWNABLE: {
			Name: "OpenZeppelin Owner Module",
			Url:  "https://docs.openzeppelin.com/contracts/4.x/api/access#Ownable",
			Type: OZOWNABLE,
			ABI:  `[{"inputs":[{"internalType":"address","name":"owner","type":"address"}],"name":"OwnableInvalidOwner","type":"error"},{"inputs":[{"internalType":"address","name":"account","type":"address"}],"name":"OwnableUnauthorizedAccount","type":"error"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"previousOwner","type":"address"},{"indexed":true,"internalType":"address","name":"newOwner","type":"address"}],"name":"OwnershipTransferStarted","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"previousOwner","type":"address"},{"indexed":true,"internalType":"address","name":"newOwner","type":"address"}],"name":"OwnershipTransferred","type":"event"},{"inputs":[],"name":"acceptOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"pendingOwner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"renounceOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"}]`,
			Functions: []Function{
				newFunction("acceptOwnership", nil, nil),
				newFunction("owner", nil, []Output{{Type: TypeAddress}}),
				newFunction("pendingOwner", nil, []Output{{Type: TypeAddress}}),
				newFunction("renounceOwnership", nil, nil),
				newFunction("transferOwnership", []Input{{Type: TypeAddress}}, nil),
			},
			Events: []Event{
				newEvent("OwnershipTransferStarted", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}}, nil),
				newEvent("OwnershipTransferred", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}}, nil),
			},
		},*/
}
