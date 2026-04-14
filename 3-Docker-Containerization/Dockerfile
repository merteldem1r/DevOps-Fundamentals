# syntax=docker/dockerfile:1

# ---------- BUILD STAGE ----------
FROM golang:1.25 AS build-stage

WORKDIR /app

# cache dependencies
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# copy source
COPY . .

# build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/app ./src/cmd/main.go


# ---------- TEST STAGE ----------
FROM build-stage AS test-stage
RUN go test -v ./...


# ---------- FINAL STAGE ----------
FROM gcr.io/distroless/base-debian12:nonroot

WORKDIR /

COPY --from=build-stage /app/app /app

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/app"]