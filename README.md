# Projeto de API com Golang, Chi Router, Weaviate, Docker e Air Reload

Este projeto é uma API desenvolvida em Golang que utiliza o Chi Router para roteamento, Weaviate como banco de dados vetorial, Docker para containerização e Air Reload para recarregamento automático durante o desenvolvimento.

## Tecnologias Utilizadas

- **Golang**: Linguagem de programação utilizada para desenvolver a API.
- **Chi Router**: Biblioteca de roteamento para Golang.
- **Weaviate(https://weaviate.io/)**: Banco de dados vetorial utilizado para armazenar e buscar dados.
- **Docker**: Plataforma de containerização utilizada para empacotar a aplicação.
- **Docker Compose**: Ferramenta para definir e gerenciar multi-containers Docker.
- **Air Reload**: Ferramenta para recarregamento automático durante o desenvolvimento.
- **Cohere(https://cohere.com/)**: Utilizado como IA para gerar a resposta final.
- **Transformers Inference**: Utiliza o modelo `sentence-transformers-multi-qa-MiniLM-L6-cos-v1` para embeddings.

## Endpoints

### 1. `/generate`

Endpoint para enviar uma pergunta e obter uma resposta.

- **URL**: `http://localhost:3333/generate`
- **Método**: `POST`
- **Payload**:
  ```json
  {
      "query": "what happened in 1953 ?"
  }
   ```

### 2. `/add-data`

Endpoint para criar dados no banco vetorial.

- **URL**: `http://localhost:3333/add-data`
- **Método**: `POST`

## Como rodar o projeto

Para rodar o projeto, siga os passos abaixo:

1. **Clone o repositório:**
   ```bash
   git clone <URL_DO_REPOSITORIO>
   cd <NOME_DO_DIRETORIO>
   docker-compose up


