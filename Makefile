bot:
	go run ./cmd/peonbot

job:
	go run ./cmd/peonjob

build:
	go build -o ./app/ ./cmd/...

build_img:
	docker build . -f ./Dockerfile --no-cache -t peonsuit

dev-compose:
	docker-compose -f ./docker/docker-compose.yml build --no-cache
	docker-compose -f ./docker/docker-compose.yml up 
