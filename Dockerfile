FROM golang:1.15.5-alpine AS build
LABEL maintainer="rafaeldasilvacosta@hotmail.co.uk"
WORKDIR /usr/src

COPY go.* ./
RUN go mod download

COPY internal/app/hello internal/app/hello
COPY cmd/hello cmd/hello

# Statically compile the binary (resulting binary will not be linked to any C libraries)
ENV CGO_ENABLED=0
RUN go build -o /usr/bin cmd/hello/hello.go

FROM scratch
COPY --chown=1001:1001 --from=build /usr/bin /usr/bin
USER 1001
ENV SERVER_PORT 8080
ENTRYPOINT ["hello"]