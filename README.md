# DevOps Fundamentals

Practical learning repository for core DevOps topics with hands-on notes, examples, and CI pipeline practice.

## Repository Structure

### 1. Version Control

Focus: Git and GitHub fundamentals, workflows, rollback strategies, stash usage, and practical command usage.

- Folder: [1-Version-Control](1-Version-Control)

### 2. CI/CD Pipelines

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


