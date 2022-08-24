# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.18-buster AS build

WORKDIR /usr/src/app

# RUN apt-get update && apt-get install -y ca-certificates git-core ssh

ARG BINARY_NAME
ARG GIT_SSH_KEY_PATH

# Uncomment lines below to allow for private repos
# # Setup for ssh onto github
# RUN mkdir -p /root/.ssh
# ADD ${GIT_SSH_KEY_PATH} /root/.ssh/id_rsa
# RUN chmod 700 /root/.ssh/id_rsa
# RUN echo "Host github.com\n\tStrictHostKeyChecking no\n" >> /root/.ssh/config
# RUN git config --global --add url."git@github.com:".insteadOf "https://github.com/"

# Download named modules
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app  ./cmd/${BINARY_NAME}/main.go

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /usr/local/bin/app  /usr/local/bin/app

USER nonroot:nonroot

ENTRYPOINT [ "app" ] 
