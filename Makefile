#DEV

build-dev:
	docker build -t video-calling -f containers/images/Dockerfile . && \
	docker build -t turn -f containers/images/Dockerfile.turn .

clean-dev:
	docker-compose -f containers/compose/docker-compose.dev.yml down

run-dev:
	docker-compose -f containers/compose/docker-compose.dev.yml up