package models

// define the tools
func GetWeatherTool() Tool {
	return Tool{
		Type: "function",
		Function: &Function{
			Name:        "get_current_weather",
			Description: "Get the current weather for a location",
			Parameters: &Parameters{
				Type: "object",
				Properties: map[string]*Parameter{
					"location": {
						Type:        "string",
						Description: "The location to get the weather for, e.g. San Francisco, CA",
					},
					"format": {
						Type:        "string",
						Description: "The format to return the weather in, e.g. 'celsius' or 'fahrenheit'",
						Enum:        []string{"celsius", "fahrenheit"},
					},
				},
				Required: []string{"location", "format"},
			},
		},
	}
}

func GetRevenueTool() Tool {
	return Tool{
		Type: "function",
		Function: &Function{
			Name:        "calculate_revenue",
			Description: "Calculate revenue for MagicVolo for a given month",
			Parameters: &Parameters{
				Type: "object",
				Properties: map[string]*Parameter{
					"month": {
						Type:        "integer",
						Description: "This is the integer representation of the month for which revenue is to be calculated. If quarter is provided, the three months of that quarter will be used.",
					},
					"year": {
						Type:        "integer",
						Description: "This is the integer representation of the year for which revenue is to be calculated.",
					},
				},
				Required: []string{"month"},
			},
		},
	}
}

func GetMultiplyTool() Tool {
	return Tool{
		Type: "function",
		Function: &Function{
			Name:        "multiply_two_numbers",
			Description: "Calculator to multiply two numbers and return the result",
			Parameters: &Parameters{
				Type: "object",
				Properties: map[string]*Parameter{
					"number1": {
						Type:        "float",
						Description: "This is the first number to multiply",
					},
					"number2": {
						Type:        "float",
						Description: "This is the second number to multiply",
					},
				},
				Required: []string{"number1", "number2"},
			},
		},
	}
}

func GetTools() []Tool {
	return []Tool{
		GetWeatherTool(),
		GetRevenueTool(),
		GetMultiplyTool(),
	}
}
