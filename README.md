# RAG Go Project

This project implements a Retrieval-Augmented Generation (RAG) concept using OpenAI and an open-source vector store. It is built with Go and follows REST API principles using the Chi router.

## Project Structure

```
go-rag
├── cmd
│   └── main.go          # Entry point of the application
├── internal
│   ├── handlers
│   │   └── handler.go   # HTTP request handlers
│   ├── routes
│   │   └── routes.go    # Route definitions
│   ├── services
│   │   └── openai.go    # OpenAI API interaction
│   └── vectorstore
│       └── vectorstore.go # Vector store implementation
├── pkg
│   └── models
│       └── model.go     # Data models
├── Dockerfile            # Docker image instructions
├── docker-compose.yml    # Docker Compose configuration
├── go.mod                # Module definition
├── go.sum                # Dependency checksums
└── README.md             # Project documentation
```

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd go-rag
   ```

2. **Build the Docker image:**
   ```
   docker build -t go-rag .
   ```

3. **Run the application using Docker Compose:**
   ```
   docker-compose up
   ```

## Usage

Once the application is running, you can interact with the API endpoints defined in the routes. The handlers will process requests related to the RAG concept, utilizing the OpenAI API and the vector store for data retrieval and generation.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.# go-rag
