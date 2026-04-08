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

### Rolling Deployment

Update instances gradually. Good for standard services and low-risk incremental rollout.

### Blue-Green Deployment

Maintain two environments and switch traffic after validation. Good for fast rollback and reduced downtime.

### Canary Deployment

Send a small portion of traffic to the new version first. Good for risk reduction and real-world validation.

### Feature Flags

Deploy code while keeping behavior hidden behind runtime toggles. Good for progressive release and separating deploy from release.

These patterns are often combined in mature delivery systems.

## Real-World Use Cases

### Web Application Team

A pull request runs lint, tests, and build validation. Merge to main publishes an artifact, deploys to staging, and promotes to production after approval.

### API Platform Team

Each merge builds and scans a container image, then publishes it to a registry. Deployment manifests are promoted through Git-based automation.

### Infrastructure Team

Terraform plan runs on every change. Apply happens only after review, policy checks, and controlled approval.

### Regulated Team

The pipeline includes audit logging, access control, retention, and explicit sign-off before production deployment.

## Example Pipeline Flow

1. Trigger on pull request.
2. Install dependencies.
3. Run lint and format checks.
4. Run unit tests.
5. Build the artifact.
6. Run security scans.
7. Store reports and outputs.
8. On merge to main, publish a release artifact.
9. Deploy to staging.
10. Run smoke tests.
11. Promote to production after approval.

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

- Keep pipelines fast enough for daily use.
- Fail early on obvious problems.
- Make stages deterministic and repeatable.
- Promote the same artifact through environments.
- Store pipeline definitions in version control.
- Separate build, test, and deploy concerns.
- Keep logs and reports visible.
- Use small, frequent changes instead of large risky batches.
- Protect production with checks, approvals, and observability.

## Common Anti-Patterns

- Rebuilding different artifacts for each environment.
- Running slow manual checks before every merge.
- Storing secrets in plain text pipeline files.
- Mixing deployment logic into application logic.
- Letting pipelines get so slow that teams bypass them.
- Relying on production-only validation.
- Treating CI as a build step and ignoring test quality.

## Testing Strategy

CI exists largely to make testing automatic and trustworthy.

- Unit tests should be fast and isolated.
- Integration tests should validate service interaction.
- End-to-end tests should cover critical user paths.
- Smoke tests should confirm the app is healthy after deployment.
- Flaky tests should be fixed or isolated quickly.

The goal is not to run every test everywhere. The goal is to run the right test at the right stage.

## Observability and Feedback

CI/CD should provide useful feedback, not just pass or fail.

Useful pipeline signals include:

- Build duration.
- Test duration.
- Failure rate by stage.
- Deployment frequency.
- Change failure rate.
- Mean time to recovery.

After deployment, watch:

- Error rate.
- Latency.
- Resource usage.
- Business-critical metrics.

Pipeline health and application health both matter.

## Rollback and Recovery

Every serious deployment strategy needs a rollback plan.

Common approaches:

- Redeploy the last known good artifact.
- Revert the commit and rebuild.
- Roll back the deployment manifest.
- Switch traffic back in blue-green or canary setups.

Rollback is only reliable when artifacts are versioned and deployment history is clear.

## Governance and Approvals

Some systems require human approval before release.

This is useful when production has strict change control, compliance requirements, or tightly managed release windows.

Approval should complement automation, not replace it.

## Common Tools

CI/CD is tool-agnostic, but these are widely used:

- GitHub Actions
- GitLab CI/CD
- Jenkins
- CircleCI
- Azure DevOps Pipelines
- Argo CD
- Tekton
- Terraform Cloud or Enterprise

The discipline matters more than the tool: versioned pipelines, reproducible builds, tested promotion, and controlled releases.

## Practical Checklist

Before introducing or reviewing a pipeline, confirm that:

- Source is in version control.
- Build commands are documented.
- Tests run in automation.
- Secrets are managed securely.
- Artifacts are versioned.
- Environments are defined.
- Rollback is possible.
- Logs and reports are accessible.

## Summary

CI/CD is the operating system of modern software delivery. CI gives fast validation, CD gives controlled promotion, and deployment automation turns code changes into reliable releases.

The best pipelines are not the most complicated ones. They are the ones that are fast, repeatable, secure, observable, and aligned with how the team actually ships software.
