# Vuejs
FROM node:alpine

WORKDIR /vuejs
RUN npm install
CMD npm run build

# Bee-go
FROM golang:1.13

RUN go get github.com/astaxie/beego
RUN go get github.com/beego/bee
RUN go get gopkg.in/mgo.v2
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/go-sql-driver/mysql
RUN go get golang.org/x/crypto/bcrypt

WORKDIR /go/src/backend
CMD bee run
