package standards

import (
	eip_pb "github.com/unpackdev/protos/dist/go/eip"
	"github.com/unpackdev/standards/utils"
)

// EIP is an interface that defines the standard methods required for
// representing Ethereum Improvement Proposals and Ethereum standards.
type EIP interface {
	// GetName returns the name of the Ethereum standard, e.g., "ERC-20 Token Standard".
	GetName() string

	// GetType returns the type of the Ethereum standard, e.g., ERC20 or ERC721.
	GetType() utils.Standard

	// GetFunctions returns a slice of Function structs, representing the
	// functions defined in the Ethereum standard.
	GetFunctions() []utils.Function

	// GetUrl returns the URL of the Ethereum standard.
	GetUrl() string

	// IsStagnant returns a boolean indicating whether the Ethereum standard is stagnant.
	IsStagnant() bool

	// GetEvents returns a slice of Event structs, representing the
	// events defined in the Ethereum standard.
	GetEvents() []utils.Event

	// GetStandard returns the complete representation of the Ethereum standard.
	GetStandard() utils.ContractStandard

	// ConfidenceCheck returns a discovery confidence information and a boolean indicating whether
	// the contract is to any level compliant with the Ethereum standard.
	ConfidenceCheck(contract *utils.ContractMatcher) (utils.Discovery, bool)

	// FunctionConfidenceCheck returns a discovery confidence information and a boolean indicating whether
	// the contract function is to any level compliant with the Ethereum standard.
	FunctionConfidenceCheck(fn *utils.Function) (utils.FunctionDiscovery, bool)

	// TokenCount returns the number of tokens associated with the Ethereum standard.
	TokenCount() int

	FunctionTokenCount(fnName string) int

	// GetABI returns the ABI of the Ethereum standard.
	GetABI() string

	// ToProto converts the Ethereum standard to its protobuf representation.
	ToProto() *eip_pb.ContractStandard

	// String returns a string representation of the Ethereum standard, typically its name.
	String() string
}
