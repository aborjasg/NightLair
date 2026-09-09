# NightLair

NightLair is a small Go-based application prototype for a hospitality and overnight-rental experience. The project is designed around the idea of creating a private, comfortable, and memorable stay for guests by combining a backend API with a simple web frontend.

## Purpose

The core vision of NightLair is to reimagine the overnight rental experience as a personal retreat: a place where guests can rest, recharge, and enjoy amenities tailored to their needs. In practice, this repository acts as the foundation for a booking and guest-experience platform, starting with a lightweight backend and a front-end interface for interacting with the system.

## Project structure

This repository is organized into two main Go services:

- `go-api/` — backend API service
- `go-webapp/` — web frontend application

### Backend API (`go-api`)

The API service is a Go HTTP server built to expose data and support a business workflow around prospects or guest-related leads. It follows a common layered structure:

- `main.go` — starts the API server and wires dependencies together
- `internal/handlers/` — HTTP handlers for API endpoints
- `internal/services/` — application logic and orchestration
- `internal/repositories/` — database access and persistence logic
- `internal/models/` — shared response and data models

The current implementation includes a `ProspectRepository` that queries a MySQL database and returns the count of rows in a `Prospects` table via a `/api/prospects` endpoint. This is a strong starting point for future functionality such as guest inquiries, bookings, or property leads.

### Web app (`go-webapp`)

The web app is a Go application that serves a simple HTML interface from the `templates/` folder and static assets from `static/`. It is designed to be a lightweight front-end for the platform and currently renders a basic page that triggers a fetch to the API.

Key pieces include:

- `main.go` — startup code for the web server
- `templates/index.html` — the main page for the UI
- `static/` — place for CSS, JS, or image files

### Relationship between the components

The two applications are intentionally separated:

- the API handles business logic and data access
- the web app presents the data to users through a browser interface

This decoupled structure makes it easier to extend the project later with richer booking flows, authentication, dashboards, or additional services.

## Current implementation status

This repository is currently in an early MVP/prototype stage. The core code establishes the foundation for a hospitality booking-style platform, but the feature set is intentionally minimal and focuses on demonstrating the project structure and API/web interaction pattern.

## How to run

### Prerequisites

- Go installed on your machine
- MySQL database available for the API service

### Start the API

```bash
cd go-api
go run .
```

The API listens on:

- `http://localhost:8080`

### Start the web app

```bash
cd go-webapp
go run .
```

The frontend listens on:

- `http://localhost:8081`

## Future direction

NightLair can evolve into a full booking and guest-management platform with features such as:

- property listings and search
- guest booking flows
- reservation management
- personalized amenities and stay preferences
- admin dashboards for operations
- secure authentication and role-based access

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

