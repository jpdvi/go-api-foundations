SHELL := '/bin/bash'
export DBHOST=localhost
export DBPORT=5432
export DBUSER=postgres
export DBPASSWORD=postgres
export SECRETKEY=boogaboogabooga42
export APPLICATIONPORT=8080
ifeq ($($@), $(test))
export DBNAME=exch_test
else
export DBNAME=exch_dev
endif

install:
	chmod +x ./scripts/install.sh
	chmod +x ./scripts/build.sh
	chmod +x ./scripts/clean.sh
	chmod +x ./scripts/run.sh
	./scripts/install.sh

build:
	./scripts/build.sh
run:
	./scripts/run.sh
migrate:
	go run migrations/migrate.go
test:   
	echo $@
	psql -c "DROP DATABASE IF EXISTS exch_test;"
	psql -c "CREATE DATABASE exch_test;"
	make migrate
	go test ./.../test
clean:
	./scripts/clean.sh
