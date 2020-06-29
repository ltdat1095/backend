# Vuejs
FROM node:14.0.0-alpine3.10

COPY webpack /webpack
WORKDIR /webpack
RUN npm install
RUN npm run build

# Bee-go
FROM golang:1.13

RUN apt-get -q update && apt-get -qy install netcat

COPY beego /go/src/backend
COPY --from=0 /webpack/dist/build.js /go/src/backend/static/js/build.js

COPY ops/bin/wait-for.sh /usr/bin/wait-for.sh
RUN chmod +x /usr/bin/wait-for.sh

WORKDIR /go/src/backend
RUN go mod vendor
RUN go build main.go
CMD ./backend
