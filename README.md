# 🚀 ADV Lembrete API

API REST construída em Go utilizando Gin, MySQL e JWT para autenticação.

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

## 📦 Subir um Docker Container

```bash
 docker-compose up --build
```

---

# 🔐 1. Autenticação

## 🔑 Login

**POST** `/api/auth/login`

### Body

```json
{
  "email": "islan_gomes@hotmail.com",
  "password": "123456"
}
```

### Resposta (200 OK)

```json
{
  "access_token": "SEU_TOKEN",
  "token_type": "Bearer",
  "user": {
    "id": 1,
    "username": "islan gomes",
    "email": "islan_gomes@hotmail.com",
    "user_type": "owner",
    "created_at": "2026-03-18T13:56:18Z"
  }
}
```

---

## 🚪 Logout

**POST** `/api/logout`

### Headers

```http
Authorization: Bearer SEU_TOKEN
```

### Resposta

```json
{
  "message": "logout realizado com sucesso",
  "user_id": 1
}
```

---

# 👥 2. Usuários

## 📄 Listar usuários (paginado)

**GET** `/api/users`

### Headers

```http
Authorization: Bearer SEU_TOKEN
```

### Resposta

```json
{
  "current_page": 1,
  "total": 4,
  "data": [
    {
      "id": 1,
      "username": "islan gomes",
      "email": "islan_gomes@hotmail.com",
      "user_type": "owner",
      "created_at": "2026-03-18T13:56:18Z"
    }
  ]
}
```

---

## 🔍 Buscar usuário por ID

**GET** `/api/users/:id`

### Resposta

```json
{
  "data": {
    "id": 1,
    "username": "islan gomes",
    "email": "islan_gomes@hotmail.com",
    "user_type": "owner",
    "created_at": "2026-03-18T13:56:18Z"
  }
}
```

---

## ➕ Criar usuário

**POST** `/api/users`

### Body

```json
{
  "username": "teste 1",
  "email": "teste_1@hotmail.com",
  "password": "123456",
  "user_type": "user"
}
```

### Resposta (201 Created)

```json
{
  "message": "Usuário criado com sucesso"
}
```

---

## ✏️ Atualizar usuário

**PUT** `/api/users/:id`

### Body

```json
{
  "username": "Islan Gomes",
  "email": "islan_gomes@hotmail.com",
  "password": "123456",
  "user_type": "owner"
}
```

### Resposta

```json
{
  "message": "usuário atualizado com sucesso"
}
```

---

## ❌ Deletar usuário

**DELETE** `/api/users/:id`

### Resposta

```json
{
  "message": "usuário deletado com sucesso"
}
```

---

# 🛠️ Tecnologias utilizadas

* Go (Golang)
* Gin Gonic
* MySQL
* JWT (autenticação)
* Docker (opcional)

---

# 📌 Observações

* Todas as rotas (exceto login) requerem autenticação via JWT
* Header obrigatório:

```http
Authorization: Bearer SEU_TOKEN
```

---
