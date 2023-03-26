LOCAL_BIN:=$(CURDIR)/bin
GOPATH?=$(HOME)/go

.PHONY: new-migration
new-migration:
ifeq ("$(wildcard $(LOCAL_BIN)/goose))","")
	go install github.com/pressly/goose/v3/cmd/goose && mv $(GOPATH)/bin/goose $(LOCAL_BIN)
endif
	cd migration && $(LOCAL_BIN)/goose create migration sql && $(LOCAL_BIN)/goose fix

.PHONY: postgres
postgres:
	docker run --name high_1 -p 5432:5432 -e POSTGRES_PASSWORD=root postgres:13.3

.PHONY: migrate
migrate:
ifeq ("$(wildcard $(LOCAL_BIN)/goose))","")
	go install github.com/pressly/goose/v3/cmd/goose && mv $(GOPATH)/bin/goose $(LOCAL_BIN)
endif
	cd  $(CURDIR)/migration && $(LOCAL_BIN)/goose postgres "postgres://postgres:root@localhost:5432/postgres?sslmode=disable"  up

