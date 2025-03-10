.PHONY: clean-docker

clean-docker:
	docker compose down --volumes --remove-orphans