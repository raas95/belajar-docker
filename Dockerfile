FROM golang:alpine

RUN apk update && apk add git 

WORKDIR /app

copy . .

RUN go mod tidy 
RUN go build -o binary

ENTRYPOINT [ "./binary" ]