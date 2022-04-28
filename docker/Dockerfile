FROM ubuntu:20.04
LABEL MAINTAINER="scientistpun <scientistpun@gmail.com>"
LABEL version="1.0.0"

ENV DIR=/home/bookmark
ENV USERNAME=root
ENV PASSWORD=admin123
ENV PORT=8080

COPY sources.list /etc/apt/sources.list
RUN apt-get clean && apt-get update
RUN apt-get install -y xz-utils

RUN mkdir $DIR
COPY bookmark.tar.xz $DIR/bookmark.tar.xz
RUN tar -C $DIR -xvf $DIR/bookmark.tar.xz
RUN mkdir $DIR/data
#RUN rm $DIR/bookmark.tar.xz

WORKDIR $DIR

VOLUME ["${DIR}/skin", "${DIR}/data"]
# RUN tar -k -C $DIR -xvf $DIR/bookmark.tar.xz
ENTRYPOINT $DIR/bin/run.serv