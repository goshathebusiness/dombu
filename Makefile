DIR=${CURDIR}

.PHONY:up
up:
	$(MAKE) postgres-service-start SERVICE=user-management
	$(MAKE) postgres-service-start SERVICE=wallet

	docker-compose up -d janus-gateway

.PHONY:build
build:
	$(MAKE) vendor
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

.PHONY:postgres-service-start
postgres-service-start:
	docker-compose up -d ${SERVICE}-postgres
	./scripts/wait-for-postgres.sh ${SERVICE}-postgres
	$(MAKE) migrations-up SERVICE=${SERVICE}
	docker-compose up -d ${SERVICE}

.PHONY:service-rebuild
service-rebuild:
	docker-compose stop ${SERVICE}
	docker-compose build ${SERVICE}
	$(MAKE) service-start SERVICE=${SERVICE}

.PHONY:migrations-up
migrations-up:
	docker run --rm -v $(DIR)/pkg/${SERVICE}/repository/postgres/migrations:/migrations \
		--network lms_default migrate/migrate -path=/migrations/ \
		-database postgres://dev:12345@${SERVICE}-postgres:5432/${SERVICE}?sslmode=disable up

.PHONY:migrations-down
migrations-down:
	docker run --rm -v $(DIR)/pkg/${SERVICE}/repository/postgres/migrations:/migrations \
		--network lms_default migrate/migrate -path=/migrations/ \
		-database postgres://dev:12345@${SERVICE}-postgres:5432/${SERVICE}?sslmode=disable down -all

.PHONY:vendor
vendor:
	go mod vendor

.PHONY:go-lint
go-lint:
	docker run -t --rm -v $(DIR):/app -v ~/.cache/golangci-lint/v1.55.2:/root/.cache -w /app golangci/golangci-lint:v1.55.2 golangci-lint run -v
