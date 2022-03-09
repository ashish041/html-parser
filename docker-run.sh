#!/bin/bash

docker build --tag main .
sudo docker run -d -p 8080:8080 main
