Service := carpark_backend

.PHONY: build

build:
	docker build --platform linux/amd64 --target release -t $(Service)-release . --no-cache
	docker build --platform linux/amd64 --target building -t $(Service)-building .
