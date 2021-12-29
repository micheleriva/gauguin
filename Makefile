DC=docker-compose

container-dev:
	$(DC) -f docker-compose-develop.yml up --build

container-prod:
	$(DC) up