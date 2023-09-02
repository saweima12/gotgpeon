devbot:
	go run ./cmd/peonbot -configPath local.yml

devjob:
	go run ./cmd/peonjob -configPath local.yml

devsetup:
	go run ./cmd/peonsetup -configPath local.yml

build:
	go build -o ./app/ ./cmd/...

build_img:
	docker build . -f ./Dockerfile --no-cache -t peonsuit
	yes | docker image prune --filter label=stage=builder 
	yes | docker image prune --filter label=stage=runtime 

dev-compose: build_img
	docker-compose up
