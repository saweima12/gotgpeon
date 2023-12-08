dev:
	go build -o ./build ./cmd/...
	docker-compose -f ./docker-compose-dev.yml up

cbuild:
	docker build . -f ./Dockerfile --no-cache -t peonsuit
	yes | docker image prune --filter label=stage=builder 
	yes | docker image prune --filter label=stage=runtime 

buildrun: cbuild
	docker-compose up
