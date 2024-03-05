# alpine linux docker image
# with only mainApp executable
FROM alpine:latest

RUN mkdir /app

COPY mainApp /app

CMD [ "/app/mainApp" ]
