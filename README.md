# OpenAI API Mock Server

A simple mock server that mimics the basic functionality of OpenAI's API endpoints. This project is designed for development and testing purposes.

## Available Endpoints

- `/v1/chat/completions` - Mock chat completions (GPT-3.5/4)
- `/v1/completions` - Mock legacy completions
- `/v1/embeddings` - Mock text embeddings
- `/v1/moderations` - Mock content moderation

## Getting Started

### Prerequisites
- Go 1.x or higher

### Running the Server

1. Clone the repository
2. Navigate to the project directory
3. Run the server:

```bash
go run main.go
```

### API Usage
All endpoints accept POST requests and return JSON responses. Here are examples of how to use each endpoint:

#### Chat Completions

```bash
curl -X POST http://localhost:8080/v1/chat/completions \
  -H "Content-Type: application/json" \
  -d '{"model": "gpt-3.5-turbo", "messages": [{"role": "user", "content": "Hello, how are you?"}]}'
```

#### Legacy Completions

```bash
curl -X POST http://localhost:8080/v1/completions \
  -H "Content-Type: application/json" \
  -d '{"model": "text-davinci-003", "prompt": "Hello, how are you?"}'
```

#### Embeddings

```bash
curl -X POST http://localhost:8080/v1/embeddings \
  -H "Content-Type: application/json" \
  -d '{"model": "text-embedding-ada-002", "input": "Hello, how are you?"}'
```

#### Moderations

```bash
curl -X POST http://localhost:8080/v1/moderations \
  -H "Content-Type: application/json" \
  -d '{"model": "text-moderation-latest", "input": "Hello, how are you?"}'
```

# Thanks for using my mock server!

