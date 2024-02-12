FROM golang

WORKDIR /pacgobot

COPY . .

ARG BOT_TOKEN=xxxx
ENV BOT_TOKEN=${BOT_TOKEN}

ARG REDIS_HOST=xxxx
ENV REDIS_HOST=${REDIS_HOST}

RUN go build 

ENTRYPOINT [ "./pacgobot" ]

