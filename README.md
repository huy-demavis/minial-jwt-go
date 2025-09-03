# Minial JWT Go

[![Go Version](https://img.shields.io/badge/go-1.22+-blue.svg)](https://golang.org/doc/devel/release.html)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

A minimal **JWT authentication service** built with [Go](https://go.dev/) and [Gin](https://gin-gonic.com/).  
This project demonstrates secure user authentication, role-based access, middleware patterns, and environment-based configuration.  

---

## ✨ Features
- 🔑 User registration and login with JWT  
- 🔒 Protected routes with middleware  
- 🛠 Configurable via `.env`  
- 🗂 Clean project structure (controllers, middleware, services)  
- ⚡ Ready for containerization with Docker Compose  

---

## 🚀 Getting Started

### Prerequisites
- Go 1.22+
- Docker (optional, for containerized setup)

### Installation
```bash
git clone https://github.com/huy-demavis/minial-jwt-go.git
cd minial-jwt-go
cp .env.example .env
go mod tidy
go run main.go
