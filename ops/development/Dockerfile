# Bee-go
FROM golang:1.13

COPY beego /go/src/backend
COPY webpack/dist/build.js /go/src/backend/static/js/build.js

# Wait for mysql initialization
RUN apt-get -q update && apt-get -qy install netcat
COPY ops/bin/wait-for.sh /usr/bin/wait-for.sh
RUN chmod +x /usr/bin/wait-for.sh

# Install bee client
RUN GO111MODULE=on go get github.com/beego/bee
RUN export PATH=$PATH:$GOPATH/bin

WORKDIR /go/src/backend
CMD /usr/bin/wait-for.sh backend-mysql:3306 -- bee run
