build:
	go build -o bin/imdbBackend && chmod 775 bin/imdbBackend

run:
	go run src/main.go

docker-build:
	docker build -t imdb_backend .

docker-run:
	docker run --rm -it -p 2333:2333 --net=host --name imdb-backend imdb_backend