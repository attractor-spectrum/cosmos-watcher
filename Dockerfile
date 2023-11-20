FROM bitnami/golang:1.18-debian-10 as build

ARG GITHUB_USERNAME
ARG GITHUB_TOKEN
ARG GITHUB_REPO_PATH

RUN git config --global \
    url."https://${GITHUB_USERNAME}:${GITHUB_TOKEN}@github.com/${GITHUB_REPO_PATH}".insteadOf \
    "https://github.com/${GITHUB_REPO_PATH}"
    
WORKDIR /app

COPY . /app

RUN apt-get update && apt-get install -y make gcc gawk bison libc-dev

RUN go build -o watcher ./cmd/watcher/main.go

RUN ls -la /app

FROM ubuntu:latest as production

RUN apt-get update && apt-get install -y curl jq coreutils dos2unix

COPY --from=build /app/libwasmvm.so /usr/lib/libwasmvm.so
COPY --from=build /app/libwasmvm.x86_64.so /usr/lib/libwasmvm.x86_64.so
COPY --from=build /app/watcher  /app/watcher
COPY --from=build /app/scripts/run.sh /run.sh

RUN dos2unix /run.sh
CMD ["/run.sh"]