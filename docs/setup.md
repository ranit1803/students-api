# 🛠️ Students API - Setup Guide

This document will walk you through the steps to set up, configure, and run the Students API project locally.

---

## 📦 Prerequisites

Make sure you have the following installed:

- [Golang](https://golang.org/dl/) 1.20+
- [MySQL Server](https://dev.mysql.com/downloads/mysql/)
- [Postman](https://www.postman.com/) or any API client for testing
- Git

---

## 🧾 Configuration

This project uses a `config/local.yaml` file for configuration via the `cleanenv` package.

### Sample `config/local.yaml`

```yaml
env: "dev"

mysql:
  host: "localhost"
  port: 3306
  user: "root"
  password: "your_password"
  dbname: "students_db"

http_server:
  address: "localhost:1803"
````

> ⚠️ Create a `.env` file or set an environment variable:

```bash
export CONFIG_PATH=config/local.yaml
```

Or use CLI flag when running:

```bash
go run cmd/students-api/main.go -config=config/local.yaml
```

---

## 🚀 Running the Project

### 1. Clone the Repository

```bash
git clone https://github.com/ranit1803/students-api.git
cd students-api
```

### 2. Run the App

```bash
go run cmd/students-api/main.go -config=config/local.yaml
```

> ✅ You should see `Server Listening at: localhost:1803`

---

## 📡 API Endpoints

### ➕ Create a Student

* **URL:** `POST /api/students`
* **Body Example (JSON):**

```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "age": 22
}
```

### 📥 Get a Student by ID

* **URL:** `GET /api/students/{id}`

### 📃 Get All Students

* **URL:** `GET /api/students`

---

## 🧪 Testing

Use Postman or CURL to test the endpoints:

```bash
curl -X POST http://localhost:1803/api/students \
-H "Content-Type: application/json" \
-d '{"name":"Alice","email":"alice@mail.com","age":20}'
```

---

## 🧹 Graceful Shutdown

To stop the server safely (e.g., with cleanup), press `CTRL+C`.

---

## 🧠 Credits

* Inspired by [Coder’s Gyan YouTube Channel](https://www.youtube.com/@codersgyan)
