# golang docker image for building mainapp
FROM golang:1.21.4-alpine as mainapp

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o mainApp .

# alpine linux docker image
# with only mainapp executable
FROM alpine:latest

RUN mkdir /app

COPY --from=mainapp /app/mainApp /app

CMD ["/app/mainApp"]