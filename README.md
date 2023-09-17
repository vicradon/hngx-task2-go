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

## Deployment

To deploy the app, you can easily use the provided Docker image. Build the image using the command below:

```bash
docker build -t <DOCKERHUB_USERNAME>/personcrudapi .
```

Then run the Docker image using the command below:

```bash
docker run -p 5000:8080 <DOCKERHUB_USERNAME>/personcrudapi
```
