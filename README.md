# TechYar

TechYar is a Persian RTL full-stack web application for browsing and presenting server products, with user authentication, profiles, and server details.

## Tech Stack

**Frontend**

* React
* Vite
* React Router
* JavaScript / JSX

**Backend**

* Go
* Gin
* GORM
* PostgreSQL
* JWT

## Features

* User registration and login
* JWT-based authentication
* User profile management
* Server/product listing
* Server search and category filtering UI
* Server details page
* Admin dashboard prototype
* Persian RTL interface

## Project Structure

```text
TechYar/
├── backend/
│   ├── cmd/
│   ├── internal/
│   │   ├── config/
│   │   ├── db/
│   │   ├── dto/
│   │   ├── handler/
│   │   ├── middleware/
│   │   ├── model/
│   │   ├── repository/
│   │   ├── router/
│   │   ├── service/
│   │   └── utils/
│   └── go.mod
│
└── frontend/
    └── net-eng/
        ├── src/
        ├── public/
        ├── package.json
        └── vite.config.js
```

## Getting Started

### Backend

```bash
cd backend
go run ./cmd
```

The backend runs on:

```text
http://localhost:8080
```

Configure PostgreSQL using the following environment variables:

```env
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=techyar
DB_PORT=5432
```

### Frontend

```bash
cd frontend/net-eng
npm install
npm run dev
```

The frontend runs on the Vite development server.

## Status

TechYar is currently under active development.
