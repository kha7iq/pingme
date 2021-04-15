FROM alpine:latest AS builder
RUN apk upgrade -U -a && \
          apk upgrade && \
          apk add --update go gcc g++ git ca-certificates curl make
WORKDIR /app
COPY ./ /app
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags "-X main.version=$RELEASE_VERSION"

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/pingme /app/pingme
ENTRYPOINT ["/app/pingme"]  