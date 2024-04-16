FROM golang:1.22 AS build-stage

WORKDIR /go/src
ENV PATH="/go/src:${PATH}"

COPY . ./

RUN go mod download
RUN go mod tidy

ENV CGO_ENABLED 1
ENV GOOS=linux

RUN \
  --mount=target=. \
  --mount=target=/root/.cache,type=cache \
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build \
  -ldflags "-s -d -w" \
  -o /FasfoodApp cmd/main.go

FROM scratch

WORKDIR /app

COPY --from=build-stage /FasfoodApp /FasfoodApp

EXPOSE 3210

ENTRYPOINT ["/FasfoodApp"]
