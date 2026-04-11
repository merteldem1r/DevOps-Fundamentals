# syntax=docker/dockerfile:1

# LAYER 1: build stage
FROM golang:1.25 AS build-stage 

# set the working directory in the container
WORKDIR /app

# copy go.mod and go.sum first to leverage Docker cache for dependencies
COPY go.mod go.sum ./
RUN go mod download

# copy the rest of the application source code
COPY . .

# build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping ./src/cmd/main.go

# LAYER 2: run tests
FROM build-stage AS run-test-stage
RUN go test -v ./...

# LAYER 3: release stage, tiny runtime image
FROM gcr.io/distroless/base-debian12:nonroot AS build-release-stage

# set the working directory in the container
WORKDIR /

# copy the built binary from the build stage
COPY --from=build-stage /docker-gs-ping /docker-gs-ping

# expose the application port
EXPOSE 8080

# set the entrypoint to run the application
ENTRYPOINT ["/docker-gs-ping"]