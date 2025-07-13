# Arvo-Fin 🌊 - Personal Finance & Budget Tracker

Arvo-Fin is a powerful personal finance management system built in **Golang**, designed to help users track expenses, manage budgets, analyze spending habits, and get smart financial suggestions. It provides a fully featured **REST API**, a modern **Next.js frontend**, and scalable **microservices** architecture using Docker and Kubernetes.

---

## ✨ Features

### 📅 Core Modules

* User Authentication (JWT based)
* Portfolio Management (multi-portfolio support)
* Budget Tracking (Monthly/Category wise)
* Transactions (Income, Expense, Transfers)
* REST API for frontend/web/mobile apps

### ⚙️ Advanced Modules

* Category & Subcategory-wise reporting
* Monthly reports & exporting (PDF, CSV)
* Search transactions
* Suggestions using AI/ML
* Reminders & Alerts (Upcoming bills, savings goal)
* Offline-first mobile app (data sync on login)
* Secure Mutations for all data operations
* Email Notifications

---

## 🧰 Architecture Overview

```
fin-api/        => REST API Service (Gin)
fin-job/        => Background Job Processor (cron jobs)
fin-consumer/   => Kafka Consumer Service
next-app/       => Frontend Web App (Next.js)

internal/
  |- service/   => Shared services used by all services
  |- store/     => Database access with GORM
  |- model/     => Business models/entities

pkg/            => Shared utility packages (logger, config, etc.)

config/         => App and service config files
scripts/        => DB migrations, seeders, and tooling
```

---

## 🚀 Tech Stack

* **Language:** Go (Golang)
* **Web Framework:** Gin
* **Frontend Framework:** Next.js (React-based)
* **ORM:** GORM (MySQL)
* **Database:** MySQL
* **Authentication:** JWT
* **Caching:** Redis
* **Message Broker:** Kafka
* **Containerization:** Docker (multi-container)
* **Orchestration:** Kubernetes (K8s)
* **Scheduler:** Cron jobs (in `fin-job`)
* **Consumers:** Kafka consumers (real-time processing)
* **Email Notifications:** SMTP-based service
* **Multi Portfolio Support:** Each user can manage multiple financial portfolios
* **Mutations:** Secure and validated data mutations for all state-changing operations
* **REST API:** Comprehensive API layer for integration

---

## 🏠 Microservices

| Service      | Responsibility                             |
| ------------ | ------------------------------------------ |
| fin-api      | Exposes REST API endpoints                 |
| fin-consumer | Kafka consumer for real-time processing    |
| fin-job      | Cron jobs (daily sync, reminders, reports) |
| next-app     | Web UI built in Next.js                    |

---

## 🔍 API Reference

### 🔑 Auth (`/auth`)

* `POST /auth/login` - Login with credentials
* `POST /auth/token` - Refresh token
* `POST /auth/send-otp` - Send OTP
* `POST /auth/register` - Register user
* `POST /auth/reset-password` - Reset password

### 👤 User (`/user`)

* `GET /user/` - Get all users
* `GET /user/:id` - Get user by ID
* `POST /user/` - Create new user
* `PUT /user/:id` - Update user by ID
* `DELETE /user/:id` - Delete user

### 💳 Account (`/account`)

* `GET /account/types` - Get account types
* `GET /account/list/:portfolioId` - Get accounts for portfolio
* `GET /account/:id` - Get account by ID
* `POST /account/` - Create account
* `PUT /account/:id` - Update account
* `DELETE /account/:id` - Delete account

### 📷 Avatar (`/avatar`)

* `POST /avatar/` - Create avatar
* `PUT /avatar/:id` - Update avatar
* `GET /avatar/:id` - Get avatar by ID
* `GET /avatar/type/:type` - Get avatars by type

### 📂 Category (`/category`)

* `GET /category/list/:portfolioId` - List categories
* `GET /category/:id` - Get category by ID
* `POST /category/` - Create category
* `PUT /category/:id` - Update category
* `DELETE /category/:portfolioId/:id` - Delete category

### 🤑 Me (`/me`)

* `GET /me/` - Get current user
* `GET /me/:userId` - Get user by ID
* `GET /me/settings/` - Get user settings
* `PUT /me/` - Update current user
* `PUT /me/settings/` - Update settings

### 🏛️ Portfolio (`/portfolio`)

* `GET /portfolio/` - List portfolios
* `GET /portfolio/:id` - Get portfolio
* `POST /portfolio/` - Create portfolio
* `PUT /portfolio/:id` - Update portfolio
* `DELETE /portfolio/:id` - Delete portfolio

### 💸 Transaction (`/transaction`)

* `GET /transaction/` - List transactions (with search)
* `GET /transaction/:id` - Get transaction
* `POST /transaction/` - Create transaction
* `PUT /transaction/:id` - Update transaction
* `DELETE /transaction/:id` - Delete transaction

---

## 🚧 Upcoming Features

* Sync conflict resolver between mobile and server
* Integration with email/SMS for reminders
* GraphQL API (optional)
* OpenAPI docs (Swagger)
* Predictive budgeting with ML model

---

## 🚀 Getting Started

### Clone the repo:

```bash
git clone https://github.com/Uttamnath64/arvo-fin.git
cd arvo-fin
```

### Run with Docker:

```bash
docker-compose --env-file ./backend/app/config/env/.env up --build
```

---

## ✅ Best Practices Followed

* Clean Architecture (Domain-Driven)
* Dependency Injection
* Layered Structure: API/Service/Store/Model
* Unit & Integration Testing Ready
* Configurable via `.env` files
* Secure Auth (JWT, Password Hashing)
* Scalable via microservices and Kafka pub-sub

---

## 🚩 Contributing

1. Fork this repo
2. Create feature branch (`git checkout -b feature/foo`)
3. Commit changes (`git commit -am 'Add foo'`)
4. Push to branch (`git push origin feature/foo`)
5. Create a new Pull Request

---

## 🌟 Author

**Uttam Nath**
Golang Developer | Learner | Builder

---

## 🎁 License

This project is licensed under the MIT License - see the `LICENSE` file for details.

---
