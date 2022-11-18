FROM golang:alpine
WORKDIR /usr/src/app
COPY . .
EXPOSE 8000
CMD [ "go","run","src/server.go" ]
