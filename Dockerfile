FROM alpine:latest

ARG TARGETPLATFORM
COPY ${TARGETPLATFORM}/pingme /usr/bin/pingme

# Expose default webhook port
EXPOSE 8080

# Set default command to start webhook server
ENTRYPOINT ["/usr/bin/pingme"]
CMD ["serve"]
