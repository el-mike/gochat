# gochat
![build](https://img.shields.io/github/workflow/status/el-Mike/gochat/Gochat%20API)

Gochat is a simple chatting application, created for learning purposes. 

# Running the application

1. Gochat uses [godotenv]("github.com/joho/godotenv") to store local env variables. Add `.env` file to the root directory, with following structure:

```
POSTGRES_USER=
POSTGRES_PASSWORD=
POSTGRES_DB=
POSTGRES_PORT=
POSTGRES_HOST=

PGADMIN_DEFAULT_EMAIL=
PGADMIN_DEFAULT_PASSWORD=

REDIS_HOST=
REDIS_PORT=
REDIS_PASSWORD=

API_SECRET=

```

2. Run `go install` to compile and install all required packages and dependencies.
3. Run `docker-compose up` to start Gochat API and all required dependencies.

## Debugging

There is VSC launch configuration available in the repository. In order to run Gochat API using VSC debugging, run `docker-compose up postgres redis` or `./scripts/run_deps.sh`, and then start `[Gochat] Launch API` VSC configuration. 

# Development

## Prerequisites

1. Install [golangci-lint](https://golangci-lint.run/usage/install/)
2. Set your IDE to use golangci-lint ([instructions](https://golangci-lint.run/usage/integrations/))
3. Install [python3](https://www.python.org/download/releases/3.0/)
4. Run `git config core.hooksPath .githooks` to wire up project's git hooks

## Conventions

This repository follows [ConventionalCommits](https://www.conventionalcommits.org/en/v1.0.0/) specification for creating commit messages. There is `prepare-commit-msg` hook set up to ensure following those rules. Branch names should also reflect the type of work it contains - one of following should be used:
* `feature/<task-description>`
* `bugfix/<task-description>`
* `chore/<task-description>`

