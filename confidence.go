package standards

import "github.com/unpackdev/standards/utils"

// CalculateDiscoveryConfidence calculates the confidence level and threshold based on the total confidence.
func CalculateDiscoveryConfidence(totalConfidence float64) (utils.ConfidenceLevel, utils.ConfidenceThreshold) {
	total := utils.ConfidenceThreshold(totalConfidence)
	switch {
	case total == utils.PerfectConfidenceThreshold:
		return utils.PerfectConfidence, utils.PerfectConfidenceThreshold
	case total >= utils.HighConfidenceThreshold:
		return utils.HighConfidence, utils.HighConfidenceThreshold
	case total >= utils.MediumConfidenceThreshold:
		return utils.MediumConfidence, utils.MediumConfidenceThreshold
	case total >= utils.LowConfidenceThreshold:
		return utils.LowConfidence, utils.LowConfidenceThreshold
	default:
		return utils.NoConfidence, utils.NoConfidenceThreshold
	}
}

// ConfidenceCheck checks the confidence of a contract against a standard EIP.
func ConfidenceCheck(standard EIP, contract *utils.ContractMatcher) (utils.Discovery, bool) {
	toReturn := utils.Discovery{
		Standard:         standard.GetType(),
		Confidence:       utils.NoConfidence,
		ConfidencePoints: 0,
		Threshold:        utils.NoConfidenceThreshold,
		MaximumTokens:    standard.TokenCount(),
		DiscoveredTokens: 0,
		Contract: &utils.ContractMatcher{
			Name:      contract.Name,
			Functions: make([]utils.Function, 0),
			Events:    make([]utils.Event, 0),
		},
	}
	foundTokenCount := 0
	discoveredFunctions := map[string]bool{}
	discoveredEvents := map[string]bool{}

	for _, standardFunction := range standard.GetFunctions() {
		contractFn := utils.Function{
			Name:    standardFunction.Name,
			Inputs:  make([]utils.Input, 0),
			Outputs: make([]utils.Output, 0),
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
				standardFunction.Inputs = make([]utils.Input, 0)
			} else {
				contractFn.Inputs = standardFunction.Inputs
			}

			if standardFunction.Outputs == nil {
				standardFunction.Outputs = make([]utils.Output, 0)
			} else {
				contractFn.Outputs = standardFunction.Outputs
			}
		}

		toReturn.Contract.Functions = append(toReturn.Contract.Functions, contractFn)
	}

	for _, event := range standard.GetEvents() {

		eventFn := utils.Event{
			Name:    event.Name,
			Inputs:  make([]utils.Input, 0),
			Outputs: make([]utils.Output, 0),
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
				event.Inputs = make([]utils.Input, 0)
			} else {
				eventFn.Inputs = event.Inputs
			}

			if event.Outputs == nil {
				event.Outputs = make([]utils.Output, 0)
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
func FunctionConfidenceCheck(standard EIP, fn *utils.Function) (utils.FunctionDiscovery, bool) {
	foundTokenCount := 0
	maximumTokens := standard.FunctionTokenCount(fn.Name)

	toReturn := utils.FunctionDiscovery{
		Standard:         standard.GetType(),
		Confidence:       utils.NoConfidence,
		ConfidencePoints: 0,
		Threshold:        utils.NoConfidenceThreshold,
		MaximumTokens:    maximumTokens,
		DiscoveredTokens: 0,
		Function: &utils.Function{
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
func FunctionMatch(newFn *utils.Function, standardFunction, contractFunction utils.Function) (int, bool) {
	totalTokenCount := 0
	newFn.Name = contractFunction.Name
	if standardFunction.Name == contractFunction.Name {
		totalTokenCount++
		for _, sfnInput := range standardFunction.Inputs {
			newInput := utils.Input{Type: sfnInput.Type, Indexed: sfnInput.Indexed}
			for _, fnInput := range contractFunction.Inputs {
				if standardInput, matched := inputMatch(standardFunction.Inputs, fnInput); matched {
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
			newOutput := utils.Output{Type: sfnOutput.Type}
			for range standardFunction.Outputs {
				for _, fnOutput := range contractFunction.Outputs {
					if _, matched := outputMatch(standardFunction.Outputs, fnOutput); matched {
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
func EventMatch(newEvent *utils.Event, standardEvent, event utils.Event) (int, bool) {
	totalTokenCount := 0

	if standardEvent.Name == event.Name {
		totalTokenCount++
		newEvent.Name = event.Name
		for _, seInput := range standardEvent.Inputs {
			newInput := utils.Input{Type: seInput.Type, Indexed: seInput.Indexed}
			for _, eventInput := range event.Inputs {
				if standardInput, matched := inputMatch(standardEvent.Inputs, eventInput); matched {
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
			newOutput := utils.Output{Type: seOutput.Type}
			for range event.Outputs {
				for _, fnOutput := range event.Outputs {
					if _, matched := outputMatch(standardEvent.Outputs, fnOutput); matched {
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

// inputMatch matches an input to a list of inputs and returns the matched input and a boolean indicating if a match was found.
func inputMatch(inputs []utils.Input, nodeInput utils.Input) (*utils.Input, bool) {
	for _, input := range inputs {
		if input.Type == nodeInput.Type {
			return &input, true
		}
	}

	return nil, false
}

// outputMatch matches an output to a list of outputs and returns the matched output and a boolean indicating if a match was found.
func outputMatch(outputs []utils.Output, nodeOutput utils.Output) (*utils.Output, bool) {
	for _, output := range outputs {
		if output.Type == nodeOutput.Type {
			return &output, true
		}
	}

	return nil, false
}
