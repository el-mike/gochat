#!/bin/bash

run_admin=0

while getopts 'a' flag;
do
    case "${flag}" in
        a) run_admin=1;;
    esac
done

if [ $run_admin -eq 1 ]; then
    docker-compose up postgres redis pgadmin
else
    docker-compose up postgres redis
fi
