#!/bin/bash

echo 'curl http://localhost:8080?url="yahoo.com"'
curl http://localhost:8080?url="yahoo.com"

echo '\n\ncurl http://localhost:8080?url="https://gmail.com"'
curl http://localhost:8080?url="https://gmail.com"

echo '\n\ncurl http://localhost:8080?url="http://wwww.google.com"'
curl http://localhost:8080?url="http://www.google.com"
