# CI/CD Pipelines

![ci-cd-wall](/images/ci-cd-wall.png)

CI/CD is the automation layer that turns source changes into tested, packaged, and deployable software. In DevOps, it connects version control, testing, security, release management, and operations into one repeatable delivery flow.

## Core Meaning

CI means Continuous Integration: merge changes frequently and validate them automatically.

CD can mean Continuous Delivery or Continuous Deployment.

- Continuous Delivery keeps every change in a deployable state, with a controlled release step.
- Continuous Deployment pushes every passing change to production automatically.

Most teams should start with Continuous Delivery and move toward Continuous Deployment only when automation, tests, and rollback are mature.

## Why CI/CD Matters

- It detects problems early.
- It reduces merge and release risk.
- It standardizes builds and deployments.
- It gives traceability from commit to environment.
- It improves release speed without losing control.
- It supports rollback and incident response.

Without CI/CD, teams tend to accumulate manual steps, inconsistent results, and fragile releases.

## Pipeline Stages

A practical pipeline usually follows this shape:

1. Trigger from a push, pull request, tag, or schedule.
2. Restore dependencies.
3. Run lint and format checks.
4. Run unit and integration tests.
5. Build the artifact or container image.
6. Run security scans and policy checks.
7. Publish versioned outputs.
8. Deploy to an environment.
9. Run smoke or post-deploy validation.
10. Report results to the team.

Not every project needs every stage, but the pipeline should be able to support them.

## Inputs and Outputs

### Inputs

- Git source code.
- Build configuration.
- Secrets from a secure store.
- Environment variables.
- Release tags or pipeline parameters.

### Outputs

- Test reports.
- Build artifacts.
- Container images.
- Deployment manifests.
- Logs and metrics.
- Release notes or version records.

Treat outputs as part of the software supply chain. They should be traceable, repeatable, and immutable where possible.

## Artifacts

An artifact is the build output that gets promoted through environments. Examples include a JAR, package archive, binary, or Docker image.

Good artifacts are:

- Versioned.
- Reproducible.
- Traceable to a commit.
- Reused across environments instead of rebuilt differently each time.

Rebuilding for every environment is a common source of drift and hard-to-debug release issues.

## Environments

Most delivery models use multiple environments:

- Development for fast iteration.
- CI or test for automated validation.
- Staging for production-like checks.
- Production for live traffic.

General rule: lower environments are easier to change, higher environments are more controlled.

## Branching and Release Strategy

CI/CD works best when branch strategy and release strategy align.

### Feature Branches

Short-lived branches are common when pull requests are the main review mechanism and the main branch must stay stable.

### Trunk-Based Development

Frequent integration into a shared trunk works well when CI is fast, merge conflicts are kept small, and releases are frequent.

### Release Branches

Release branches are useful when you need stabilization, hotfix separation, or controlled release windows.

## Triggers

Common triggers include:

- Push to a branch.
- Pull request creation or update.
- Merge to main.
- Release tag creation.
- Manual approval.
- Scheduled execution.

Use triggers deliberately. Not every action should run the entire pipeline.

## Quality Gates

A quality gate is a rule that must pass before progression.

Common gates include:

- Linting and formatting.
- Test success.
- Coverage thresholds.
- Security scan results.
- Artifact signing or approval.
- Environment policy checks.

Quality gates protect the main branch and reduce the chance of broken code reaching users.

## Security in CI/CD

Security should be embedded in the pipeline.

Typical checks:

- Secret scanning.
- Dependency scanning.
- Static analysis.
- Container image scanning.
- Infrastructure as Code scanning.
- Deployment policy validation.

Best practices:

- Keep secrets out of source and pipeline files.
- Use a secret manager or CI secret store.
- Limit credential scope.
- Rotate credentials regularly.
- Sign artifacts when supply-chain integrity matters.

## Deployment Patterns

Common rollout patterns:

- Rolling deployment: update instances gradually.
- Blue-green deployment: switch traffic between two environments after validation.
- Canary deployment: expose a small slice of traffic first.
- Feature flags: deploy code while controlling runtime behavior.

These patterns reduce release risk and are often combined.

## Real-World Use Cases

Typical use cases:

- Web apps: PR checks, artifact build, staging deploy, production approval.
- APIs: container build, image scan, registry publish, manifest promotion.
- Infrastructure: plan on every change, apply after review and policy checks.
- Regulated systems: audit logging, retention, and explicit release approval.

## Example Pipeline Flow

1. Trigger on pull request.
2. Install dependencies.
3. Run lint, tests, and security scans.
4. Build and store the artifact.
5. Deploy to staging, run smoke tests, then promote after approval.

This keeps validation early and release control explicit.

## Minimal GitHub Actions Example

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
			- name: Set up runtime
				run: echo "configure runtime"
			- name: Install dependencies
				run: echo "install dependencies"
			- name: Lint
				run: echo "run lint"
			- name: Test
				run: echo "run tests"
```

This is intentionally minimal. Real pipelines replace the echo steps with actual commands for the stack being used.

## Best Practices

- Keep pipelines fast and deterministic.
- Promote the same artifact through environments.
- Keep build, test, and deploy concerns separate.
- Store pipeline definitions in version control.
- Protect production with checks, approvals, and observability.

## Common Anti-Patterns

- Rebuilding different artifacts for each environment.
- Storing secrets in plain text pipeline files.
- Mixing deployment logic into application logic.
- Letting pipelines become slow enough that teams bypass them.

## Testing Strategy

CI exists to make testing automatic and trustworthy.

- Unit tests: fast and isolated.
- Integration tests: service interaction and contracts.
- End-to-end tests: critical user paths.
- Smoke tests: post-deploy health checks.

Run the right test at the right stage.

## Observability and Feedback

CI/CD should provide useful feedback, not just pass or fail.

- Pipeline metrics: build duration, failure rate, deployment frequency, MTTR.
- Runtime metrics: error rate, latency, resource usage, business signals.

Pipeline health and application health both matter.

## Rollback and Recovery

Every serious deployment strategy needs a rollback plan.

- Redeploy the last known good artifact.
- Revert the commit and rebuild.
- Roll back the deployment manifest.
- Shift traffic back in blue-green or canary setups.

Rollback is only reliable when artifacts are versioned and deployment history is clear.

## Governance and Approvals

Some systems require human approval before release, especially when production has strict change control, compliance requirements, or limited release windows. Approval should complement automation, not replace it.

## Common Tools

CI/CD is tool-agnostic, but common choices include GitHub Actions, GitLab CI/CD, Jenkins, CircleCI, Azure DevOps Pipelines, Argo CD, Tekton, and Terraform Cloud or Enterprise. The discipline matters more than the tool: versioned pipelines, reproducible builds, tested promotion, and controlled releases.

## GitHub Actions Fundamentals

GitHub Actions is GitHub's built-in automation platform. It runs workflows in response to repository events such as pushes, pull requests, tags, releases, or schedules.

### Main Building Blocks

- Workflow: the full automation definition stored in `.github/workflows/*.yml`.
- Event: what triggers the workflow, such as `push`, `pull_request`, or `workflow_dispatch`.
- Job: a set of steps that runs on a runner.
- Step: one action or command inside a job.
- Runner: the machine that executes the job, such as `ubuntu-latest`.
- Action: a reusable unit of automation, often published by GitHub or the community.

### What You Should Learn First

Start with `push` and `pull_request`, then learn how to run shell commands, check out the repository, set up the runtime, read logs, and add secrets or artifacts later.

### Core Features To Know

- `workflow_dispatch` for manual runs.
- `jobs.<job_id>.runs-on` to choose the runner.
- `needs` to control job order.
- `secrets` and `permissions` to control access.
- `strategy.matrix` to test multiple versions or platforms.
- `actions/upload-artifact` and `actions/download-artifact` for build outputs.

## Real Practice 

### Good First Practice Workflow

Create a workflow that does simple quality checks on every pull request and push to `main`:

- Confirm Markdown files exist.
- Check that the version-control and CI/CD guides are present.
- Optionally count lines or search for required headings.
- Fail the workflow if expected files or sections are missing.

This gives you a real CI pipeline without needing an application runtime.

### Suggested Learning Path

1. Create a workflow file at `.github/workflows/ci.yml`.
2. Trigger it on `push` and `pull_request`.
3. Add a job on `ubuntu-latest` that checks out the repo and validates the markdown files.
4. Push the workflow, watch the Actions tab, and fix failures from the logs.

### Example Validation Workflow

```yaml
name: docs-ci

on:
	push:
		branches: [main]
	pull_request:

jobs:
	validate:
		runs-on: ubuntu-latest
		steps:
			- name: Checkout repository
				uses: actions/checkout@v4

			- name: Check required files
				run: |
					test -f README.md
					test -f 1-Version-Control/version-control.md
					test -f 2-CI&CD-Pipelines/ci-cd-pipelines.md

			- name: Verify markdown content exists
				run: |
					grep -q "# Version Control" 1-Version-Control/version-control.md
					grep -q "# CI/CD Pipelines" 2-CI&CD-Pipelines/ci-cd-pipelines.md

			- name: Show line counts
				run: |
					wc -l README.md 1-Version-Control/version-control.md 2-CI&CD-Pipelines/ci-cd-pipelines.md
```

This is a strong first workflow because it teaches event triggers, runners, shell steps, and failure handling without introducing unnecessary complexity.

### How To Think About It

Treat GitHub Actions as code that enforces delivery standards. A good workflow is not just a script that runs commands. It is a policy implementation for how changes enter the repository and move toward release.
