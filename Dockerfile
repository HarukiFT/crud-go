FROM golang:alpine
WORKDIR /build
COPY . ./
RUN go mod download
RUN go build -o app cmd/go-crud/main.go
ENTRYPOINT [ "./app" ]