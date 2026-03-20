# 🚀 ADV Lembrete API

API REST construída em Go utilizando Gin, MySQL e JWT para autenticação.

---

## 📚 Índice

* [▶️ Como rodar](#️-como-rodar-o-projeto)
* [🔐 Autenticação](#-1-autenticação)
* [👥 Usuários](#-2-usuários)
* [🏢 Entidades](#-3-entidades)
* [🛠️ Tecnologias](#️-tecnologias-utilizadas)

---

## ▶️ Como rodar o projeto

```bash
go run cmd/main.go
```

---

## 📦 Dependências

```bash
go get github.com/gin-gonic/gin
go get github.com/go-sql-driver/mysql
go get golang.org/x/crypto/bcrypt
go get github.com/joho/godotenv
go get github.com/golang-jwt/jwt/v5
go get github.com/gin-contrib/cors
```

---

## 📦 Subir com Docker

```bash
docker-compose up --build
```

---

# 🔐 1. Autenticação

## 🔑 Login

**POST** `/api/auth/login`

```json
{
  "email": "islan_gomes@hotmail.com",
  "password": "123456"
}
```

---

## 🚪 Logout

**POST** `/api/logout`

```http
Authorization: Bearer SEU_TOKEN
```

---

# 👥 2. Usuários

## 📄 Listar usuários

**GET** `/api/users`

---

## 🔍 Buscar por ID

**GET** `/api/users/:id`

---

## ➕ Criar usuário

**POST** `/api/users`

---

## ✏️ Atualizar usuário

**PUT** `/api/users/:id`

---

## ❌ Deletar usuário

**DELETE** `/api/users/:id`

---

# 🏢 3. Entidades

## 📄 Listar entidades (paginado)

**GET** `/api/lembretes`

### Headers

```http
Authorization: Bearer SEU_TOKEN
```

### Query Params

* `nome` → string
* `page` → int
* `limit` → int

---

## 🔍 Buscar entidade por ID

**GET** `/api/lembretes/:id`

---

## ➕ Criar entidade

**POST** `/api/lembretes`

```json
{
  "nome_entidade": "teste teste"
}
```

---

## ✏️ Atualizar entidade

**PUT** `/api/lembretes/:id`

---

## ❌ Deletar entidade

**DELETE** `/api/lembretes/:id`

---

# 🛠️ Tecnologias utilizadas

* Go (Golang)
* Gin Gonic
* MySQL
* JWT
* Docker

---

# 📌 Observações

* Todas as rotas (exceto login) requerem autenticação
* Header obrigatório:

```http
Authorization: Bearer SEU_TOKEN
```
