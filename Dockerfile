FROM alpine:edge

COPY users /opt/users

EXPOSE 8080
CMD ["/opt/users"]
