# GRPC with Golang

## Description

Practicing/learning grpc with Golang, as well as, make a crud app with mongo using docker-compose.

## Install App

1. Install [ Docker Engine ](https://docs.docker.com/engine/install/) :fire:
2. `cd blog && docker-compose up -d`
3. `cd .. && make blog` gonna create the binaries...
4. `./bin/blog/server`, open another terminal and execute `./bin/blog/client`
5. With that you gonna see the execute of the functions called in main.go inside client folder.

PS: the best practices are:

1. Create proto first.
2. Create logic on the server.
3. Create logic on the client.
