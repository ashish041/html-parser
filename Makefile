
# make all to run the app in local
all: install-go up

install-go:
	sh install-go.sh

up: test build run call-api

build:
	sh ./build.sh
run:
	sh ./run.sh
test:
	sh ./test.sh
call-api:
	sh ./call-api.sh

# make docker to run the app in docker
docker: docker-run call-api
docker-run:
	sh ./docker-run.sh
