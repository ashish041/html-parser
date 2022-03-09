# HTML Parser
Web application to parse website and retrieve following informations

- HTML Version
- Page Title
- Headings count by level
- Amount of internal and external links
- Amount of inaccessible links
- If a page contains a login form

Command to install Go and run web application
```
make
```

Command to run web application
```
make up
```

Command to run web application in docker
```
make docker
```

Command to generate output

```
curl http://localhost:8080?url="yahoo.com"
curl http://localhost:8080?url="https://gmail.com"
curl http://localhost:8080?url="http://www.google.com
```

Please make sure the web application is not running on the host machine before running in docker.
Command to check whether web application running on port 8080.
```
netstat -nlp |grep 8080
```
