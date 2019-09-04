FROM golang as builder

ENV GO111MODULE=on

WORKDIR /api
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM scratch

# copy binary to app directory in second stage
COPY --from=builder /api/main /api/
EXPOSE 8585
ENTRYPOINT [ "/api/main" ]
