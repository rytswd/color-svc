# Build
FROM golang:1.15-alpine as build

WORKDIR /go/src/github.com/rytwsd/color-svc

COPY go.mod go.mod
# COPY go.sum go.sum

RUN go mod download

COPY . .
RUN go build -v -i -o dist/color-svc ./cmd/color-svc/

# Run
FROM rytswd/toolkit-alpine:0.1.1

COPY --from=build /go/src/github.com/rytwsd/color-svc/dist/color-svc /usr/local/bin/

# ENTRYPOINT ["color-svc"]
