run:
	go run main.go

build:
	go build -o ./build/consumer main.go

rabbitmq:
	docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3