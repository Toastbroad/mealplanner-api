#!/bin/bash 
source env.sh

rm ./docker-compose/docker-compose.yml
rm ./database/database.go

envsubst < "./docker-compose/template.yml" > "./docker-compose/docker-compose.yml"
envsubst < "./database/database.go-template" > "./database/database.go"