FROM golang:1.17 AS builder
LABEL bayt.cloud.image.authors="laith@bayt.cloud"
ARG GITHUB_ID
ENV GITHUB_ID=$GITHUB_ID
ARG MY_GITHUB_TOKEN
ENV MY_GITHUB_TOKEN=$MY_GITHUB_TOKEN
WORKDIR /app
USER $APP_USER
ADD src .
RUN --mount=type=secret,id=MY_GITHUB_TOKEN,required \
  export MY_GITHUB_TOKEN=$(cat /run/secrets/MY_GITHUB_TOKEN) && \
  git config \
  --global \
  url."https://${GITHUB_ID}:${MY_GITHUB_TOKEN}@github.com".insteadOf \
  "https://github.com"

ENV GOPRIVATE="github.com/laithrafid"
RUN go mod download
RUN go mod verify
RUN go build -o /userapi


FROM alpine:3.15.0 AS runner
ARG USERS_API_ADDRESS
ARG MYSQLDB_DRIVER
ARG MYSQLDB_SOURCE
ENV MYSQLDB_DRIVER=$MYSQLDB_DRIVER
ENV MYSQLDB_SOURCE=$MYSQLDB_SOURCE
ENV USERS_API_ADDRESS=$USERS_API_ADDRESS
WORKDIR /
COPY --from=builder /userapi /userapi
EXPOSE $USERS_API_ADDRESS
ENTRYPOINT ["/userapi"]
