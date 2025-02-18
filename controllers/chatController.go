package controllers

import (
	"bytes"
	"chat-tool-calling/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"io"

	"github.com/gin-gonic/gin"
)

type ChatController struct {
}

func NewChatController() *ChatController {
	return &ChatController{}
}

func (ctrl *ChatController) GetChat(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
	apiKey := c.Request.Header.Get("Authorization")
	// remove Bearer
	apiKey = apiKey[7:]

	chatBody := models.ChatBody{}
	if err := c.ShouldBindJSON(&chatBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// chatResponse := []models.Message{
	// 	models.Message{
	// 		Role:    chatBody.Messages[0].Role,
	// 		Content: chatBody.Messages[0].Content,
	// 	},
	// }
	chatResponse := models.UserResponse{
		Messages: chatBody.Messages,
	}

	// add tools to chatBody
	chatBody.Tools = models.GetTools()

	log.Println(apiKey)
	log.Println(chatBody)

	responseBody, err := ctrl.CallLLM(chatBody, apiKey)
	chatResponse.ChatResponse = responseBody

	// extract function calls
	funcResponse := extractFunctionCall(responseBody.Choices[0].Message.ToolCalls)

	if funcResponse != "" {
		chatResponse.Messages = append(chatResponse.Messages, models.Message{
			Role:    "function",
			Content: funcResponse,
		})
	} else {
		chatResponse.Messages = append(chatResponse.Messages, models.Message{
			Role:    "assistant",
			Content: responseBody.Choices[0].Message.Content,
		})
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": chatResponse})

}

// make a request to POST https://openrouter.ai/api/v1/chat/completions
func (ctrl *ChatController) CallLLM(body models.ChatBody, apiKey string) (models.ChatResponse, error) {
	appUrl := "https://openrouter.ai/api/v1/chat/completions"
	responseBody := models.ChatResponse{}

	// call LLM
	client := &http.Client{}

	// add tools to chatBody
	// body.Tools = models.GetTools()

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		log.Printf("error marshalling json: %v", err)
		return responseBody, err
	}

	log.Println(string(jsonBytes))

	req, err := http.NewRequest("POST", appUrl, bytes.NewBuffer(jsonBytes))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+apiKey)
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("error calling LLM: %v", err)
		return responseBody, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		log.Printf("error decoding response: %v", err)
		return responseBody, err
	}

	log.Printf("response body: %+v", responseBody)

	return responseBody, nil
}

func extractFunctionCall(toolCalls []models.ToolCall) string {
	response := ""

	for _, toolCall := range toolCalls {
		functionName := toolCall.Function.Name
		arguments := toolCall.Function.Arguments

		log.Printf("Function name: %s, Arguments: %s", functionName, arguments)

		// Remove the outer double quotes from the arguments string
		arguments = strings.Trim(arguments, "\"")

		// Unmarshal the arguments JSON string into a map
		var argsMap map[string]interface{}
		err := json.Unmarshal([]byte(arguments), &argsMap)
		if err != nil {
			log.Printf("error unmarshalling arguments: %v", err)
			return ""
		}

		// Call the relevant function with parameters
		switch functionName {
		case "calculate_revenue":
			resp := calculateRevenue(int(argsMap["year"].(float64)), int(argsMap["month"].(float64)))
			if resp != "" {
				response += resp
			}
		case "get_current_weather":
			resp := getCurrentWeather(argsMap["location"].(string), argsMap["format"].(string))
			if resp != "" {
				response += resp
			}
		case "multiply_two_numbers":
			resp := multiplyTwoNumbers(argsMap["number1"].(float64), argsMap["number2"].(float64))
			if resp != "" {
				response += resp
			}
		default:
			log.Println("Unknown function name")
		}
	}

	return response
}

func calculateRevenue(year int, month int) string {
	// Implement the calculateRevenue function
	log.Println("Calculating revenue for year", year, "and month", month)
	return fmt.Sprintf("Revenue for year %d and month %d is $1000", year, month)
}

func getCurrentWeather(location string, format string) string {
	// Implement the getCurrentWeather function
	log.Println("Getting current weather for location", location, "in format", format)

	// wttr.in/London?format=3
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://wttr.in/"+location+"?format=%f", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var weather string
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error reading response: %v", err)
		return ""
	}
	weather = string(body)

	// return fmt.Sprintf("Current weather for location %s in format %s is %s", location, format, weather)
	return fmt.Sprintf("Currently it feels like %s in %s", weather, location)
}

func multiplyTwoNumbers(number1 float64, number2 float64) string {
	// Implement the multiplyTwoNumbers function
	log.Println("Multiplying", number1, "and", number2)
	return fmt.Sprintf("%f * %f = %f", number1, number2, number1*number2)
}
