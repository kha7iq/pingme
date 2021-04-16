FROM alpine:latest  

ENTRYPOINT ["entrypoint.sh"]
COPY scripts/entrypoint.sh /usr/bin/entrypoint.sh
RUN chmod +x /usr/bin/entrypoint.sh
COPY pingme /usr/bin/pingme