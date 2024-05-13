build:
	docker build -t faraway-server -f cmd/server/Dockerfile .
	docker build -t faraway-client -f cmd/client/Dockerfile .

run:
	docker-compose up -d --remove-orphans --build --force-recreate