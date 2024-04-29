package standards

import (
	"github.com/unpackdev/standards/contracts"
	"github.com/unpackdev/standards/errors"
	"github.com/unpackdev/standards/shared"
)

// Constants representing various Ethereum standards and EIPs.
const (
	ERC20     shared.Standard = "ERC20"     // ERC-20 Token Standard.
	ERC721    shared.Standard = "ERC721"    // ERC-721 Non-Fungible Token Standard.
	ERC1822   shared.Standard = "ERC1822"   // ERC-1822 Universal Proxy Standard (UPS).
	ERC1820   shared.Standard = "ERC1820"   // ERC-1820 Pseudo-introspection Registry Contract.
	ERC777    shared.Standard = "ERC777"    // ERC-777 Token Standard.
	ERC1155   shared.Standard = "ERC1155"   // ERC-1155 Multi Token Standard.
	ERC1337   shared.Standard = "ERC1337"   // ERC-1337 Subscription Standard.
	ERC1400   shared.Standard = "ERC1400"   // ERC-1400 Security Token Standard.
	ERC1410   shared.Standard = "ERC1410"   // ERC-1410 Partially Fungible Token Standard.
	ERC165    shared.Standard = "ERC165"    // ERC-165 Standard Interface Detection.
	ERC820    shared.Standard = "ERC820"    // ERC-820 Registry Standard.
	ERC1014   shared.Standard = "ERC1014"   // ERC-1014 Create2 Standard.
	ERC1948   shared.Standard = "ERC1948"   // ERC-1948 Non-Fungible Data Token Standard.
	ERC1967   shared.Standard = "ERC1967"   // ERC-1967 Proxy Storage Slots Standard.
	ERC2309   shared.Standard = "ERC2309"   // ERC-2309 Consecutive Transfer Standard.
	ERC2535   shared.Standard = "ERC2535"   // ERC-2535 Diamond Standard.
	ERC2771   shared.Standard = "ERC2771"   // ERC-2771 Meta Transactions Standard.
	ERC2917   shared.Standard = "ERC2917"   // ERC-2917 Interest-Bearing Tokens Standard.
	ERC3156   shared.Standard = "ERC3156"   // ERC-3156 Flash Loans Standard.
	ERC3664   shared.Standard = "ERC3664"   // ERC-3664 BitWords Standard.
	UNISWAPV2 shared.Standard = "UNISWAPV2" // Uniswap V2 Core.
	OZOWNABLE shared.Standard = "OZOWNABLE" // OpenZeppelin Ownable.
)

func GetContractByStandard(standard shared.Standard) (shared.EIP, error) {
	if standard, ok := standards[standard]; ok {
		return contracts.NewContract(standard), nil
	}
	return nil, errors.ErrStandardNotFound
}

// LoadStandards loads list of supported Ethereum EIPs into the registry.
func LoadStandards() error {
	for name, standard := range standards {
		if err := RegisterStandard(name, contracts.NewContract(standard)); err != nil {
			return err
		}
	}

	return nil
}
