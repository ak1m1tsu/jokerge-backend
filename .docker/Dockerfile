FROM golang:alpine3.19 as modules
COPY go.mod go.sum /modules/ 
WORKDIR /modules
RUN go mod download

FROM golang:alpine3.19 as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN GOOS=linux GOARCH=amd64 go build -tags migrate -o /bin/jokerge ./cmd/app

FROM scratch
COPY --from=builder /bin/jokerge /jokerge
CMD [ "/jokerge" ]
