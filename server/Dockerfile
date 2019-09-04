FROM golang

ENV GO111MODULE=on

WORKDIR /api
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
EXPOSE 8585
ENTRYPOINT [ "/api/main" ]
