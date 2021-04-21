
NAME=ghbackup
APP_VERSION=0.0.1

build:
	docker build -t $(NAME):$(APP_VERSION) .

start-console: build
	docker run -it \
		--rm \
		-v `pwd`:/go/src/github.com/fujisanmagazine/ghbackup \
		--name $(NAME) \
		$(NAME):$(APP_VERSION) /bin/bash
