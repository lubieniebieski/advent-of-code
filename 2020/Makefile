.PHONY: up start down destroy run advent build
FILES = $(wildcard *.rb)

build:
	docker-compose -f docker-compose.yml build $(c)
up:
	docker-compose -f docker-compose.yml up -d $(c)

run:
	docker-compose -f docker-compose.yml run --rm $(c)

advent:
	docker-compose -f docker-compose.yml run --rm --name=test script bundle exec rubocop . --auto-correct
	$(foreach file, $(FILES), docker-compose -f docker-compose.yml run --rm --name=${file} script ruby ${file};)

start:
	docker-compose -f docker-compose.yml start $(c)

down:
	docker-compose -f docker-compose.yml down $(c)

destroy:
	docker-compose -f docker-compose.yml down -v $(c)

stop:
	docker-compose -f docker-compose.yml stop $(c)
