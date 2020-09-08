.PHONY: build check-env

build:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o "testbot"
docker-build: check-env
	docker build -t us.gcr.io/${CLOUD_PROJ_ID}/testbot:test .
docker-push: check-env
	gcloud docker --  push us.gcr.io/${CLOUD_PROJ_ID}/testbot
check-env:
ifndef CLOUD_PROJ_ID
$(error CLOUD_PROJ_ID is undefined.  export CLOUD_PROJ_ID="your gcp project id")
endif
