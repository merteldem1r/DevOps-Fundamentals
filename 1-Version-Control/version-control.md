# Version Control

Version control tracks changes to code and configuration over time. In DevOps, it enables collaboration, rollback, release traceability, and automation.

- Git: distributed version control system.
- GitHub: platform for hosting Git repos, pull requests, review, CI/CD integration, and governance.

## Why It Matters

- Preserves complete change history.
- Enables team collaboration without file overwrites.
- Makes rollback and incident recovery practical.
- Powers CI/CD and Infrastructure as Code workflows.
- Provides auditability for compliance and operations.

## Core Git Concepts

- Repository: project and its history.
- Commit: snapshot with hash, author, message, parent.
- Branch: isolated line of development.
- Merge: combines branch histories.
- Rebase: rewrites commit base for linear history.
- Remote: shared repo endpoint (usually `origin`).
- Pull Request: review and merge workflow on GitHub.

## Fast Workflow (Recommended)

1. Create/switch to a feature branch.
2. Make changes.
3. Stage and commit.
4. Push branch.
5. Open pull request.
6. Address review feedback.
7. Merge after checks pass.

## Essential Commands

### Start and Inspect

```bash
git init
git clone https://github.com/owner/repo.git
git status
git log --oneline --graph --decorate --all
```

### Stage and Commit

```bash
git add file.txt
git add .
git commit -m "Add concise version-control guide"
```

### Sync with Remote

```bash
git push origin main
git push -u origin feature/my-change
git fetch --all --prune
git pull
git pull --rebase
```

Use `pull --rebase` when you want a cleaner, linear local history.

## Branching and Checkout

### Create and Switch Branch

```bash
git switch -c feature/login-flow
# older equivalent
git checkout -b feature/login-flow
```

### Switch Existing Branch

```bash
git switch feature/login-flow
# older equivalent
git checkout feature/login-flow
```

`git switch` is clearer for branch operations. `git checkout` is older and overloaded (branch switching, commit checkout, file restore).

### List/Rename/Delete Branches

```bash
git branch
git branch -a
git branch -m old-name new-name
git branch -d feature/old-work
git branch -D feature/old-work
```

Use `-D` only when you intentionally force-delete.

## Inspecting Changes

```bash
git diff
git diff --staged
git show <commit-hash>
git show HEAD:file.txt
git blame file.txt
```

## Checkout a Commit Hash (Detached HEAD)

```bash
git checkout <commit-hash>
# modern equivalent
git switch --detach <commit-hash>
```

This is useful for reviewing old snapshots or debugging regressions.

If you want to keep new work from that commit:

```bash
git switch -c fix/from-old-commit
# older equivalent
git checkout -b fix/from-old-commit
```

Return to previous branch:

```bash
git switch -
```

## Undo and Rollback

### Local Undo

```bash
git restore --staged file.txt   # unstage
git restore file.txt            # discard working file changes
git reset --soft HEAD~1         # undo commit, keep changes
git reset --hard HEAD~1         # undo commit, discard changes
```

### Safe Undo for Shared Branches

```bash
git revert <commit-hash>
```

Creates a new commit that reverses an old commit without rewriting history.

### Roll Back 2-3 Pushed Commits

Option 1 (recommended on shared/protected branches):

```bash
git revert --no-edit HEAD~3..HEAD
git push origin main
```

Option 2 (history rewrite, team coordination required):

```bash
git reset --hard HEAD~3
git push --force-with-lease origin main
```

Prefer `--force-with-lease` over `--force`.

Preview before rollback:

```bash
git log --oneline --graph -n 10
```

## Stash (Temporary Save)

Use stash when work is not ready to commit but you must switch context.

```bash
git stash
git stash push -m "WIP: version-control notes"
git stash list
git stash apply
git stash apply stash@{0}
git stash pop
git stash drop stash@{0}
git stash clear
```

- `apply`: restore but keep stash.
- `pop`: restore and remove stash.

## Merge and Rebase

```bash
git switch main
git merge feature/login-flow

git switch feature/login-flow
git rebase main
```

- Merge preserves branch topology.
- Rebase creates linear history by rewriting commit ancestry.

### Conflict Resolution Flow

1. Open conflicted files.
2. Resolve markers.
3. Stage fixed files.
4. Continue merge/rebase.

## Remotes, Tags, and Releases

```bash
git remote -v
git remote add origin https://github.com/owner/repo.git
git remote rename origin upstream
git remote remove origin

git tag -a v1.0.0 -m "Release 1.0.0"
git tag
git push origin v1.0.0
git push origin --tags
```

Use annotated tags for production releases.

## .gitignore Essentials

Common ignores:

```gitignore
node_modules/
.env
dist/
.DS_Store
```

Do not commit secrets. If a secret is committed, rotate it and clean history appropriately.

## GitHub Practices for DevOps

- Keep `main` stable with branch protection.
- Require pull requests, reviews, and status checks.
- Keep branches short-lived and focused.
- Use clear commit messages and small PRs.
- Use tags/releases for deploy traceability.

## Compact Command Cheat Sheet

```bash
git init
git clone <repo-url>
git status
git add .
git commit -m "message"
git log --oneline --graph --decorate --all
git switch -c <branch>
git switch <branch>
git checkout <branch>
git checkout <commit-hash>
git diff
git diff --staged
git fetch --all --prune
git pull --rebase
git push
git merge <branch>
git rebase <branch>
git stash
git stash pop
git revert <commit>
git reset --hard HEAD~1
git push --force-with-lease
git tag -a v1.0.0 -m "Release"
```

## Common Mistakes to Avoid

- Large, mixed-purpose commits.
- Direct pushes to protected branches.
- Unsafe force-pushes without coordination.
- Ignoring conflicts for too long.
- Committing secrets or generated artifacts.

## Summary

Professional Git usage is about controlled change: small commits, branch-based collaboration, safe rollback, and clear history. Master these commands and workflows, and you will have a strong DevOps foundation.