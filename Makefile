export CHAT_COMMIT ?= $(shell git rev-parse HEAD)

run:
	chmod a+x ./deploy.sh
	bash ./deploy.sh