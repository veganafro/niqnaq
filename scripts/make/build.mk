.PHONY: deepclean

deepclean: # Deep clean local docker
	docker container prune -f
	docker image prune -a -f
	docker volume prune -f

