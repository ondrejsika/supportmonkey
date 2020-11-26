FROM golang as build
WORKDIR /build
COPY supportmonkey.go .
COPY go.mod .
COPY go.sum .
ENV CGO_ENABLED=0
RUN go build -a -ldflags \
  '-extldflags "-static"'

FROM debian:10 as certs
RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates

FROM scratch
COPY --from=certs /etc/ssl/certs /etc/ssl/certs
COPY --from=build /build/supportmonkey .
ENTRYPOINT [ "/supportmonkey" ]
CMD [ "-h" ]
