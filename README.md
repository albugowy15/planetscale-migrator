# PlanetScale Migrator

This is a command-line interface tool designed to facilitate the migration of databases from PlanetScale to other MySQL-compatible database services.

## Prerequisites

### Go 1.22.0

Ensure that you have Go installed on your system. Always use the latest version of Go for optimal performance and compatibility.

### Exporting Data from PlanetScale

To export data from PlanetScale, refer to the instructions provided in this article: [Hobby tier deprecation - FAQ](https://planetscale.com/docs/concepts/hobby-plan-deprecation-faq#how-do-i-migrate-off-of-planetscale-)

## How to run

1. Begin by cloning this repository

```bash
git clone https://github.com/albugowy15/planetscale-migrator.git
```

2. Create a .env file by copying the provided .env.example file and fill in all the required environment variables

```bash
cp .env.example .env
```

3. Build the Go binary

```bash
go build -o pscale-migrator main.go
```

4. Execute the binary and specify either a .sql file or a directory containing .sql files

```bash
# Execute a single file, such as sample.sql
./pscale-migrator -file sample.sql

# Execute all SQL files within the migrations directory
./pscale-migrator -dir migrations/
```
