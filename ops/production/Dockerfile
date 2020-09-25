# Vuejs
FROM node:14.0.0-alpine3.10

COPY webpack /webpack
WORKDIR /webpack
RUN npm install
RUN npm run build

# Bee-go
FROM golang:1.13

COPY beego /go/src/backend
COPY --from=0 /webpack/dist/build.js /go/src/backend/static/js/build.js

WORKDIR /go/src/backend
ENV BEEGO_RUNMODE=prod
RUN go mod vendor
RUN go build main.go

EXPOSE 8080

ENTRYPOINT ./main
