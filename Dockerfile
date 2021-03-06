FROM golang:1.13.11-alpine3.10 AS golang-build
RUN mkdir -p /go/src/github.com/AirHelp/zendesk-mock
WORKDIR /go/src/github.com/AirHelp/zendesk-mock
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o zendesk-mock .

FROM scratch
COPY --from=golang-build /go/src/github.com/AirHelp/zendesk-mock/zendesk-mock .
CMD ["/zendesk-mock"]
