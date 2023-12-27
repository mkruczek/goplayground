#!/bin/bash

docker build -t my-postgres .

docker run --name my-postgres -p 5432:5432 -d my-postgres

sleep 10

go run main.go

docker rm -f my-postgres