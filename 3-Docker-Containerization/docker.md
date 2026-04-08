# Docker Fundamentals

Docker is a containerization platform that packages applications and their dependencies into lightweight, portable, reproducible units called containers. In DevOps, Docker is foundational because it solves the "it works on my machine" problem and powers modern deployment workflows.

## What Docker Solves

Without Docker, teams face:

- Environment inconsistency across machines.
- Dependency conflicts and version mismatches.
- Slow, fragile deployment processes.
- Difficulty reproducing bugs or scaling reliably.

Docker solves these by bundling code, runtime, libraries, and configuration into a single immutable artifact that runs identically everywhere: your laptop, CI, staging, production.

## Core Concepts

![docker-arch](/images/docker-architecture.webp)

### Image

An image is a static blueprint for a container. It contains application code, runtime, dependencies, environment variables, and build instructions. Images are immutable and versioned.

### Container

A container is a running instance of an image. It is isolated from the host and other containers, with its own filesystem, processes, and network namespace. Multiple containers can run from the same image.

### Dockerfile

A Dockerfile is a text file with instructions to build an image. Each instruction creates a layer, and layers are cached for speed.

### Registry

A registry is a remote repository for images. Docker Hub is the default public registry. Teams often run private registries (Amazon ECR, Azure Container Registry, Harbor, etc.).

### Layer

Docker images are built in layers. Each Dockerfile instruction (FROM, RUN, COPY, etc.) creates a new layer. Layers are cached and reused, speeding up builds.

## How Docker Works

```text
Your project files + Dockerfile
        ↓
docker build
        ↓
Docker CLI sends request
        ↓
Docker daemon reads Dockerfile
        ↓
Pull base image if needed
        ↓
Execute instructions (COPY, RUN, etc.)
        ↓
Create image layers
        ↓
Store final image locally
        ↓
docker run
        ↓
Docker creates container from image
        ↓
Adds writable layer + isolation + network
        ↓
Starts main process
        ↓
Your app is now running inside the container
```

1. Write a Dockerfile.
2. Build the image: docker build -t myapp:1.0 .
3. Run a container: docker run myapp:1.0
4. Push to registry: docker push registry/myapp:1.0
5. Pull and run elsewhere: docker pull registry/myapp:1.0 && docker run registry/myapp:1.0

The same image runs identically in all environments.

## Dockerfile Fundamentals

A basic Dockerfile structure:

```dockerfile
# Start from a base image
FROM python:3.11-slim

# Set working directory
WORKDIR /app

# Copy files
COPY requirements.txt .

# Install dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Copy app code
COPY . .

# Expose port
EXPOSE 8000

# Define entrypoint
CMD ["python", "app.py"]
```

### Common Instructions

- FROM: base image to build on.
- WORKDIR: set the working directory inside container.
- COPY: copy files from host to container.
- RUN: execute commands during build (e.g., install packages).
- ENV: set environment variables.
- EXPOSE: declare which ports the app listens on (informational).
- CMD: default command when container starts.
- ENTRYPOINT: configure container as executable.
- ARG: build-time variables.

### Dockerfile Best Practices

- Use specific base image tags (python:3.11-slim, not python:latest).
- Minimize layers to reduce image size.
- Order instructions so frequently-changed layers come late.
- Use .dockerignore to exclude unnecessary files.
- Keep base images minimal (slim, alpine flavors).
- Use multi-stage builds for smaller final images.
- Never RUN as root unless required.

## Multi-Stage Builds

Multi-stage builds reduce final image size by using intermediate stages.

```dockerfile
# Build stage
FROM golang:1.21 AS builder
WORKDIR /build
COPY . .
RUN go build -o app .

# Runtime stage
FROM alpine:3.20
WORKDIR /app
COPY --from=builder /build/app .
EXPOSE 8080
CMD ["./app"]
```

The final image contains only the compiled binary and minimal runtime, not build tools.

## Docker Compose

Docker Compose defines multi-container applications in a YAML file.

```yaml
version: "3.8"

services:
  web:
    build: .
    ports:
      - "8000:8000"
    environment:
      DATABASE_URL: postgresql://db:5432/myapp
    depends_on:
      - db

  db:
    image: postgres:16
    environment:
      POSTGRES_PASSWORD: secret
    volumes:
      - db_data:/var/lib/postgresql/data

volumes:
  db_data:
```

Run with docker-compose up. Compose resolves service names (web, db) as hostnames inside containers.

## Common Docker Commands

### Build

```bash
docker build -t myapp:1.0 .
```

### Run

```bash
docker run -d -p 8000:8000 --name myapp myapp:1.0
```

- -d: detached (background)
- -p: port mapping (host:container)
- --name: container name

### View Running Containers

```bash
docker ps
```

### View Image Layers

```bash
docker history myapp:1.0
docker inspect myapp:1.0
```

### Push to Registry

```bash
docker tag myapp:1.0 registry.example.com/myapp:1.0
docker push registry.example.com/myapp:1.0
```

### Pull from Registry

```bash
docker pull registry.example.com/myapp:1.0
```

### Logs and Exec

```bash
docker logs myapp
docker exec -it myapp bash
```

## Container Networking

By default, containers are isolated. To communicate:

- Port mapping (-p) exposes container port on host.
- Docker networks let containers discover each other by name.
- Docker Compose creates networks automatically and resolves service names as hostnames.

## Volumes and Persistence

Containers are ephemeral. To persist data:

- Volumes: Docker-managed storage. Survive container deletion.
- Bind mounts: host directory mounted inside container. Good for development.
- tmpfs: in-memory storage.

```bash
docker run -v myvolume:/data myapp:1.0
docker run -v /host/path:/container/path myapp:1.0
```

## Image Optimization

### Minimize Layer Count

Each RUN, COPY, ADD instruction creates a layer. Combine related commands:

```dockerfile
RUN apt-get update && \
    apt-get install -y curl && \
    apt-get clean
```

### Use Alpine Base Images

Alpine is ~5MB vs 200MB+ for full Linux. Trade-off: less tooling, but smaller final product.

```dockerfile
FROM alpine:3.20
```

### Leverage Build Cache

Dockerfile instructions are cached by layer. Put frequently-changing instructions late.

```dockerfile
COPY requirements.txt .
RUN pip install -r requirements.txt
COPY . .
```

If you change app code, only the last layer rebuilds.

## Real-World Use Cases

### Web Application

Dockerfile builds a web service, pushed to registry. CI/CD pipeline pulls image and runs it in staging and production.

### Database Migration Job

Containerize migration tools and database schema. Run as one-off container on deployment.

### Microservices

Each service has its own Dockerfile/image. Docker Compose or Kubernetes orchestrates multiple containers.

### CI/CD Pipeline Environment

CI system pulls a Docker image with pre-installed tools (Go, Python, Node) to run tests.

## Docker and DevOps

Docker is core to modern DevOps because:

- Infrastructure as Code: Dockerfile is versioned like application code.
- Reproducibility: same image runs everywhere.
- Immutability: images do not change after push, reducing surprises.
- Scalability: orchestrate many containers uniformly (Kubernetes, Swarm).
- Supply chain: images are signed/scanned for security.

## Security Considerations

- Scan images for vulnerabilities: docker scout cves myimage:1.0
- Use minimal base images to reduce attack surface.
- Run containers as non-root unless required.
- Never store secrets in Dockerfile or images.
- Use secret management: environment variables, secret stores.
- Sign images for production integrity.

## Common Mistakes

- Using latest tag: unpredictable upgrades.
- Large image sizes: slow pulls, expensive storage.
- Running as root: security risk.
- Storing secrets in layers: baked into image.
- No health checks: orchestrators cannot detect failed containers.
- Not versioning images: cannot reproduce problems.

## Practical Checklist

Before using a Docker image in production:

- Image is versioned (not latest).
- Image is scanning for vulnerabilities.
- Dockerfile uses specific base image versions.
- Multi-stage builds applied if applicable.
- Image size is reasonable.
- Layers are optimized.
- Container runs as non-root.
- Logs and health checks are configured.
