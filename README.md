Of course! Here's your friendly 💻 API guide all wrapped up in Markdown for easy copy-paste or saving as `README.md`:

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