package standards

import "github.com/unpackdev/standards/utils"

func init() {
	// Initialize the storage map so it can be accessed globally.
	storage = make(map[utils.Standard]EIP)
}
