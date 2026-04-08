# Version Control

Version control is the discipline of tracking changes to code, configuration, and documentation over time. In DevOps, it is not just a convenience. It is the foundation for collaboration, automation, traceability, rollback, and reliable delivery.

Git is the most widely used version control system. GitHub is a platform built around Git that adds collaboration, code review, issue tracking, automation, security, and repository management.

## Why Version Control Matters

Version control gives teams the ability to:

- Track every change with author, timestamp, and history.
- Collaborate without overwriting each other's work.
- Compare versions and understand what changed.
- Revert mistakes quickly and safely.
- Create repeatable release processes.
- Support CI/CD pipelines and infrastructure as code.
- Audit decisions and reproduce old states when needed.

Without version control, software delivery becomes risky, opaque, and hard to recover when something breaks.

## Core Concepts

### Repository

A repository is a project that stores files and change history. A repository can be local on your machine, remote on a platform like GitHub, or both.

### Commit

A commit is a snapshot of your project at a point in time. Each commit has:

- A unique hash.
- An author.
- A message describing the change.
- A parent commit, which links history together.

Good commits are small, focused, and easy to understand.

### Working Tree, Staging Area, and History

Git works in three main areas:

- Working tree: the files you are editing.
- Staging area: changes prepared for the next commit.
- Repository history: committed snapshots stored in Git.

This model lets you choose exactly what goes into each commit.

### Branch

A branch is a movable pointer to a line of work. Branches let you develop features, fix bugs, or experiment without disturbing the main code line.

### Merge

Merge combines changes from one branch into another. If changes touch the same lines, Git may require conflict resolution.

### Remote

A remote is a shared repository, usually hosted on GitHub or another Git service. Common remote names include `origin`.

### Pull Request

A pull request is a review request for merging one branch into another. It is a collaboration and quality control mechanism, not just a merge button.

## Git vs GitHub

Git and GitHub are related but not the same.

- Git is the version control engine.
- GitHub is a hosting and collaboration platform for Git repositories.

Git handles commits, branches, merges, and history. GitHub adds pull requests, issues, code reviews, Actions, releases, access controls, and repository insights.

## Git Workflow Model

A typical Git workflow looks like this:

1. Clone or initialize a repository.
2. Create a branch for a change.
3. Modify files in the working tree.
4. Stage the relevant changes.
5. Commit with a clear message.
6. Push the branch to a remote repository.
7. Open a pull request.
8. Review, update, and merge.

This workflow encourages smaller changes, better reviewability, and safer releases.

## Getting Started Commands

### Initialize a New Repository

```bash
git init
```

Creates a new Git repository in the current directory.

### Clone an Existing Repository

```bash
git clone https://github.com/owner/repo.git
```

Downloads a remote repository and sets up the `origin` remote.

### Check Repository Status

```bash
git status
```

Shows modified, staged, and untracked files.

### View History

```bash
git log
git log --oneline --graph --decorate --all
```

Use the second form for a compact visual history of branches and merges.

## Daily Git Commands

### Stage Changes

```bash
git add file.txt
git add .
git add -A
```

- `git add file.txt` stages one file.
- `git add .` stages changes in the current directory.
- `git add -A` stages all tracked and untracked changes across the repository.

### Commit Changes

```bash
git commit -m "Add version control guide"
```

Write commit messages that describe the purpose of the change, not just the files touched.

### Push Commits

```bash
git push origin main
git push -u origin feature/version-control-notes
```

The `-u` flag sets the upstream branch so future pushes and pulls are simpler.

### Pull Remote Changes

```bash
git pull
git pull --rebase
```

`git pull` fetches and merges. `git pull --rebase` fetches and replays your commits on top of the updated remote branch.

## Branching Commands

### Create a Branch

```bash
git branch feature/login-flow
```

### Switch Branches

```bash
git switch feature/login-flow
git checkout feature/login-flow
```

Use `git switch` for branch movement in modern Git. `git checkout` is the older, more overloaded command that can also switch branches, restore files, and check out commits. In practice, `git switch` is clearer for branch changes, while `git checkout` still appears in older tutorials and in repositories that have not moved to the newer workflow.

### Create and Switch in One Step

```bash
git switch -c feature/login-flow
git checkout -b feature/login-flow
```

Use `git switch -c` as the modern equivalent of `git checkout -b`.

### List Branches

```bash
git branch
git branch -a
```

`-a` shows both local and remote-tracking branches.

### Rename a Branch

```bash
git branch -m old-name new-name
```

### Delete a Branch

```bash
git branch -d feature/old-work
git branch -D feature/old-work
```

Use `-d` when the branch is already merged. Use `-D` only when you are sure you want to force deletion.

## Comparing and Inspecting Changes

### Show Modified Files and Diffs

```bash
git diff
git diff --staged
```

- `git diff` shows unstaged changes.
- `git diff --staged` shows staged changes ready to commit.

### Inspect a File at a Commit

```bash
git show HEAD:file.txt
```

### View a Commit

```bash
git show <commit-hash>
```

### See Who Changed a Line

```bash
git blame file.txt
```

This is useful for investigating when and why a line changed.

## Undoing Changes Safely

### Unstage a File

```bash
git restore --staged file.txt
```

### Discard Working Tree Changes

```bash
git restore file.txt
```

### Undo the Last Commit but Keep Changes

```bash
git reset --soft HEAD~1
```

### Undo the Last Commit and Discard Changes

```bash
git reset --hard HEAD~1
```

Be careful with `--hard`; it removes local work that has not been saved elsewhere.

### Revert a Commit Safely

```bash
git revert <commit-hash>
```

`git revert` creates a new commit that undoes an earlier commit. It is the preferred way to undo changes on shared branches.

## Stashing Changes

`git stash` temporarily saves your uncommitted changes and restores a clean working tree. It is useful when you need to switch branches, pull updates, or handle an urgent fix without losing work that is not ready to commit.

Common use cases:

- You started a feature but need to switch to another branch quickly.
- You want to pull the latest changes before your work is ready.
- You need a clean tree for testing, debugging, or emergency fixes.

### Save Current Work

```bash
git stash
git stash push -m "WIP on version-control notes"
```

### View Stashes

```bash
git stash list
```

### Reapply Stashed Changes

```bash
git stash apply
git stash apply stash@{0}
```

`git stash apply` restores the changes but keeps the stash entry.

### Reapply and Remove the Stash

```bash
git stash pop
```

Use `pop` when you want to restore the changes and delete the stash entry in one step.

### Remove a Stash Manually

```bash
git stash drop stash@{0}
git stash clear
```

`drop` removes one stash entry. `clear` removes all stashes.

## Merging and Rebasing

### Merge

```bash
git switch main
git merge feature/version-control-notes
```

Merge preserves the complete branching history.

### Rebase

```bash
git switch feature/version-control-notes
git rebase main
```

Rebase rewrites your branch history so it appears on top of the latest `main`. It creates a cleaner linear history, but it should be used carefully on shared branches.

### Merge Conflicts

A merge conflict happens when Git cannot automatically combine changes.

General resolution process:

1. Open the conflicted file.
2. Review conflict markers.
3. Choose or combine the correct content.
4. Stage the resolved file.
5. Complete the merge or rebase.

Conflict markers look like this:

```text
<<<<<<< HEAD
your changes
=======
incoming changes
>>>>>>> feature-branch
```

## Branching Strategies

### Feature Branch Workflow

The most common approach for small and medium teams.

- Keep `main` stable.
- Create short-lived branches for work.
- Merge through pull requests.

### Git Flow

A more structured model with branches such as `main`, `develop`, `feature/*`, `release/*`, and `hotfix/*`.

Useful for teams that want clear release separation, but it can be heavier than necessary for fast-moving products.

### Trunk-Based Development

Developers integrate frequently into a shared trunk branch.

- Branches are short-lived.
- Work is merged often.
- CI must be strong.

This approach reduces long-lived merge pain and supports continuous delivery well.

## Good Commit Practices

- Keep commits small and focused.
- Use imperative commit messages, such as `Fix login redirect`.
- Do not mix unrelated changes in one commit.
- Commit often enough to preserve progress.
- Avoid committing secrets, build artifacts, or temporary files.

### Commit Message Examples

Good:

```bash
git commit -m "Add branch workflow documentation"
git commit -m "Fix merge conflict handling in README"
git commit -m "Refactor Git examples for clarity"
```

Less useful:

```bash
git commit -m "update"
git commit -m "changes"
git commit -m "fix stuff"
```

## Useful Git Configuration

### Set Your Identity

```bash
git config --global user.name "Your Name"
git config --global user.email "you@example.com"
```

### View Configuration

```bash
git config --list
```

### Useful Defaults

```bash
git config --global init.defaultBranch main
git config --global pull.rebase true
git config --global core.editor "code --wait"
```

These settings help standardize behavior across projects.

## Remote Repository Management

### Inspect Remotes

```bash
git remote -v
```

### Add a Remote

```bash
git remote add origin https://github.com/owner/repo.git
```

### Rename a Remote

```bash
git remote rename origin upstream
```

### Remove a Remote

```bash
git remote remove origin
```

### Fetch Remote Updates

```bash
git fetch
git fetch --all --prune
```

Fetching updates your remote-tracking branches without changing your working tree.

## Tags and Releases

Tags mark important points in history, usually releases.

### Create a Tag

```bash
git tag v1.0.0
git tag -a v1.0.0 -m "Release 1.0.0"
```

### List Tags

```bash
git tag
```

### Push Tags

```bash
git push origin v1.0.0
git push origin --tags
```

Annotated tags are preferred for releases because they store metadata and a message.

## Ignoring Files

Use `.gitignore` to exclude files that should not be tracked, such as:

- Secrets
- Logs
- Dependency folders
- Build output
- Editor-specific files

Example:

```gitignore
node_modules/
.env
dist/
.DS_Store
```

Do not rely on `.gitignore` to protect secrets that were already committed.

## GitHub and DevOps

Version control is tightly connected to DevOps because it enables automation and repeatability.

- CI/CD pipelines watch branches and run tests on every push.
- Infrastructure as code lives in Git and is reviewed like application code.
- Release tags create traceable deployment points.
- Branch protection enforces approvals and required checks.
- Git history helps with audits, compliance, and incident response.

## Practical Command Cheat Sheet

```bash
git init
git clone <repo-url>
git status
git add <file>
git add .
git commit -m "message"
git log --oneline --graph --decorate --all
git diff
git diff --staged
git branch
git switch <branch>
git switch -c <new-branch>
git merge <branch>
git rebase <branch>
git fetch
git pull
git push
git remote -v
git restore <file>
git restore --staged <file>
git revert <commit>
git tag -a v1.0.0 -m "Release 1.0.0"
git blame <file>
```

## Common Mistakes To Avoid

- Committing directly to protected branches.
- Making huge commits that are hard to review.
- Using `reset --hard` without understanding the impact.
- Rewriting shared history without coordination.
- Ignoring merge conflicts until the end of a long branch.
- Committing secrets or environment files.
- Treating GitHub as storage only and not using reviews, checks, and branch policies.

## Summary

Version control is the source of truth for modern software delivery. Git gives you the mechanics to track, branch, merge, and recover. GitHub gives you the collaboration layer that makes version control useful at team scale.

If you understand commits, branches, remotes, pull requests, merging, rebasing, and safe rollback patterns, you already have the core of professional DevOps version control practice.
