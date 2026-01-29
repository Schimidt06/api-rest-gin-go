# API Rest Gin Go

Este projeto é uma API REST desenvolvida em Go usando o framework Gin e GORM para operações de banco de dados com PostgreSQL (ou SQLite, dependendo da configuração). A API gerencia um cadastro de alunos, permitindo operações de CRUD (Create, Read, Update, Delete).

## 🚀 Tecnologias Utilizadas

*   **Go**: Linguagem de programação.
*   **Gin**: Web framework para criar APIs de alta performance.
*   **GORM**: Biblioteca ORM para Go.
*   **PostgreSQL**: Banco de dados relacional.
*   **Docker** (Opcional): Para containerização do banco de dados/aplicação.

## 📋 Pré-requisitos

*   Go instalado (versão 1.16+)
*   Banco de Dados (PostgreSQL configurado ou ajuste para SQLite)

## 🔧 Instalação e Execução

1.  Clone o repositório:
    ```bash
    git clone https://github.com/SEU_USUARIO/api-rest-gin-go.git
    cd api-rest-gin-go
    ```

2.  Instale as dependências:
    ```bash
    go mod tidy
    ```

3.  Execute a aplicação:
    ```bash
    go run main.go
    ```

A API estará rodando em `http://localhost:8080` (porta padrão do Gin).

## 📍 Endpoints da API

Abaixo estão as rotas disponíveis na aplicação:

### Alunos

| Método | Rota | Descrição |
| :--- | :--- | :--- |
| `GET` | `/alunos` | Retorna todos os alunos cadastrados. |
| `GET` | `/alunos/:id` | Busca um aluno pelo ID. |
| `GET` | `/alunos/cpf/:cpf` | Busca um aluno pelo CPF. |
| `POST` | `/alunos` | Cria um novo aluno. |
| `DELETE` | `/alunos/:id` | Deleta um aluno pelo ID. |
| `PATCH` | `/alunos/:id` | Atualiza os dados de um aluno. |

### Utilitários

| Método | Rota | Descrição |
| :--- | :--- | :--- |
| `GET` | `/:nome` | Retorna uma saudação personalizada (Saudacao). |

## 📦 Estrutura do Modelo (Aluno)

```json
{
  "nome": "string",
  "cpf": "string",
  "rg": "string"
}
```

## 🤝 Contribuindo

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou enviar pull requests.
