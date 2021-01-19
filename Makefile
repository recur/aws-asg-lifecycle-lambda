SHELL := /bin/bash

menu:
	@perl -ne 'printf("%10s: %s\n","$$1","$$2") if m{^([\w+-]+):[^#]+#\s(.+)$$}' Makefile

build: # Build recur/aws-asg-lifecycle-lambda
	docker build -t recur/aws-asg-lifecycle-lambda .

push: # Push recur/aws-asg-lifecycle-lambda
	docker push recur/aws-asg-lifecycle-lambda

pull : # Pull recur/aws-asg-lifecycle-lambda
	docker pull recur/aws-asg-lifecycle-lambda

bash: # Run bash shell with recur/aws-asg-lifecycle-lambda
	docker run --rm -ti --entrypoint bash recur/aws-asg-lifecycle-lambda

clean:
	docker-compose down --remove-orphans

up:
	docker-compose up -d --remove-orphans

down:
	docker-compose rm -f -s

recreate:
	$(MAKE) clean
	$(MAKE) up

logs:
	docker-compose logs -f

pr:
	gh pr create --web

fmt:
	go fmt

test:
	go test

ci-go-test:
	/env.sh figlet -f /j/chunky.flf go
	/env.sh $(MAKE) test
