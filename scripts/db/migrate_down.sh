#!/bin/bash

USER=$(cat .env | grep POSTGRES_USER | awk -F"=" '{print $2}')
PASSWORD=$(cat .env | grep POSTGRES_PASSWORD | awk -F"=" '{print $2}')
DB=$(cat .env | grep POSTGRES_DB | awk -F"=" '{print $2}')
HOST=$(cat .env | grep POSTGRES_HOST | awk -F"=" '{print $2}')
PORT=$(cat .env | grep POSTGRES_PORT | awk -F"=" '{print $2}')

CONNECTION_STRING="postgres://${USER}:${PASSWORD}@${HOST}:${PORT}/${DB}?sslmode=disable"

migrate -database ${CONNECTION_STRING} -path db/migrations down
