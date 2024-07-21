FROM golang:1.22.5 as build

WORKDIR /go/src/app
ARG cmd="cmd/app/main.go"
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/app $cmd

FROM gcr.io/distroless/static-debian11

COPY --from=build /go/bin/app /
CMD ["/app"]