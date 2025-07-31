```markdown
# 🚀 Go PostgreSQL User API

A simple RESTful API built with Golang to manage users via PostgreSQL. Vibecoding encouraged 😎

---

## 🔧 Prerequisites

- ✅ Go installed → [https://golang.org/dl](https://golang.org/dl)
- 🐘 PostgreSQL running locally or remotely
- 📦 A database table named `users` with:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email TEXT NOT NULL,
    name TEXT NOT NULL,
    age INT
);
```

---

## ⚙️ Configuration

To get started:

1. Open `config/db.go`
2. Replace the connection string with your PostgreSQL credentials:

```go
connStr := "user=yourusername password=yourpassword dbname=yourdb sslmode=disable"
```

> 📍 Note: Database credentials are managed in the `config` folder, not in `router`.

---

## 📁 Project Structure

```bash
go-rest-api/
├── main.go
├── config/
│   └── db.go
├── controller/
│   └── user_controller.go
├── service/
│   └── user_service.go
├── repository/
│   └── user_repository.go
├── model/
│   └── user.go
└── router/
    └── router.go
```

---

## 🏗️ Build & Run

1. 💥 Build the project:
   ```bash
   go build -o api-server
   ```

2. 🚀 Run the API server:
   ```bash
   ./api-server
   ```

> The server will listen on `http://localhost:8080`

---

## 📡 REST API Endpoints

| Method | Endpoint                | Description          |
|--------|-------------------------|----------------------|
| GET    | `/users`                | List all users       |
| POST   | `/users/create`         | Create a new user    |
| PUT    | `/users/update`         | Update user info     |
| DELETE | `/users/delete?id=1`    | Delete user by ID    |

---

## 🧪 Example `curl` Commands

### ➕ Create a user
```bash
curl -X POST -H "Content-Type: application/json" \
-d '{"email":"cat@dev.io","name":"Catarina","age":27}' \
http://localhost:8080/users/create
```

### 🔄 Update a user
```bash
curl -X PUT -H "Content-Type: application/json" \
-d '{"id":1,"email":"newcat@dev.io","name":"Cat","age":28}' \
http://localhost:8080/users/update
```

### ❌ Delete a user
```bash
curl -X DELETE "http://localhost:8080/users/delete?id=1"
```

---  

## 🔐 JWT Authentication

All API endpoints require JWT authentication.

### 1. Set Up Your JWT Private Key

Create a `.env` file in your project root with:

```env
JWT_PRIVATE_KEY=your_super_secret_private_key_here
```

Replace `your_super_secret_private_key_here` with a secure, random string.

### 2. Required Headers

For every API request, include:

- `Authorization: Bearer <your_jwt_token>`
- `X-JWT-KEY: <your_decryption_key>`

The server will validate the JWT using a combination of your provided key and the private key from the environment.

### 3. Example `curl` Command

```bash
curl -X GET http://localhost:8080/users \
  -H "Authorization: Bearer <your_jwt_token>" \
  -H "X-JWT-KEY: <your_decryption_key>"
```

Replace `<your_jwt_token>` and `<your_decryption_key>` with your actual values.

> ⚠️ All endpoints (`/users`, `/users/create`, `/users/update`, `/users/delete`) require these headers for authentication.

---