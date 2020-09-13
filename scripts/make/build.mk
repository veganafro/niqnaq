.PHONY: build clean deepclean logs proto start status stop

build: # Build docker image
	docker-compose build

coldstart: # Build and start docker images
	docker-compose up --detach --build

clean:stop # Stop docker containers, clean data and workspace
	docker-compose down -v --remove-orphans

deepclean: # Deep clean local docker
	docker container prune -f
	docker image prune -a -f
	docker volume prune -f

logs: # Get logs of containers
	docker-compose logs -f

start: # Start docker containers
	docker-compose up -d

status: # Get status of containers
	docker-compose ps

stop: # Stop docker containers
	docker-compose stop

