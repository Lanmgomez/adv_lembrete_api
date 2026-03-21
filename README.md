# 🚀 ADV Lembrete API

API REST construída em Go utilizando Gin, MySQL e JWT para autenticação.

---

## 📚 Índice

* [▶️ Como rodar](#️-como-rodar-o-projeto)
* [🔐 Autenticação](#-1-autenticação)
* [👥 Usuários](#-2-usuários)
* [🏢 Entidades](#-3-entidades)
* [⏰ Lembretes](#-4-lembretes)
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
go get github.com/robfig/cron/v3
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
Content-Type: application/json
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
```json
{
  "username": "teste 5",
  "email": "teste_5@hotmail.com",
  "password": "123456",
  "user_type": "user"
}
```

---

## ✏️ Atualizar usuário

**PUT** `/api/users/:id`

---

## ❌ Deletar usuário

**DELETE** `/api/users/:id`

---

# 🏢 3. Entidades

## 📄 Listar entidades (paginado)

**GET** `/api/entidades`

### Headers

```http
Content-Type: application/json
Authorization: Bearer SEU_TOKEN
```

### Query Params

* `nome` → string
* `page` → int
* `limit` → int

---

## 🔍 Buscar entidade por ID

**GET** `/api/entidades/:id`

---

## ➕ Criar entidade

**POST** `/api/entidades`

```json
{
  "nome_entidade": "teste teste"
}
```

---

## ✏️ Atualizar entidade

**PUT** `/api/entidades/:id`

---

## ❌ Deletar entidade

**DELETE** `/api/entidades/:id`

---

# ⏰ 4. Lembretes

## 📄 Listar lembretes (paginado)

**GET** `/api/lembretes`

### Headers

```http
Content-Type: application/json
Authorization: Bearer SEU_TOKEN
```

### Query Params (opcionais)

* `nome` → string
* `page` → int
* `limit` → int

### Resposta (200 OK)

```json
{
  "current_page": 1,
  "total": 1,
  "data": [
    {
      "id": 1,
      "entidade_id": 1,
      "nome_lembrete": "Enviar relatório",
      "descricao": "Relatório mensal da entidade",
      "status": "pendente",
      "data_vencimento": "2026-04-10T00:00:00Z",
      "dias_antecedencia": 10,
      "email_notificacao": "email@dominio.com",
      "created_at": "2026-03-21T10:00:00Z",
      "updated_at": "2026-03-21T10:00:00Z",
      "dias_restantes": "20 dias restantes"
    }
  ]
}
```

## 🔍 Buscar lembrete por ID

**GET** `/api/lembrete/:id`

### Headers

```http
Content-Type: application/json
Authorization: Bearer SEU_TOKEN
```

### Resposta (200 OK)

```json
{
  "data": {
    "id": 1,
    "entidade_id": 1,
    "nome_lembrete": "Enviar relatório",
    "descricao": "Relatório mensal da entidade",
    "status": "pendente",
    "data_vencimento": "2026-04-10T00:00:00Z",
    "dias_antecedencia": 10,
    "email_notificacao": "email@dominio.com",
    "created_at": "2026-03-21T10:00:00Z",
    "updated_at": "2026-03-21T10:00:00Z",
    "dias_restantes": "20 dias restantes"
  }
}
```

---

## ➕ Criar lembrete

### Headers

```http
Content-Type: application/json
Authorization: Bearer SEU_TOKEN
```

**POST** `/api/lembrete`

```json
{
  "entidade_id": 1,
  "nome_lembrete": "Enviar relatório",
  "descricao": "Relatório mensal da entidade",
  "data_vencimento": "2026-04-10",
  "dias_antecedencia": 10,
  "email_notificacao": "email@dominio.com"
}
```

## Resposta (201 Created)
```json
{
  "message": "lembrete criado com sucesso"
}
```

## Regras
* O lembrete só pode ser criado se a entidade existir
* O sistema define o status inicial como pendente

---

## ✏️ Atualizar lembrete

**PUT** `/api/lembrete/:id`

### Headers

```http
Content-Type: application/json
Authorization: Bearer SEU_TOKEN
```

```json
{
  "entidade_id": 1,
  "nome_lembrete": "Teste 01",
  "descricao": "Testando crud...",
  "status": "concluido",
  "data_vencimento": "2026-04-22",
  "dias_antecedencia": 1,
  "email_notificacao": "islan_gomes@hotmail.com"
}
```

## Resposta (200 OK)
```json
{
  "message": "lembrete atualizado com sucesso"
}
```

---

## ❌ Deletar entidade

**DELETE** `/api/entidades/:id`

### Headers

```http
Content-Type: application/json
Authorization: Bearer SEU_TOKEN
```

## Resposta (200 OK)
```json
{
  "message": "lembrete deletado com sucesso"
}
```

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
