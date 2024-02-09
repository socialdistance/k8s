FROM golang:1.21-alpine

COPY . /app

WORKDIR /app

RUN ls -la . && \
    go install . && \
    which k8s

ENTRYPOINT [ "/go/bin/k8s" ]