FROM uselagoon/commons AS builder

RUN apk add go
WORKDIR /app
COPY . /app
RUN go build -o temp-handler /app/main.go

ENTRYPOINT ["/app/temp-handler"]


