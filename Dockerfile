FROM golang:alpine as builder

WORKDIR /usr/src/app

COPY ./go.mod ./go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN GOOS=linux go build  -v -o /app 


FROM alpine:latest as deployment
EXPOSE 80
WORKDIR /usr/local/bin/myapp
COPY --from=builder /app /usr/local/bin/myapp/app
COPY ./templates ./templates
COPY ./postgres-data ./postgres-data


CMD ["./app"]