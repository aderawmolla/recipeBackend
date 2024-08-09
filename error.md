
command-1

go get -u github.com/gin-gonic/gin
 
error-1 

go: go.mod file not found in current directory or any parent directory.
To build and install a command, use 'go install' with a version,
like 'go install example.com/cmd@latest'
For more information, see https://golang.org/doc/go-get-install-deprecation
or run 'go help get' or 'go help install'.

answer-1

go mod init project-name


command-2
 docker compose up -d
 error-2
 > ERROR [golang_app build 5/5] RUN go build -o /app/start                                               0.4s
------
 > [golang_app build 5/5] RUN go build -o /app/start:
0.321 main.go:4:2: //go:build comment without // +build comment
0.321 /go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/context.go:23:2: //go:build comment without // +build comment

answer-2

adding 
