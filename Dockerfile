# Container image that runs your code
FROM golang:1.14

# Copies your code file from your action repository to the filesystem path `/` of the container
COPY . /prepare_release

WORKDIR /prepare_release

RUN go build -o prepare_release

# Code file to execute when the docker container starts up (`entrypoint.sh`)
ENTRYPOINT ["/prepare_release/prepare_release"]
