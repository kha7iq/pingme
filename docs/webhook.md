# Webhook Server

PingMe can also run as a small HTTP server that accepts webhooks and forwards them to any of the supported services (Telegram, Slack, Discord, email, etc.).
This is useful when other tools want a webhook service like grafana alerts.

## What the server does

- Listens on a port (default `8080`).
- Exposes a `/webhook` endpoint that accepts JSON.
- Reads credentials from the same environment variables the CLI already uses.
- Dispatches the incoming payload to the requested service.

You still configure services with environment variables; the webhook is just another way to trigger them.

---

## Starting the server

Basic run:

```bash
pingme serve
```

Defaults:

- Host: `0.0.0.0`
- Port: `8080`

Custom host/port:

```bash
pingme serve --host 127.0.0.1 --port 9000
```

Environment overrides also work:

```bash
export PINGME_HOST=0.0.0.0
export PINGME_PORT=8080
pingme serve
```

---

## Configuring services (env vars)

Same idea as the CLI: set env vars before you start the server.

Example: Telegram

```bash
export TELEGRAM_TOKEN="123456:ABC-your-bot-token"
export TELEGRAM_CHANNELS="-123456789"

pingme serve
```

Example: Slack

```bash
export SLACK_TOKEN="xoxb-..."
export SLACK_CHANNELS="alerts,ops"

pingme serve
```

You can set multiple services at the same time; each webhook chooses which one to use via the `service` field.

---

## Request format

`POST /webhook` with JSON body:

```bash
{
  "service": "telegram",
  "message": "Your message here",
  "title": "Optional title",
  "priority": 1
}
```

Fields:

- `service` (string, required): which integration to use, e.g. `"telegram"`, `"slack"`, `"email"`, `"pushover"`, etc.
- `message` (string, required): main message body.
- `title` (string, optional): subject/title where supported (email, pushover, etc.).
- `priority` (int, optional): used by services that support it (e.g. Pushover, Gotify).

---

## Simple examples

### Telegram example

Assuming `TELEGRAM_TOKEN` and `TELEGRAM_CHANNELS` are set:

```bash
curl -X POST http://localhost:8080/webhook \
  -H "Content-Type: application/json" \
  -d '{
    "service": "telegram",
    "title": "Deploy finished",
    "message": "Backend has been deployed to production."
  }'
```

### Slack example

With `SLACK_TOKEN` and `SLACK_CHANNELS` set:

```bash
curl -X POST http://localhost:8080/webhook \
  -H "Content-Type: application/json" \
  -d '{
    "service": "slack",
    "title": "High CPU",
    "message": "CPU usage on node-1 is above 90%."
  }'
```

### Email example

With `EMAIL_*` env vars configured:

```bash
curl -X POST http://localhost:8080/webhook \
  -H "Content-Type: application/json" \
  -d '{
    "service": "email",
    "title": "Nightly backup",
    "message": "Backup completed successfully."
  }'
```

---

## Authentication (optional but recommended)

By default, anyone who can reach `/webhook` can send messages. For local use that's fine; for anything exposed to the outside, turn on auth.

### API key

```bash
export PINGME_AUTH_METHOD="apikey"
export PINGME_API_KEYS="secret-key-1,secret-key-2"

pingme serve
```

Call with:

```bash
curl -X POST http://localhost:8080/webhook \
  -H "Authorization: Bearer secret-key-1" \
  -H "Content-Type: application/json" \
  -d '{"service":"telegram","message":"API key protected"}'
```

### Basic auth

```bash
export PINGME_AUTH_METHOD="basic"
export PINGME_BASIC_USER="pingme"
export PINGME_BASIC_PASS="changeme"

pingme serve
```

Call with:

```bash
curl -u pingme:changeme \
  -H "Content-Type: application/json" \
  -d '{"service":"telegram","message":"Basic auth"}' \
  http://localhost:8080/webhook
```

### HMAC

For when you want to verify the payload hasn't been tampered with:

```bash
export PINGME_AUTH_METHOD="hmac"
export PINGME_HMAC_SECRET="super-secret"

pingme serve
```

Clients send a hex HMAC-SHA256 in `X-Signature` based on the raw body.

Example client (Python-style pseudocode):

```bash
import hmac
import hashlib
import requests

secret = "super-secret"
payload = '{"service":"telegram","message":"Hello"}'

signature = hmac.new(
    secret.encode(),
    payload.encode(),
    hashlib.sha256
).hexdigest()

requests.post(
    "http://localhost:8080/webhook",
    data=payload,
    headers={
        "Content-Type": "application/json",
        "X-Signature": signature
    }
)
```

---

## Endpoints

- `POST /webhook`  
  Main endpoint; accepts JSON as described above.

- `GET /health`  
  Simple health check. Returns HTTP 200 with a small JSON body.

- `GET /`  
  Basic info about the server and available endpoints.

---

## Typical setups

### CI/CD

Use the webhook to notify after jobs:

```bash
- name: Notify PingMe
  if: always()
  run: |
    curl -X POST "$PINGME_URL/webhook" \
      -H "Content-Type: application/json" \
      -d '{
        "service": "slack",
        "title": "CI run finished",
        "message": "Workflow '${{ github.workflow }}' finished with status: ${{ job.status }}"
      }'
```

### Monitoring / alerting

Most tools that support "webhook" or "HTTP" notifications can just post JSON to `/webhook`, and you map fields into `service`, `title`, `message`.

---

## Health checks

Your server has a health endpoint you can plug into Docker, Kubernetes, or uptime checks:

```bash
curl http://localhost:8080/health
```

Returns something like:

```bash
{"status":"ok","service":"pingme"}
```

Use this for:

- Docker health checks
- Kubernetes liveness/readiness probes
- Load balancer health checks
- External uptime monitors

---

## Quick checklist

- [ ] Set env vars for at least one service (`TELEGRAM_*`, `SLACK_*`, etc.).
- [ ] Optionally set auth (`PINGME_AUTH_METHOD` + related vars).
- [ ] Start the server with `pingme serve`.
- [ ] Hit `GET /health` to confirm it's running.
- [ ] Send a test `POST /webhook` with your desired `service` and `message`.
