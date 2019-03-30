FROM golang:alpine AS build-env
LABEL maintainer="thapakazi"
WORKDIR /src
ADD . /src
RUN cd /src &&  go build -o api-hostname

FROM alpine
EXPOSE 2048
WORKDIR /app
COPY --from=build-env /src/api-hostname /app/
ENTRYPOINT ./api-hostname