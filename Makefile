build:
	go build -o ./bin/carservice ./cmd
run:
	./bin/carservice

docker-build:
	docker build -t carservice:latest -f Dockerfile .

docker-run:
	docker run -d -p 80:80 --rm --name carservice_container carservice