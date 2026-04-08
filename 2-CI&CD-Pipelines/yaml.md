# YAML File Fundamentals

YAML is a human-readable data serialization format used to define structured configuration. In DevOps, it is widely used because it is easier to read and maintain than many alternatives for nested configuration.

YAML is not a programming language. It does not execute logic by itself. It describes data that tools interpret.

## What YAML Is

YAML originally meant "Yet Another Markup Language," and later became a recursive acronym: "YAML Ain't Markup Language."

Key idea:

- Use indentation and simple key-value structure to represent data.
- Keep configuration readable for humans and machine-parsable for tools.

![yaml-comparation](/images/yaml-comparing.jpg)

## Why YAML Is Popular in DevOps

Tools such as Docker Compose, Kubernetes, GitHub Actions, Ansible, and CI/CD platforms rely on YAML because:

- It is concise and readable for nested config.
- It maps naturally to objects, lists, strings, numbers, and booleans.
- It is easy to version in Git and review in pull requests.
- Teams can standardize infrastructure and pipeline definitions as code.

YAML became a practical standard for "configuration as code."

## Where You Will See YAML

Common DevOps use cases:

- CI/CD pipelines: workflow definitions and jobs.
- Container orchestration: Kubernetes manifests.
- Multi-container local/prod setups: Docker Compose.
- Automation tools: Ansible playbooks.
- Application and deployment config files.

## Core YAML Syntax

### 1. Key-Value Pairs

```yaml
name: devops-fundamentals
version: 1
enabled: true
```

### 2. Indentation Defines Structure

YAML uses spaces for hierarchy. Use consistent indentation (usually 2 spaces).

```yaml
app:
	name: web
	port: 8080
```

Do not mix tabs and spaces.

### 3. Lists

```yaml
environments:
	- dev
	- staging
	- production
```

List of objects:

```yaml
services:
	- name: api
		port: 8080
	- name: worker
		port: 9090
```

### 4. Nested Objects

```yaml
database:
	host: localhost
	port: 5432
	credentials:
		user: admin
		ssl: false
```

### 5. Strings and Quoting

Use unquoted strings when simple, quoted strings when needed.

```yaml
plain: hello
quoted: "hello:world"
single: "literal value"
```

Use quotes if values contain special characters (`:`, `#`, `{`, `}`) or could be misread.

### 6. Null, Boolean, Number

```yaml
retries: 3
debug: false
description: null
```

### 7. Comments

```yaml
# This is a comment
timeout: 30
```

## Multi-Line Values

YAML supports block scalars for multi-line text.

```yaml
message: |
	Build started.
	Running tests.
	Build finished.
```

`|` preserves newlines. `>` folds lines into spaces.

## Common YAML Pitfalls

- Incorrect indentation.
- Mixing tabs and spaces.
- Writing `key:value` instead of `key: value`.
- Unquoted special values being interpreted unexpectedly.
- Duplicate keys in the same object.

When YAML fails, the problem is often formatting, not logic.

## YAML in Docker Compose

Docker Compose files use YAML to define services, images, volumes, networks, and environment values.

Example:

```yaml
services:
	web:
		image: nginx:latest
		ports:
			- "8080:80"
	db:
		image: postgres:16
		environment:
			POSTGRES_PASSWORD: example
```

Why YAML works well here:

- Services are naturally hierarchical.
- Lists such as ports, volumes, and networks are easy to express.
- Config stays readable as stacks grow.

## YAML in Kubernetes

Kubernetes resources are defined with YAML manifests.

Example Deployment:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
	name: api
spec:
	replicas: 2
	selector:
		matchLabels:
			app: api
	template:
		metadata:
			labels:
				app: api
		spec:
			containers:
				- name: api
					image: myorg/api:1.0.0
					ports:
						- containerPort: 8080
```

Why YAML is preferred here:

- Kubernetes objects are deeply nested.
- Definitions are versioned as Infrastructure as Code.
- Teams can review and promote manifest changes through Git workflows.

## YAML in GitHub Actions

GitHub Actions workflows are YAML files in `.github/workflows/`.

Example:

```yaml
name: ci

on:
	pull_request:
	push:
		branches: [main]

jobs:
	test:
		runs-on: ubuntu-latest
		steps:
			- uses: actions/checkout@v4
			- name: Run checks
				run: echo "run tests here"
```

YAML helps express triggers, jobs, and steps in a clear hierarchy.

## YAML Anchors and Reuse (Intermediate)

YAML supports anchors and aliases to reuse blocks.

```yaml
defaults: &defaults
	retries: 3
	timeout: 30

job_a:
	<<: *defaults
	name: build

job_b:
	<<: *defaults
	name: test
```

Some platforms fully support this, some have partial behavior depending on parser/version. Test in your target tool.

## Validation and Linting

For professional usage, always validate YAML before merging:

- Use a YAML linter.
- Use platform-specific validation (Kubernetes schema checks, Actions workflow checks).
- Run checks in CI so formatting errors fail early.

Readable YAML is good; validated YAML is reliable.

## Best Practices

- Keep files focused and modular.
- Use consistent indentation (2 spaces is common).
- Keep naming predictable.
- Prefer explicit over clever shortcuts.
- Add comments only where intent is not obvious.
- Avoid storing secrets directly in YAML.
- Validate and lint in every pull request.

## Security Notes

YAML files often contain references to sensitive values. Do not hardcode secrets like API keys, tokens, or passwords in repository-tracked YAML files.

Instead, use:

- Secret stores.
- Environment-specific secret injection.
- CI/CD secret management.

## Quick Mental Model

Think of YAML as a structured tree:

- Objects are branches.
- Lists are repeated nodes.
- Leaf values are actual settings.

If your indentation is correct and your keys are clear, your YAML is usually understandable and maintainable.

## Summary

YAML is one of the most important foundational formats in DevOps. It is used to define pipelines, infrastructure, containers, and automation workflows in a readable and version-controlled way.

If you master YAML syntax, structure, and validation habits, you remove a major source of CI/CD and deployment errors early in your DevOps journey.
