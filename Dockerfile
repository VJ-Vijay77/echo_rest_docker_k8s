FROM golang:1.19-alpine

RUN mkdir /app

WORKDIR /app

COPY binary /app/

CMD [ "/app/binary" ]