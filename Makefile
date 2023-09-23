DIR=${CURDIR}

.PHONY:up
up:
	$(MAKE) service-start SERVICE=user-management
	$(MAKE) service-start SERVICE=wallet

	docker-compose up -d janus-gateway

.PHONY:build
build:
	docker-compose build

.PHONY:stop
stop:
	docker-compose stop

.PHONY:down
down:
	docker-compose down

.PHONY:service-start
service-start:
	docker-compose up -d ${SERVICE}

.PHONY:service-rebuild
service-rebuild:
	docker-compose stop ${SERVICE}
	docker-compose build ${SERVICE}
	$(MAKE) service-start SERVICE=${SERVICE}