package confidence

import (
	"github.com/unpackdev/standards/shared"
)

// CalculateDiscoveryConfidence calculates the confidence level and threshold based on the total confidence.
func CalculateDiscoveryConfidence(totalConfidence float64) (shared.ConfidenceLevel, shared.ConfidenceThreshold) {
	total := shared.ConfidenceThreshold(totalConfidence)
	switch {
	case total == shared.PerfectConfidenceThreshold:
		return shared.PerfectConfidence, shared.PerfectConfidenceThreshold
	case total >= shared.HighConfidenceThreshold:
		return shared.HighConfidence, shared.HighConfidenceThreshold
	case total >= shared.MediumConfidenceThreshold:
		return shared.MediumConfidence, shared.MediumConfidenceThreshold
	case total >= shared.LowConfidenceThreshold:
		return shared.LowConfidence, shared.LowConfidenceThreshold
	default:
		return shared.NoConfidence, shared.NoConfidenceThreshold
	}
}

// ConfidenceCheck checks the confidence of a contract against a standard EIP.
func ConfidenceCheck(standard shared.EIP, contract *shared.ContractMatcher) (shared.Discovery, bool) {
	toReturn := shared.Discovery{
		Standard:         standard.GetType(),
		Confidence:       shared.NoConfidence,
		ConfidencePoints: 0,
		Threshold:        shared.NoConfidenceThreshold,
		MaximumTokens:    standard.TokenCount(),
		DiscoveredTokens: 0,
		Contract: &shared.ContractMatcher{
			Name:      contract.Name,
			Functions: make([]shared.Function, 0),
			Events:    make([]shared.Event, 0),
		},
	}
	foundTokenCount := 0
	discoveredFunctions := map[string]bool{}
	discoveredEvents := map[string]bool{}

	for _, standardFunction := range standard.GetFunctions() {
		contractFn := shared.Function{
			Name:    standardFunction.Name,
			Inputs:  make([]shared.Input, 0),
			Outputs: make([]shared.Output, 0),
		}

		for _, contractFunction := range contract.Functions {
			if _, found := discoveredFunctions[contractFunction.Name]; !found {
				if tokensFound, found := FunctionMatch(&contractFn, standardFunction, contractFunction); found {
					discoveredFunctions[contractFunction.Name] = true
					contractFn.Matched = true
					foundTokenCount += tokensFound
				}
			}
		}

		if !contractFn.Matched {
			contractFn.Matched = false

			if standardFunction.Inputs == nil {
				standardFunction.Inputs = make([]shared.Input, 0)
			} else {
				contractFn.Inputs = standardFunction.Inputs
			}

			if standardFunction.Outputs == nil {
				standardFunction.Outputs = make([]shared.Output, 0)
			} else {
				contractFn.Outputs = standardFunction.Outputs
			}
		}

		toReturn.Contract.Functions = append(toReturn.Contract.Functions, contractFn)
	}

	for _, event := range standard.GetEvents() {

		eventFn := shared.Event{
			Name:    event.Name,
			Inputs:  make([]shared.Input, 0),
			Outputs: make([]shared.Output, 0),
		}

		for _, contractEvent := range contract.Events {
			if _, found := discoveredEvents[contractEvent.Name]; !found {
				if tokensFound, found := EventMatch(&eventFn, event, contractEvent); found {
					discoveredEvents[contractEvent.Name] = true
					eventFn.Matched = true
					foundTokenCount += tokensFound
				}
			}
		}

		if !eventFn.Matched {
			eventFn.Matched = false

			if event.Inputs == nil {
				event.Inputs = make([]shared.Input, 0)
			} else {
				eventFn.Inputs = event.Inputs
			}

			if event.Outputs == nil {
				event.Outputs = make([]shared.Output, 0)
			} else {
				eventFn.Outputs = event.Outputs
			}
		}

		toReturn.Contract.Events = append(toReturn.Contract.Events, eventFn)
	}

	toReturn.DiscoveredTokens = foundTokenCount

	// Calculate the total confidence based on the discovered tokens and maximum tokens
	confidencePoints := float64(foundTokenCount) / float64(standard.TokenCount())
	level, threshold := CalculateDiscoveryConfidence(confidencePoints)
	toReturn.Confidence = level
	toReturn.ConfidencePoints = confidencePoints
	toReturn.Threshold = threshold

	return toReturn, foundTokenCount > 0
}

// FunctionConfidenceCheck checks for function confidence against provided EIP standard
func FunctionConfidenceCheck(standard shared.EIP, fn *shared.Function) (shared.FunctionDiscovery, bool) {
	foundTokenCount := 0
	maximumTokens := standard.FunctionTokenCount(fn.Name)

	toReturn := shared.FunctionDiscovery{
		Standard:         standard.GetType(),
		Confidence:       shared.NoConfidence,
		ConfidencePoints: 0,
		Threshold:        shared.NoConfidenceThreshold,
		MaximumTokens:    maximumTokens,
		DiscoveredTokens: 0,
		Function: &shared.Function{
			Name: fn.Name,
		},
	}

	for _, standardFunction := range standard.GetFunctions() {
		if fn.Name == standardFunction.Name {
			if tokensFound, found := FunctionMatch(toReturn.Function, standardFunction, *fn); found {
				fn.Matched = true
				toReturn.Function.Matched = true
				foundTokenCount += tokensFound
			}
		}
	}

	toReturn.DiscoveredTokens = foundTokenCount
	confidencePoints := float64(foundTokenCount) / float64(maximumTokens)
	level, threshold := CalculateDiscoveryConfidence(confidencePoints)
	toReturn.Confidence = level
	toReturn.ConfidencePoints = confidencePoints
	toReturn.Threshold = threshold

	return toReturn, foundTokenCount > 0
}

// FunctionMatch matches a function from a contract to a standard function and returns the total token count and a boolean indicating if a match was found.
func FunctionMatch(newFn *shared.Function, standardFunction, contractFunction shared.Function) (int, bool) {
	totalTokenCount := 0
	newFn.Name = contractFunction.Name
	if standardFunction.Name == contractFunction.Name {
		totalTokenCount++
		for _, sfnInput := range standardFunction.Inputs {
			newInput := shared.Input{Type: sfnInput.Type, Indexed: sfnInput.Indexed}
			for _, fnInput := range contractFunction.Inputs {
				if standardInput, matched := InputMatch(standardFunction.Inputs, fnInput); matched {
					totalTokenCount += 2 // Counting the input match and type match...
					if standardInput.Indexed == fnInput.Indexed {
						totalTokenCount++
					}
					newInput.Matched = true
					break
				}
			}
			newFn.Inputs = append(newFn.Inputs, newInput)
		}

		for _, sfnOutput := range standardFunction.Outputs {
			newOutput := shared.Output{Type: sfnOutput.Type}
			for range standardFunction.Outputs {
				for _, fnOutput := range contractFunction.Outputs {
					if _, matched := OutputMatch(standardFunction.Outputs, fnOutput); matched {
						totalTokenCount += 2 // Counting the input match and type match...
					}
					newOutput.Matched = true
					break
				}
			}
			newFn.Outputs = append(newFn.Outputs, newOutput)
		}
	}

	return totalTokenCount, totalTokenCount > 0
}

// EventMatch matches an event from a contract to a standard event and returns the total token count and a boolean indicating if a match was found.
func EventMatch(newEvent *shared.Event, standardEvent, event shared.Event) (int, bool) {
	totalTokenCount := 0

	if standardEvent.Name == event.Name {
		totalTokenCount++
		newEvent.Name = event.Name
		for _, seInput := range standardEvent.Inputs {
			newInput := shared.Input{Type: seInput.Type, Indexed: seInput.Indexed}
			for _, eventInput := range event.Inputs {
				if standardInput, matched := InputMatch(standardEvent.Inputs, eventInput); matched {
					totalTokenCount += 2 // Counting the input match and type match...
					if standardInput.Indexed == eventInput.Indexed {
						totalTokenCount++
					}
					newInput.Matched = true
					break
				}
			}
			newEvent.Inputs = append(newEvent.Inputs, newInput)
		}

		for _, seOutput := range standardEvent.Outputs {
			newOutput := shared.Output{Type: seOutput.Type}
			for range event.Outputs {
				for _, fnOutput := range event.Outputs {
					if _, matched := OutputMatch(standardEvent.Outputs, fnOutput); matched {
						totalTokenCount += 2 // Counting the input match and type match...
					}
					newOutput.Matched = true
					break
				}
			}
			newEvent.Outputs = append(newEvent.Outputs, newOutput)
		}
	}

	return totalTokenCount, totalTokenCount > 0
}

// InputMatch matches an input to a list of inputs and returns the matched input and a boolean indicating if a match was found.
func InputMatch(inputs []shared.Input, nodeInput shared.Input) (*shared.Input, bool) {
	for _, input := range inputs {
		if input.Type == nodeInput.Type {
			return &input, true
		}
	}

	return nil, false
}

// OutputMatch matches an output to a list of outputs and returns the matched output and a boolean indicating if a match was found.
func OutputMatch(outputs []shared.Output, nodeOutput shared.Output) (*shared.Output, bool) {
	for _, output := range outputs {
		if output.Type == nodeOutput.Type {
			return &output, true
		}
	}

	return nil, false
}
