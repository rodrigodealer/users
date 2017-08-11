FROM alpine:edge

RUN apk add --update curl && \
    rm -rf /var/cache/apk/*

COPY users /opt/users

EXPOSE 8080
CMD ["/opt/users"]
