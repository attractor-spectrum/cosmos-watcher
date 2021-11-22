FROM bitnami/golang:1.16-debian-10 as build

WORKDIR /app

COPY . /app

RUN apt-get update && apt-get install -y make gcc gawk bison libc-dev

RUN go build -o watcher ./cmd/watcher/main.go

RUN ls -la /app

FROM ubuntu:latest as production

RUN apt-get update && apt-get install -y curl jq

#COPY --from=build /app/libwasmvm.so /usr/lib/libwasmvm.so
COPY --from=build /app/watcher  /app/watcher
COPY --from=build /app/scripts/run.sh /run.sh

CMD ["/run.sh"]
