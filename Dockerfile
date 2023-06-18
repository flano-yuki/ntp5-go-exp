FROM golang:1.20-alpine as build

WORKDIR /app
COPY ./ ./
RUN go mod download
RUN go build -o /ntp5-go-exp ./app/ntpv5-cli/main.go  

FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /ntp5-go-exp /ntp5-go-exp
USER nonroot:nonroot
ENTRYPOINT ["/ntp5-go-exp"]
