package standards

import (
	"github.com/unpackdev/standards/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEIPConfidenceDiscovery(t *testing.T) {
	tests := []struct {
		name       string
		standard   EIP
		outputPath string
		contracts  []struct {
			name                 string
			outputFile           string
			contract             *utils.ContractMatcher
			expectedLevel        utils.ConfidenceLevel
			expectedThreshold    utils.ConfidenceThreshold
			standardTokenCount   int
			discoveredTokenCount int
			shouldMatch          bool
			expectedEip          string
			expectedProto        string
		}
		expectedError string
	}{
		{
			name: "Test ERC20",
			standard: func() EIP {
				standard, err := GetContractByStandard(ERC20)
				assert.NoError(t, err)
				assert.NotNil(t, standard)
				return standard
			}(),
			outputPath: "eip/",
			contracts: []struct {
				name                 string
				outputFile           string
				contract             *utils.ContractMatcher
				expectedLevel        utils.ConfidenceLevel
				expectedThreshold    utils.ConfidenceThreshold
				standardTokenCount   int
				discoveredTokenCount int
				shouldMatch          bool
				expectedEip          string
				expectedProto        string
			}{
				/*	{
								name:       "Full Match",
								outputFile: "eip20_full_match",
								contract: &ContractMatcher{
									Name: "ERC20 Full Match",
									Functions: []Function{
										newFunction("totalSupply", nil, []Output{{Type: TypeUint256}}),
										newFunction("balanceOf", []Input{{Type: TypeAddress}}, []Output{{Type: TypeUint256}}),
										newFunction("transfer", []Input{{Type: TypeAddress}, {Type: TypeUint256}}, []Output{{Type: TypeBool}}),
										newFunction("transferFrom", []Input{{Type: TypeAddress}, {Type: TypeAddress}, {Type: TypeUint256}}, []Output{{Type: TypeBool}}),
										newFunction("approve", []Input{{Type: TypeAddress}, {Type: TypeUint256}}, []Output{{Type: TypeBool}}),
										newFunction("allowance", []Input{{Type: TypeAddress}, {Type: TypeAddress}}, []Output{{Type: TypeUint256}}),
									},
									Events: []Event{
										newEvent("Transfer", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}, {Type: TypeUint256}}, nil),
										newEvent("Approval", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}, {Type: TypeUint256}}, nil),
									},
								},
								expectedLevel:        PerfectConfidence,
								expectedThreshold:    PerfectConfidenceThreshold,
								standardTokenCount:   68,
								discoveredTokenCount: 68,
								shouldMatch:          true,
								expectedEip:          tests.ReadJsonBytesForTest(t, "eip/eip20_full_match").Content,
								expectedProto:        tests.ReadJsonBytesForTest(t, "eip/eip20_full_match.proto").Content,
							},
							{
								name:       "High Match",
								outputFile: "eip20_high_match",
								contract: &ContractMatcher{
									Name: "ERC20 High Match",
									Functions: []Function{
										newFunction("totalSupply", nil, []Output{{Type: TypeUint256}}),
										newFunction("balanceOf", []Input{{Type: TypeAddress}}, []Output{{Type: TypeUint256}}),
										newFunction("transfer", []Input{{Type: TypeAddress}, {Type: TypeUint256}}, []Output{{Type: TypeBool}}),
										newFunction("transferFrom", []Input{{Type: TypeAddress}, {Type: TypeAddress}, {Type: TypeUint256}}, []Output{{Type: TypeBool}}),
										newFunction("approve", []Input{{Type: TypeAddress}, {Type: TypeUint256}}, []Output{{Type: TypeBool}}),
										newFunction("allowance", []Input{{Type: TypeAddress}, {Type: TypeAddress}}, []Output{}),
									},
									Events: []Event{
										newEvent("Transfer", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}, {Type: TypeUint256}}, nil),
										newEvent("Approval", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}, {Type: TypeUint256}}, nil),
									},
								},
								expectedLevel:        HighConfidence,
								expectedThreshold:    HighConfidenceThreshold,
								standardTokenCount:   68,
								discoveredTokenCount: 66,
								shouldMatch:          true,
								expectedEip:          tests.ReadJsonBytesForTest(t, "eip/eip20_high_match").Content,
								expectedProto:        tests.ReadJsonBytesForTest(t, "eip/eip20_high_match.proto").Content,
							},
							{
								name:       "Medium Match",
								outputFile: "eip20_medium_match",
								contract: &ContractMatcher{
									Name: "ERC20 Medium Match",
									Functions: []Function{
										newFunction("totalSupply", nil, []Output{{Type: TypeUint256}}),
										newFunction("balanceOf", []Input{{Type: TypeAddress}}, []Output{{Type: TypeUint256}}),
										newFunction("transfer", []Input{{Type: TypeAddress}, {Type: TypeUint256}}, []Output{{Type: TypeBool}}),
										newFunction("transferFrom", []Input{{Type: TypeAddress}, {Type: TypeAddress}, {Type: TypeUint256}}, []Output{{Type: TypeBool}}),
										newFunction("approve", []Input{{Type: TypeAddress}, {Type: TypeUint256}}, []Output{{Type: TypeBool}}),
									},
									Events: []Event{
										newEvent("Transfer", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}, {Type: TypeUint256}}, nil),
										newEvent("Approval", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}, {Type: TypeUint256}}, nil),
									},
								},
								expectedLevel:        MediumConfidence,
								expectedThreshold:    MediumConfidenceThreshold,
								standardTokenCount:   68,
								discoveredTokenCount: 59,
								shouldMatch:          true,
								expectedEip:          tests.ReadJsonBytesForTest(t, "eip/eip20_medium_match").Content,
								expectedProto:        tests.ReadJsonBytesForTest(t, "eip/eip20_medium_match.proto").Content,
							},
							{
								name:       "Low Match",
								outputFile: "eip20_low_match",
								contract: &ContractMatcher{
									Name: "ERC20 Low Match",
									Functions: []Function{
										newFunction("totalSupply", nil, []Output{{Type: TypeUint256}}),
									},
									Events: []Event{
										newEvent("Transfer", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}, {Type: TypeUint256}}, nil),
									},
								},
								expectedLevel:        LowConfidence,
								expectedThreshold:    LowConfidenceThreshold,
								standardTokenCount:   68,
								discoveredTokenCount: 13,
								shouldMatch:          true,
								expectedEip:          tests.ReadJsonBytesForTest(t, "eip/eip20_low_match").Content,
								expectedProto:        tests.ReadJsonBytesForTest(t, "eip/eip20_low_match.proto").Content,
							},
							{
								name:       "No Match",
								outputFile: "eip20_no_match",
								contract: &ContractMatcher{
									Name:      "ERC20 No Match",
									Functions: []Function{},
									Events:    []Event{},
								},
								expectedLevel:        NoConfidence,
								expectedThreshold:    NoConfidenceThreshold,
								standardTokenCount:   68,
								discoveredTokenCount: 0,
								shouldMatch:          true,
								expectedEip:          tests.ReadJsonBytesForTest(t, "eip/eip20_no_match").Content,
								expectedProto:        tests.ReadJsonBytesForTest(t, "eip/eip20_no_match.proto").Content,
							},
						},
						expectedError: "",
					},
					{
						name: "Test EIP721",
						standard: func() EIP {
							standard, err := GetContractByStandard(ERC721)
							assert.NoError(t, err)
							assert.NotNil(t, standard)
							return standard
						}(),
						outputPath: "eip/",
						contracts: []struct {
							name                 string
							outputFile           string
							contract             *ContractMatcher
							expectedLevel        ConfidenceLevel
							expectedThreshold    ConfidenceThreshold
							standardTokenCount   int
							discoveredTokenCount int
							shouldMatch          bool
							expectedEip          string
							expectedProto        string
						}{
							{
								name:       "Full Match",
								outputFile: "eip721_full_match",
								contract: &ContractMatcher{
									Name: "ERC721 Full Match",
									Functions: []Function{
										newFunction("name", nil, []Output{{Type: TypeString}}),
										newFunction("symbol", nil, []Output{{Type: TypeString}}),
										newFunction("totalSupply", nil, []Output{{Type: TypeUint256}}),
										newFunction("balanceOf", []Input{{Type: TypeAddress}}, []Output{{Type: TypeUint256}}),
										newFunction("ownerOf", []Input{{Type: TypeUint256}}, []Output{{Type: TypeAddress}}),
										newFunction("transferFrom", []Input{{Type: TypeAddress}, {Type: TypeAddress}, {Type: TypeUint256}}, nil),
										newFunction("approve", []Input{{Type: TypeAddress}, {Type: TypeUint256}}, nil),
										newFunction("setApprovalForAll", []Input{{Type: TypeAddress}, {Type: TypeBool}}, nil),
										newFunction("getApproved", []Input{{Type: TypeUint256}}, []Output{{Type: TypeAddress}}),
										newFunction("isApprovedForAll", []Input{{Type: TypeAddress}, {Type: TypeAddress}}, []Output{{Type: TypeBool}}),
									},
									Events: []Event{
										newEvent("Transfer", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}, {Type: TypeUint256}}, nil),
										newEvent("Approval", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}, {Type: TypeUint256}}, nil),
										newEvent("ApprovalForAll", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}, {Type: TypeBool}}, nil),
									},
								},
								expectedLevel:        PerfectConfidence,
								expectedThreshold:    PerfectConfidenceThreshold,
								standardTokenCount:   90,
								discoveredTokenCount: 90,
								shouldMatch:          true,
								expectedEip:          tests.ReadJsonBytesForTest(t, "eip/eip721_full_match").Content,
								expectedProto:        tests.ReadJsonBytesForTest(t, "eip/eip721_full_match.proto").Content,
							},
						},
					},
					{
						name: "Test EIP1155",
						standard: func() EIP {
							standard, err := GetContractByStandard(ERC1155)
							assert.NoError(t, err)
							assert.NotNil(t, standard)
							return standard
						}(),
						outputPath: "eip/",
						contracts: []struct {
							name                 string
							outputFile           string
							contract             *ContractMatcher
							expectedLevel        ConfidenceLevel
							expectedThreshold    ConfidenceThreshold
							standardTokenCount   int
							discoveredTokenCount int
							shouldMatch          bool
							expectedEip          string
							expectedProto        string
						}{
							{
								name:       "Full Match",
								outputFile: "eip1155_full_match",
								contract: &ContractMatcher{
									Name: "ERC1155 Full Match",
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
								expectedLevel:        PerfectConfidence,
								expectedThreshold:    PerfectConfidenceThreshold,
								standardTokenCount:   115,
								discoveredTokenCount: 115,
								shouldMatch:          true,
								expectedEip:          tests.ReadJsonBytesForTest(t, "eip/eip1155_full_match").Content,
								expectedProto:        tests.ReadJsonBytesForTest(t, "eip/eip1155_full_match.proto").Content,
							},
						},
					},
					{
						name: "Test EIP1820",
						standard: func() EIP {
							standard, err := GetContractByStandard(ERC1820)
							assert.NoError(t, err)
							assert.NotNil(t, standard)
							return standard
						}(),
						outputPath: "eip/",
						contracts: []struct {
							name                 string
							outputFile           string
							contract             *ContractMatcher
							expectedLevel        ConfidenceLevel
							expectedThreshold    ConfidenceThreshold
							standardTokenCount   int
							discoveredTokenCount int
							shouldMatch          bool
							expectedEip          string
							expectedProto        string
						}{
							{
								name:       "Full Match",
								outputFile: "eip1820_full_match",
								contract: &ContractMatcher{
									Name: "ERC1820 Full Match",
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
								expectedLevel:        PerfectConfidence,
								expectedThreshold:    PerfectConfidenceThreshold,
								standardTokenCount:   67,
								discoveredTokenCount: 67,
								shouldMatch:          true,
								expectedEip:          tests.ReadJsonBytesForTest(t, "eip/eip1820_full_match").Content,
								expectedProto:        tests.ReadJsonBytesForTest(t, "eip/eip1820_full_match.proto").Content,
							},
						},
					},
					{
						name: "Test EIP1822",
						standard: func() EIP {
							standard, err := GetContractByStandard(ERC1822)
							assert.NoError(t, err)
							assert.NotNil(t, standard)
							return standard
						}(),
						outputPath: "eip/",
						contracts: []struct {
							name                 string
							outputFile           string
							contract             *ContractMatcher
							expectedLevel        ConfidenceLevel
							expectedThreshold    ConfidenceThreshold
							standardTokenCount   int
							discoveredTokenCount int
							shouldMatch          bool
							expectedEip          string
							expectedProto        string
						}{
							{
								name:       "Full Match",
								outputFile: "eip1822_full_match",
								contract: &ContractMatcher{
									Name: "ERC1822 Full Match",
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
								expectedLevel:        PerfectConfidence,
								expectedThreshold:    PerfectConfidenceThreshold,
								standardTokenCount:   29,
								discoveredTokenCount: 29,
								shouldMatch:          true,
								expectedEip:          tests.ReadJsonBytesForTest(t, "eip/eip1822_full_match").Content,
								expectedProto:        tests.ReadJsonBytesForTest(t, "eip/eip1822_full_match.proto").Content,
							},
						},
					},
					{
						name: "Test EIP1967",
						standard: func() EIP {
							standard, err := GetContractByStandard(ERC1967)
							assert.NoError(t, err)
							assert.NotNil(t, standard)
							return standard
						}(),
						outputPath: "eip/",
						contracts: []struct {
							name                 string
							outputFile           string
							contract             *ContractMatcher
							expectedLevel        ConfidenceLevel
							expectedThreshold    ConfidenceThreshold
							standardTokenCount   int
							discoveredTokenCount int
							shouldMatch          bool
							expectedEip          string
							expectedProto        string
						}{
							{
								name:       "Full Match",
								outputFile: "eip1967_full_match",
								contract: &ContractMatcher{
									Name: "ERC1967 Full Match",
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
										newEvent("AdminChanged", []Input{{Type: utils.TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}}, nil),
									},
								},
								expectedLevel:        utils.PerfectConfidence,
								expectedThreshold:    utils.PerfectConfidenceThreshold,
								standardTokenCount:   67,
								discoveredTokenCount: 67,
								shouldMatch:          true,
								expectedEip:          tests.ReadJsonBytesForTest(t, "eip/eip1967_full_match").Content,
								expectedProto:        tests.ReadJsonBytesForTest(t, "eip/eip1967_full_match.proto").Content,
							},
						},
					},*/
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, contract := range tt.contracts {
				t.Run(contract.name, func(t *testing.T) {
					discovery, found := tt.standard.ConfidenceCheck(contract.contract)

					// Assert the confidence level and threshold against the expected values
					assert.Equal(t, contract.expectedLevel, discovery.Confidence)
					assert.NotEmpty(t, discovery.Confidence.String())
					assert.Equal(t, contract.expectedThreshold, discovery.Threshold)
					assert.Equal(t, contract.standardTokenCount, discovery.MaximumTokens)
					assert.Equal(t, contract.discoveredTokenCount, discovery.DiscoveredTokens)

					// Assert that the function found a match in the contract
					assert.True(t, contract.shouldMatch, found)

					assert.NotNil(t, discovery.ToProto())

					jsonDiscovery, err := utils.ToJSON(discovery)
					assert.NoError(t, err)
					assert.NotNil(t, jsonDiscovery)

					jsonPrettyDiscovery, err := utils.ToJSONPretty(discovery)
					assert.NoError(t, err)
					assert.NotNil(t, jsonPrettyDiscovery)

					protoDiscovery, err := utils.ToProtoJSON(discovery)
					assert.NoError(t, err)
					assert.NotNil(t, protoDiscovery)

					protoPrettyDiscovery, err := utils.ToJSONPretty(discovery.ToProto())
					assert.NoError(t, err)

					// Assert that the JSON output matches the expected output
					assert.Equal(t, contract.expectedEip, string(jsonPrettyDiscovery))
					assert.Equal(t, contract.expectedProto, string(protoPrettyDiscovery))
				})
			}
		})
	}
}

func TestFunctionMatch(t *testing.T) {
	tests := []struct {
		name           string
		newFn          *utils.Function
		standardFn     utils.Function
		contractFn     utils.Function
		expectedTokens int
		expectedMatch  bool
	}{
		{
			name: "Matching function",
			newFn: &utils.Function{
				Name: "transfer",
				Inputs: []utils.Input{
					{Type: utils.TypeAddress, Indexed: false, Matched: false},
					{Type: utils.TypeUint256, Indexed: false, Matched: false},
				},
				Outputs: []utils.Output{{Type: utils.TypeBool, Matched: false}},
				Matched: false,
			},
			standardFn: utils.Function{
				Name: "transfer",
				Inputs: []utils.Input{
					{Type: utils.TypeAddress, Indexed: false, Matched: false},
					{Type: utils.TypeUint256, Indexed: false, Matched: false},
				},
				Outputs: []utils.Output{{Type: utils.TypeBool, Matched: false}},
				Matched: false,
			},
			contractFn: utils.Function{
				Name: "transfer",
				Inputs: []utils.Input{
					{Type: utils.TypeAddress, Indexed: false, Matched: false},
					{Type: utils.TypeUint256, Indexed: false, Matched: false},
				},
				Outputs: []utils.Output{{Type: utils.TypeBool, Matched: false}},
				Matched: false,
			},
			expectedTokens: 9,
			expectedMatch:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTokens, gotMatch := FunctionMatch(tt.newFn, tt.standardFn, tt.contractFn)
			assert.Equal(t, tt.expectedTokens, gotTokens)
			assert.Equal(t, tt.expectedMatch, gotMatch)
		})
	}
}

func TestEventMatch(t *testing.T) {
	tests := []struct {
		name           string
		newEvent       *utils.Event
		standardEvent  utils.Event
		contractEvent  utils.Event
		expectedTokens int
		expectedMatch  bool
	}{
		{
			name: "Matching event",
			newEvent: &utils.Event{
				Name: "Transfer",
				Inputs: []utils.Input{
					{Type: utils.TypeAddress, Indexed: true, Matched: false},
					{Type: utils.TypeAddress, Indexed: true, Matched: false},
					{Type: utils.TypeUint256, Indexed: false, Matched: false},
				},
				Outputs: []utils.Output{},
				Matched: false,
			},
			standardEvent: utils.Event{
				Name: "Transfer",
				Inputs: []utils.Input{
					{Type: utils.TypeAddress, Indexed: true, Matched: false},
					{Type: utils.TypeAddress, Indexed: true, Matched: false},
					{Type: utils.TypeUint256, Indexed: false, Matched: false},
				},
				Outputs: []utils.Output{},
				Matched: false,
			},
			contractEvent: utils.Event{
				Name: "Transfer",
				Inputs: []utils.Input{
					{Type: utils.TypeAddress, Indexed: true, Matched: false},
					{Type: utils.TypeAddress, Indexed: true, Matched: true},
					{Type: utils.TypeUint256, Indexed: false, Matched: false},
				},
				Outputs: []utils.Output{},
				Matched: false,
			},
			expectedTokens: 10,
			expectedMatch:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTokens, gotMatch := EventMatch(tt.newEvent, tt.standardEvent, tt.contractEvent)
			assert.Equal(t, tt.expectedTokens, gotTokens)
			assert.Equal(t, tt.expectedMatch, gotMatch)
		})
	}
}
