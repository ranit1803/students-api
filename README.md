 # 🧑‍🎓 **Students API - RESTful Service in Go**

A simple, production-style REST API built using **Golang** and **MySQL** to manage student records. This project demonstrates key backend concepts including routing, request validation, error handling, and graceful server shutdowns.

---

### 🚀 Features

* Create student records (POST)
* Get a student by ID (GET)
* Get all students (GET)
* Proper input validation using `go-playground/validator`
* Graceful server shutdown using `context` and OS signal handling
* Configurable via YAML files using `cleanenv`

---

### 🛠️ Tech Stack

* **Language:** Golang
* **Database:** MySQL
* **Libraries:**

  * `net/http` (standard library)
  * `slog` for structured logging
  * `go-sql-driver/mysql` for DB connection
  * `go-playground/validator/v10` for validation
  * `cleanenv` for config loading
* **Testing Tool:** Postman

---

### 📁 Project Structure

```
students-api/
├── cmd/                  # Entry point (main.go)
├── internal/
│   ├── config/           # Configuration loader
│   ├── http/
│   │   └── handlers/     # Request handlers
│   ├── storage/          # Storage interface + MySQL implementation
│   └── utils/
│       └── responses/    # Common response formatting
├── config/               # YAML configuration file
├── go.mod / go.sum
└── README.md
```

---

### 📦 Requirements

* Go 1.20+
* MySQL installed and running
* Git
* Postman (Testing the endpoints)

---

### 🧪 Quick Start

```bash
# Clone the project
git clone https://github.com/ranit1803/students-api.git
cd students-api

# Run with config path (make sure MySQL is running and config is set correctly)
go run cmd/students-api/main.go -config config/local.yaml
```

---

### 📚 Further Documentation

For API usage, setup instructions, and custom configs, check the [`docs/`](./docs/) folder.
---

### 🧠 Things I Learned
This project helped me strengthen my backend development skills using Golang. Here are the key concepts and technologies I learned and applied:
* 📦 Project Structuring in Go using cmd/, internal/, and modular packaging
* 🗃️ MySQL Integration with Go using database/sql and go-sql-driver/mysql
* ⚙️ Configuration Management using YAML files and the cleanenv library
* 🌐 RESTful API Design using net/http and http.ServeMux
* 🔁 Request Routing & Handlers with pattern-based routing and path parameters
* ✅ Input Validation with go-playground/validator/v10 for user input safety
* 🧵 Goroutines & Graceful Shutdown using os.Signal, syscall, and context.WithTimeout
* 📄 Standardized API Responses with reusable response format utilities
* 🔍 Structured Logging using the log/slog package
* 🧪 API Testing with Postman
* 💻 Git & GitHub Workflow for project version control and collaboration
