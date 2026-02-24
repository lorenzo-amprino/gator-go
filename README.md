# gator

Simple CLI RSS feed manager.

## Install from GitHub

Prerequisites:
- Go 1.25+ (see [go.mod](go.mod))
- PostgreSQL database

Option A — Install directly:
```sh
go install github.com/lorenzo-amprino/gator-go@latest
```

Option B — Build from source:
```sh
git clone https://github.com/lorenzo-amprino/gator-go.git
cd gator-go
go build -o gator .
# run:
./gator <command>
```

Config
- Create ~/.gatorconfig.json (used by the app: see [internal/config/config.go](internal/config/config.go))
Example:
```json
{
  "db_url": "postgres://user:pass@localhost:5432/dbname?sslmode=disable",
  "current_user_name": ""
}
```

Database schema
- Apply the SQL files in `sql/schema` to your Postgres instance:
```sh
psql "<DB_URL>" -f sql/schema/001_users.sql
psql "<DB_URL>" -f sql/schema/002_feeds.sql
psql "<DB_URL>" -f sql/schema/003_feed_follow.sql
```

Run
- The main entry point is [main.go](main.go). After installation/build run the binary and use the commands documented in the source.