# Bee-go
FROM golang:1.13

RUN GO111MODULE=on go get github.com/beego/bee
RUN export PATH=$PATH:$GOPATH/bin

RUN apt-get -q update && apt-get -qy install netcat
COPY ops/bin/wait-for.sh /usr/bin/wait-for.sh
RUN chmod +x /usr/bin/wait-for.sh

COPY webpack/dist/build.js /go/src/backend/static/js/build.js

WORKDIR /go/src/backend
RUN ls -la
# RUN go mod vendor
RUN bee run main.go

CMD /usr/bin/wait-for.sh backend-mysql:3306 -- ./main
