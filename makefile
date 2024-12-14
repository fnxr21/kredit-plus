.PHONY: run db
run:
	go run main.go

# Start the Docker containers in the 'db' directory
dc:
	docker-compose up 
dc-d:
	docker-compose up -d