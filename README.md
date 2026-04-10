# DevOps Fundamentals

Practical learning repository for core DevOps topics with hands-on notes, examples, and CI pipeline practice.

**Search Resources**: [Build and Deploy a Production API](https://www.youtube.com/watch?v=H5FAxTBuNM8&list=WL&index=18&t=6724s) | [DevOps Tutorial - GeekForGeeks](https://www.geeksforgeeks.org/devops/devops-tutorial/) | [GitHub Actions](https://docs.github.com/en/actions/tutorials) | [Containerization using Docker - GeekForGeeks](https://www.geeksforgeeks.org/blogs/containerization-using-docker/) | [Docker - Golang Tutorial](https://docs.docker.com/guides/golang/)

## Project Sections

### 1. Version Control

![git](/images/version-control-branching.png)

Version control tracks changes to code and configuration over time. In DevOps, it enables collaboration, rollback, release traceability, and automation.

Focus: Git and GitHub fundamentals, workflows, rollback strategies, stash usage, and practical command usage.

- Git: distributed version control system.
- GitHub: platform for hosting Git repos, pull requests, review, CI/CD integration, and governance.

- Folder: [1-Version-Control](1-Version-Control)

### 2. CI/CD Pipelines

![ci/cd](/images/ci-cd-wall.png)

CI/CD is the automation layer that turns source changes into tested, packaged, and deployable software. In DevOps, it connects version control, testing, security, release management, and operations into one repeatable delivery flow.

Focus: CI/CD foundations, GitHub Actions practice, YAML syntax, and a Go-based pipeline testing setup.

- Folder: [2-CI-CD-Pipelines](2-CI-CD-Pipelines)

Simple folder structure:

```text
2-CI-CD-Pipelines/
|-- go.mod                       # Go module definition
|-- go.sum                       # Locked dependency checksums
|-- directive/
|   |-- ci-cd-pipelines.md       # CI/CD theory and workflow notes
|   |-- yaml.md                  # YAML fundamentals
|   |-- example.yaml             # YAML syntax and workflow
|-- src/
|   |-- main.go                  # entry
|   |-- internal/
|   |   |-- tests/
|   |       |-- uuid_test.go     # simple uuid test
|   |-- utils/                   # utilities
```

### CI Workflow

This repository includes a GitHub Actions pipeline used to validate and test CI behavior.

- Workflow file: [.github/workflows/pipeline.yaml](.github/workflows/pipeline.yaml)

### 3. Docker - Containerization

![docker](/images/docker-architecture.webp)

Docker is a containerization platform that packages applications and their dependencies into lightweight, portable, reproducible units called containers. In DevOps, Docker is foundational because it solves the "it works on my machine" problem and powers modern deployment workflows.

Focus: Docker fundamentals, containerization, Dockerfile syntax, image optimization, and Docker Compose for multi-container setups.

- Folder: [3-Docker-Container](3-Docker-Containerization)
