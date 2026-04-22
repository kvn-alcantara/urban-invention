# Urban Invention

A small Go HTTP service scaffolded for future Dockerization.

## Run

```bash
go run ./cmd/app
```

Set `PORT` to change the listening port.

## Test

```bash
go test ./...
```

## Discord Alerts (GitHub Actions)

This repository includes a workflow at `.github/workflows/discord-alerts.yml` that sends notifications to Discord for:

- every `push` (new commits)
- every merged pull request into `main`

### Setup

1. Create a Discord webhook in your server/channel.
2. In GitHub, go to `Settings > Secrets and variables > Actions`.
3. Add a repository secret named `DISCORD_WEBHOOK_URL` with the webhook URL.

After this, each commit push and each merge to `main` will post an alert message to your Discord channel.
