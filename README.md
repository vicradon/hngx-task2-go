# Person CRUD API

A CRUD API for a person object implemented in Go.

## On Startup

Install goose using the command:

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

## Create a Migration

```bash
goose create make_email_unique_and_auto_increment_id sql
```

## Run Migrations

```bash
goose sqlite3 ../database.db up
```
