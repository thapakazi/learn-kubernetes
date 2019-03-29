# Dockerfile References: https://docs.docker.com/engine/reference/builder/
FROM golang:1.12 AS build-env
LABEL maintainer="thapakazi"
# # Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/thapakazi/api-hostname

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .
# Download all the dependencies
# https://stackoverflow.com/questions/28031603/what-do-three-dots-mean-in-go-command-line-invocations
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...


# Build only
FROM alpine
EXPOSE 2048
WORKDIR /app
COPY --from=build-env /go/bin/api-hostname /app/bin
# Run the executable
ENTRYPOINT ./bin