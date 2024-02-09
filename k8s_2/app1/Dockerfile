FROM golang:1.21-alpine

COPY . /app

WORKDIR /app

RUN ls -la . && \
    go install . && \
    which app1

ENTRYPOINT [ "/go/bin/app1" ]