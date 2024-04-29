package shared

import eip_pb "github.com/unpackdev/protos/dist/go/eip"

// ConfidenceLevel represents the confidence level of a discovery.
type ConfidenceLevel int

// String returns the string representation of the confidence level.
func (c ConfidenceLevel) String() string {
	switch c {
	case PerfectConfidence:
		return "perfect"
	case HighConfidence:
		return "high"
	case MediumConfidence:
		return "medium"
	case LowConfidence:
		return "low"
	case NoConfidence:
		return "no_confidence"
	default:
		return "unknown"
	}
}

// ToProto converts a ConfidenceLevel to its protobuf representation.
func (c ConfidenceLevel) ToProto() eip_pb.ConfidenceLevel {
	return eip_pb.ConfidenceLevel(c)
}

// ConfidenceThreshold represents the threshold value for a confidence level.
type ConfidenceThreshold float64

// ToProto converts a ConfidenceThreshold to its protobuf representation.
func (c ConfidenceThreshold) ToProto() eip_pb.ConfidenceThreshold {
	return eip_pb.ConfidenceThreshold(c)
}

const (
	// PerfectConfidenceThreshold represents a perfect confidence threshold value.
	PerfectConfidenceThreshold ConfidenceThreshold = 1.0

	// HighConfidenceThreshold represents a high confidence threshold value.
	HighConfidenceThreshold ConfidenceThreshold = 0.9

	// MediumConfidenceThreshold represents a medium confidence threshold value.
	MediumConfidenceThreshold ConfidenceThreshold = 0.5

	// LowConfidenceThreshold represents a low confidence threshold value.
	LowConfidenceThreshold ConfidenceThreshold = 0.1

	// NoConfidenceThreshold represents no confidence threshold value.
	NoConfidenceThreshold ConfidenceThreshold = 0.0

	// PerfectConfidence represents a perfect confidence level.
	PerfectConfidence ConfidenceLevel = 4

	// HighConfidence represents a high confidence level.
	HighConfidence ConfidenceLevel = 3

	// MediumConfidence represents a medium confidence level.
	MediumConfidence ConfidenceLevel = 2

	// LowConfidence represents a low confidence level.
	LowConfidence ConfidenceLevel = 1

	// NoConfidence represents no confidence level.
	NoConfidence ConfidenceLevel = 0
)
