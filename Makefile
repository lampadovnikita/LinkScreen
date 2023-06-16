.PHONY:

build:
		go build -o ./.bin/bot cmd/bot/main.go

run:	build
		./.bin/bot

service-update: build
	systemctl restart link-screen.service