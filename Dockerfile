FROM golang:1.5-onbuild

COPY ./netcore /netcore

EXPOSE 53 67
ENTRYPOINT ["/netcore"]
