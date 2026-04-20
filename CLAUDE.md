# CLAUDE.md

This file provides guidance to Claude Code when working with code in this repository.

## Project

Generated Go SDK for the Paraph API. Source is regenerated from `openapi.yaml` in the `paraph` repo via `make sdk-go`. Do not hand-edit generated files — edit `openapi.yaml` and regenerate instead.

## Common Commands

```bash
# Regenerate from the paraph repo:
# make sdk-go  (run from ../paraph)
```

## Git Workflow

- Always create a feature branch for changes — never commit directly to main
- Use `git worktree add ../paraph-go-<branch> -b <branch>` to work in an isolated directory
- Push the feature branch and create a PR for review before merging
- Never use `git push` to main directly
