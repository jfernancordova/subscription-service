## Subscription Service

![Github](https://img.shields.io/github/actions/workflow/status/jfernancordova/subscription-service/go.yml)&nbsp;
![Language](https://img.shields.io/badge/language-Go-orange.svg)&nbsp;
![Platform](https://img.shields.io/badge/platform-docker-blue.svg)&nbsp;
![Database](https://img.shields.io/badge/database-postgreSQL-pink.svg)&nbsp;
![Cache](https://img.shields.io/badge/cache-redis-red.svg)&nbsp;

A flexible subscription management system. The idea is providing features such as: sending emails, handling errors and shutdown using concurrency (channels, waitgroups and goroutines).

### Features

- User registration and authentication.
- Subscription plans.
- Email notifications.
- Invoices management.
- User manuals.

### UI vs API

This project is a **server-rendered web UI** (Go + HTML templates). It is **not** a REST/JSON API service.

- Use it through a browser (HTML forms + session cookies).
- There is no `/api` surface that returns JSON.
- Protected routes under `/members/*` require a logged-in session (cookie), not an API token.

### Pages / routes

| Path | Method | What it does |
|------|--------|----------------|
| `/` | GET | Home |
| `/login` | GET, POST | Login form |
| `/register` | GET, POST | Registration form |
| `/logout` | GET | Log out |
| `/activate` | GET | Account activation (from email link) |
| `/members/plans` | GET | Choose a subscription plan (auth required) |
| `/members/subscribe` | GET | Subscribe to a plan via `?id=…` (auth required) |

### Prerequisites

Before you begin, ensure you have met the following requirements:

- Go (Golang). You can download it from [here](https://golang.org/dl/).
- Docker. You can download it from [here](https://docs.docker.com/get-docker/).

### Testing

To check the coverage and run all the tests:

```bash
go test -v ./cmd/web
```

### Getting Started

Follow these steps to get the Subscription Service up and running on your local machine:

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/jfernancordova/subscription-service.git
   ```

2. Navigate to the project directory:

   ```bash
   cd subscription-service
   ```

3. Run the application:

   ```bash
   make start
   ```

   This starts Postgres, Redis, and MailHog via Docker Compose, then builds and runs the Go web app.

4. Access the service in a browser:

   - App: http://localhost:8080
   - MailHog (emails): http://localhost:8025

5. Populate data to the database:

   ```bash
   make load
   ```

6. Log in with:

   ```text
   email: admin@example.com
   password: verysecret
   ```

7. To restart or stop the subscription service app:

   ```bash
   make restart
   make stop
   ```

### Port configuration

The web server listens on port **8080** by default (ports below 1024, such as 80, require root and will fail with `bind: permission denied` for a normal user).

To use a different port:

```bash
PORT=3000 ./subscription-service
```

### Troubleshooting

**`listen tcp :80: bind: permission denied`**

Binding to port 80 needs elevated privileges. This project defaults to port **8080**. Rebuild/restart after pulling the change:

```bash
make restart
```

Then open http://localhost:8080

**Containers are up but the site does not load**

`make start` brings up Docker services even if the Go process fails. Check that a `subscription-service` process is running, or look for a panic/log right after `Started!`. Then run `make restart`.
