FROM golang:1.9.7

RUN mkdir /app

ADD bin/imdbBackend /app/imdbBackend

WORKDIR /app

EXPOSE 2333

ENTRYPOINT /app/imdbBackend
