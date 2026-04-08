# DevOps Fundamentals

This repository is for practicing the foundations of DevOps in a practical and structured way.

## What Is DevOps?

DevOps is a way of working where Development and Operations collaborate to deliver software faster, safer, and more reliably.

It combines:

- Culture: shared ownership, collaboration, and fast feedback
- Practices: automation, continuous integration, continuous delivery, monitoring
- Tools: platforms that support build, test, deploy, observe, and secure systems

The main goal is to shorten the software delivery lifecycle while improving quality and stability.

## Core DevOps Stages

DevOps is often represented as an iterative lifecycle:

1. Plan

- Define requirements, priorities, and architecture.
- Output: user stories, tasks, design decisions.

2. Code

- Write application and infrastructure code.
- Output: source code in version control (usually Git).

3. Build

- Compile/package code and create artifacts (for example, Docker images).
- Output: versioned build artifacts.

4. Test

- Run automated tests (unit, integration, end-to-end, security checks).
- Output: quality and security feedback.

5. Release

- Approve and prepare deployment changes.
- Output: release candidates, changelogs, release notes.

6. Deploy

- Move changes to environments (staging, production), often through CI/CD.
- Output: running application updates.

7. Operate

- Keep services healthy, scalable, and cost-efficient.
- Output: reliable runtime operations.

8. Monitor

- Track metrics, logs, and traces to detect issues and improve continuously.
- Output: alerts, dashboards, and performance insights.

## Common DevOps Technologies

### Version Control

- Git
- GitHub, GitLab, Bitbucket

### CI/CD

- GitHub Actions
- GitLab CI/CD
- Jenkins
- CircleCI
- Argo CD (GitOps deployments)

### Containers and Orchestration

- Docker
- Kubernetes
- Helm

### Infrastructure as Code (IaC)

- Terraform
- Ansible
- Pulumi
- CloudFormation

### Cloud Platforms

- AWS
- Azure
- Google Cloud

### Monitoring and Observability

- Prometheus + Grafana
- ELK/EFK Stack (Elasticsearch, Logstash/Fluentd, Kibana)
- Datadog
- New Relic
- OpenTelemetry

### Security (DevSecOps)

- Snyk
- Trivy
- SonarQube
- OWASP ZAP
- HashiCorp Vault

## Key DevOps Principles

- Automate repetitive work
- Keep deployments small and frequent
- Shift testing and security left
- Use Infrastructure as Code
- Monitor everything that matters
- Learn from incidents and improve continuously
