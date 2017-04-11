FROM golang:1.8
RUN go get github.com/astaxie/beego && go get github.com/beego/bee && go get github.com/go-sql-driver/mysql && go get github.com/astaxie/beego/orm
RUN go get github.com/lisijie/webcron
ARG workdir
ENV workdir=${workdir}
WORKDIR ${workdir}