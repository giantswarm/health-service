FROM alpine:3.8

RUN mkdir -p /opt
ADD ./health-service /opt/health-service

EXPOSE 8000
ENTRYPOINT ["/opt/health-service"]
