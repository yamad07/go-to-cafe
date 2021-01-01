migrate:
	docker-compose up -d mysql
	docker-compose run --rm dbmate migrate

run:
	docker-compose up -d mysql
	docker-compose up -d elasticsearch
	docker-compose up app
