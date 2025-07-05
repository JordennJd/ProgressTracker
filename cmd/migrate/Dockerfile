FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY "cmd/migrate/main.go" .
COPY migrations /app/migrations

RUN CGO_ENABLED=0 GOOS=linux go build -o /migrate ./main.go

CMD ["/migrate"]