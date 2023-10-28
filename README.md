## Subscription Service

![Language](https://img.shields.io/badge/language-Go-orange.svg)&nbsp;
![Platform](https://img.shields.io/badge/platform-docker-blue.svg)&nbsp;
![Database](https://img.shields.io/badge/database-postgreSQL-pink.svg)&nbsp;
![Cache](https://img.shields.io/badge/cache-redis-red.svg)&nbsp;

A flexible subscription management system. The idea is providing features such us: sending emails, handling errors and shutdown using concurrency (channels, waitgroups and goroutines).

### Features

- User registration and authentication.
- Subscription plans.
- Email notifications.
- Invoices management.
- User manuals.

### Prerequisites

Before you begin, ensure you have met the following requirements:

- Go (Golang). You can download it from [here](https://golang.org/dl/).
- Docker. You can download it from [here](https://docs.docker.com/get-docker/).


### Testing

To check the coverage and run all the tests:

   ```bash
      cd cmd/web
      go test -v
   ```

### Getting Started

Follow these steps to get the Subscription Service up and running on your local machine:

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/jfernancordova/subscription-service.git

2. Navigate to the project directory:

   ```bash
   cd subscription-service

3. Run the application:
   ```bash
   make start

4. Access the service locally by opening a web browser and visiting http://localhost.

5. To see the mails through [mailhog](https://github.com/mailhog/MailHog): http://localhost:8025

6. To populate data to the database:
   ```bash
   make load

7. To access:
   ```bash
   email: admin@example.com
   password: verysecret

8. To `restart/stop` the subscription service app
   ```bash
   make restart
   make stop
