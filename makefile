.PHONY: run db
run:
	go run main.go

# Start the Docker containers in the 'db' directory
dc:
	docker-compose up 
dc-up:
	docker-compose up -d
dc-build:
	docker build -t nxrfandi/kredit-plus:0.0.0.2 .

dc-push:
	docker push nxrfandi/kredit-plus:0.0.0.2
