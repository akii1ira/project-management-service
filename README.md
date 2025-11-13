# Project Management Service

A simple **Project Management Service** built with **Go**, **PostgreSQL**, and **Docker**.  
This service allows you to create, update, and manage tasks in projects.

---

## 🚀 Features

- Create, update, delete tasks
- PostgreSQL database integration
- Dockerized for easy setup
- REST API endpoints

---

## 🛠 Tech Stack

- **Backend:** Go 1.25+
- **Database:** PostgreSQL 15
- **Routing:** `chi` / `gorilla/mux`
- **Containerization:** Docker, Docker Compose

---

## 📦 Prerequisites

Make sure you have installed:

- [Go 1.25+](https://golang.org/dl/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

---

## ⚡ Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/akii1ira/project-management-service.git
cd project-management-service


🔧 API Endpoints
Method	Endpoint	Description
GET	/tasks	Get all tasks
GET	/tasks/{id}	Get task by ID
POST	/tasks	Create a new task
PUT	/tasks/{id}	Update a task
DELETE	/tasks/{id}	Delete a task


📝 Environment Variables
Create a .env file or set environment variables for PostgreSQL:
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=projectdb
POSTGRES_HOST=db
POSTGRES_PORT=5432
```
