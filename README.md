# gh permalink

[![CI](https://github.com/spenserblack/gh-permalink/actions/workflows/ci.yml/badge.svg)](https://github.com/spenserblack/gh-permalink/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/spenserblack/gh-permalink/branch/main/graph/badge.svg?token=HOyIAY5JyM)](https://codecov.io/gh/spenserblack/gh-permalink)

Create a permalink from the CLI

## Installation

```shell
gh extension install spenserblack/gh-permalink
```

## Usage

```shell
# Permalink a file
gh permalink my-file

# Permalink a line
gh permalink my-file 1

# Permalink a range of lines
gh permalink my-file 1-5
```

### Advanced

#### Permalinking Commits

This extension creates a permalink from the current HEAD commit. If you need to create a permalink
to a different commit, checkout that commit first.

```shell
git checkout v0.1.0
gh permalink my-file
# now you have a permalink to my-file at the commit tagged by v0.1.0
git checkout -  # go back to where you were previously checked out
```
