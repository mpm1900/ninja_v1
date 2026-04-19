# Ninja Battler V1

A modern web-based turn-based combat game inspired by the Naruto universe. Features an animated background, real-time combat logic, and a TanStack-powered frontend.

## Prerequisites

Ensure you have the following installed:

- **Go**: 1.25+
- **Docker & Docker Compose**: For the PostgreSQL database
- **Bun**: (Recommended) or Node.js/pnpm for the frontend
- **Goose**: For database migrations (`go install github.com/pressly/goose/v3/cmd/goose@latest`)
- **Air**: (Optional) For Go hot-reloading (`go install github.com/air-verse/air@latest`)

## Getting Started

### 1. Environment Setup

Copy the example environment file and adjust if necessary:

```bash
cp .env.example .env
```

### 2. Infrastructure

Start the PostgreSQL database using Docker Compose:

```bash
make up
```

### 3. Database Migrations

Run the migrations to set up the database schema:

```bash
make migrate
```

### 4. Running the Backend

You can run the backend in two ways:

**Standard run:**
```bash
make run
```

**With hot-reloading (requires Air):**
```bash
air
```

The server will start on the port specified in your `.env` (default is `:3005`).

### 5. Running the Frontend

Navigate to the `web` directory, install dependencies, and start the development server:

```bash
cd web
bun install
bun dev
```

The frontend will be available at [http://localhost:3000](http://localhost:3000).
