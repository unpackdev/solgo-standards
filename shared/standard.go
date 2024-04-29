package shared

import (
	eip_pb "github.com/unpackdev/protos/dist/go/eip"
	"strings"
)

// Standard represents the type for Ethereum standards and EIPs.
type Standard string

// String returns the string representation of the Standard.
func (s Standard) String() string {
	return string(s)
}

// ToProto converts a string representation of an Ethereum standard
// to its corresponding protobuf enum value. If the standard is not recognized,
// it returns unknown.
func (s Standard) ToProto() eip_pb.Standard {
	if standardValue, ok := eip_pb.Standard_value[strings.ToUpper(string(s))]; ok {
		return eip_pb.Standard(standardValue)
	}
	return eip_pb.Standard_UNKNOWN
}
