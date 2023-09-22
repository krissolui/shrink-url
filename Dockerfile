FROM alpine:latest

RUN mkdir /app

COPY ./bin/shrink-url /app/
COPY ./appsettings.json /app/

WORKDIR /app

CMD [ "./shrink-url" ]