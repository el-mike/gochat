# gochat
https://img.shields.io/github/workflow/status/el-Mike/gochat/Gochat%20API

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

There is VSC launch configuration available in the repository. In order to run Gochat API using VSC debugging, run `docker-compose up postgres redis`, and then start `[Gochat] Launch API` VSC configuration. 
