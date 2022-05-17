FROM golang:1.18.2 as builder

ENV CGO_ENABLED=0

COPY go.* /src/
RUN cd /src && go get ./...

COPY . /src/
RUN cd /src && go build -o /frontend ./cmd/frontend
RUN cd /src && go build -o /proxy ./cmd/proxy

FROM gcr.io/distroless/static

COPY --from=builder /frontend /frontend
COPY --from=builder /proxy /proxy

