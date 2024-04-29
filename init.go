package standards

import "github.com/unpackdev/standards/shared"

func init() {
	// Initialize the storage map so it can be accessed globally.
	storage = make(map[shared.Standard]shared.EIP)
}
