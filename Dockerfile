FROM alpine:latest
ENTRYPOINT ["/usr/bin/pingme"]
COPY pingme /usr/bin/pingme