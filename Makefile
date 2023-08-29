bot:
	go run ./cmd/peonbot

job:
	go run ./cmd/peonjob

build:
	go build -o ./app/ ./cmd/...

build_img:
	docker build . -f ./docker/Dockerfile --no-cache -t peon


