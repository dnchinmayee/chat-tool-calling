# Chat Tool API

This project provides a simple REST API for a chat application using the Gin framework in Go.

## Features

- **GET /api/chat:**  This endpoint handles chat requests.  (Further details on request parameters and response structure would need to be added based on the implementation in `controllers/chatController.go` and `models/chat.go`)

## Getting Started

1. **Clone the repository:**
   ```bash
   git clone [repository_url]
   ```
2. **Install dependencies:**
   ```bash
   go mod tidy
   ```
3. **Run the server:**
   ```bash
   go run main.go
   ```
The API will be available at `http://localhost:8080`.

## API Documentation

## API Documentation

### Request Body (`/api/chat`)

The API accepts a JSON payload with the following structure:

```json
{
  "model": "gpt-3.5-turbo", // Model to use
  "messages": [
    {
      "role": "system",
      "content": "You are a helpful assistant."
    },
    {
      "role": "user",
      "content": "What is the capital of France?"
    }
  ],
  "tools": [
    {
      "type": "function",
      "function": {
        "name": "calculate_revenue",
        "description": "Calculates revenue for a given year and month.",
        "parameters": {
          "type": "object",
          "properties": {
            "year": { "type": "integer", "description": "The year." },
            "month": { "type": "integer", "description": "The month (1-12)." }
          },
          "required": ["year", "month"]
        }
      }
    },
    {
      "type": "function",
      "function": {
        "name": "get_current_weather",
        "description": "Gets the current weather for a given location.",
        "parameters": {
          "type": "object",
          "properties": {
            "location": { "type": "string", "description": "The location (e.g., London)." },
            "format": { "type": "string", "description": "The format (e.g., '3')." }
          },
          "required": ["location", "format"]
        }
      }
    },
    {
      "type": "function",
      "function": {
        "name": "multiply_two_numbers",
        "description": "Multiplies two numbers.",
        "parameters": {
          "type": "object",
          "properties": {
            "number1": { "type": "number", "description": "The first number." },
            "number2": { "type": "number", "description": "The second number." }
          },
          "required": ["number1", "number2"]
        }
      }
    }
  ]
}
```

### Response Body

The API returns a JSON payload with the following structure:

```json
{
  "id": "string",
  "object": "chat.completion",
  "created": 1678886000,
  "model": "gpt-3.5-turbo-0301",
  "choices": [
    {
      "index": 0,
      "finish_reason": "stop",
      "message": {
        "role": "assistant",
        "content": "Paris"
      }
    }
  ],
  "usage": {
    "prompt_tokens": 12,
    "completion_tokens": 1,
    "total_tokens": 13
  }
}
```

### Available Functions

- `calculate_revenue(year: integer, month: integer)`: Calculates revenue.
- `get_current_weather(location: string, format: string)`: Gets current weather.
- `multiply_two_numbers(number1: number, number2: number)`: Multiplies two numbers.


## Contributing

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

MIT License

Copyright (c) 2024

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
