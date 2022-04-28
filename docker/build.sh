go build -o ../bin/run.serv ../main.go
tar -Jcvf bookmark.tar.xz ../admin/ ../skin/ ../bin/
docker build -t scientistpun/bookmark:1.0.2 -t scientistpun/bookmark:latest .
rm bookmark.tar.xz
