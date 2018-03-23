# golang slackbot (WIP)
*Mostly i followed this tutorial [Here](https://rsmitty.github.io/Slack-Bot/)*

The goal was an mvp, the other part was learning how to deploy on [GCP](https://cloud.google.com/)

# Prereqs

- golang (currently using 1.10.0. havent checked earlier)
- docker Version 17.12.0-ce-mac55 (23011)
- google cloud sdk
- Cloud Project ID from your Google Cloud Console [dashboard](https://console.cloud.google.com/home/dashboard)


# To Build
- run `export GOOGLE_PROJ_ID="yourProjectID"`
- run `make build` to generate a bin file for linux architecture from the testbot.go file
- run make docker-build to generate the docker container
- run `make docker-push` to push it to your docker registry
