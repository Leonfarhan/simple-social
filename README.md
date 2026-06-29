# Simple Social

Simple Social is a minimal Go HTTP API starter for a social application. It currently exposes a versioned health check endpoint and uses `chi` for routing.

## Tech Stack

- Go 1.25
- [chi](https://github.com/go-chi/chi) for HTTP routing and middleware
- [godotenv](https://github.com/joho/godotenv) for local environment variables
- [Air](https://github.com/air-verse/air) for optional live reload during development

## Project Structure

```text
.
├── cmd/
│   └── api/          # API entrypoint, router, and handlers
├── internal/
│   └── env/          # Environment helper functions
├── scripts/          # Utility scripts
├── docs/             # Project documentation
├── .air.toml         # Air live reload config
├── .env.example      # Environment variable template
├── go.mod
└── go.sum
```

## Requirements

- Go 1.25 or newer
- Air, optional, for hot reload

Install Air if you want live reload:

```bash
go install github.com/air-verse/air@latest
```

## Environment Variables

Copy the example environment file, then adjust the values if needed:

```bash
cp .env.example .env
```

Default local configuration:

```env
ADDR=:8080
DB_ADDR=
EXTERNAL_URL=localhost:8080
FRONTEND_URL=http://localhost:5173
ENV=development
FROM_EMAIL=
MAILTRAP_API_KEY=
SENDGRID_API_KEY=
```

| Variable | Description | Example |
| --- | --- | --- |
| `ADDR` | Address used by the HTTP server | `:8080` |
| `DB_ADDR` | PostgreSQL connection string | `postgres://admin:adminpassword@localhost/socialnetwork?sslmode=disable` |
| `EXTERNAL_URL` | Public API host used by generated Swagger docs | `localhost:8080` |
| `FRONTEND_URL` | Frontend origin used to build activation links | `http://localhost:5173` |
| `ENV` | Runtime environment; `production` disables email sandbox mode | `development` |
| `FROM_EMAIL` | Sender address used by Mailtrap or SendGrid | `noreply@example.com` |
| `MAILTRAP_API_KEY` | Mailtrap SMTP API key used by the active mailer | `mailtrap-api-key` |
| `SENDGRID_API_KEY` | SendGrid API key kept for parity with upstream implementation | `sendgrid-api-key` |

## Getting Started

Clone the repository and install dependencies:

```bash
git clone https://github.com/Leonfarhan/simple-social.git
cd simple-social
go mod download
```

Run the API:

```bash
go run ./cmd/api
```

Or run with live reload:

```bash
air
```

The server starts on the address configured by `ADDR`.

## API Endpoints

### Health Check

```http
GET /v1/health
```

Response:

```text
ok
```

Example:

```bash
curl http://localhost:8080/v1/health
```

## Build

Build the API binary:

```bash
go build -o ./bin/main ./cmd/api
```

Run the built binary:

```bash
./bin/main
```

## Development Notes

- The app loads environment variables from `.env` on startup.
- `.env` is ignored by Git because it can contain local or sensitive values.
- Keep `.env.example` updated when adding new required environment variables.
- The HTTP server uses request ID, real IP, logging, and recovery middleware.
- Current API version prefix: `/v1`.
