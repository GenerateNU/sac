# SAC (Student Activity Calendar) CLI

[![CLI](https://github.com/GenerateNU/sac/actions/workflows/cli.yml/badge.svg)](https://github.com/GenerateNU/sac/actions/workflows/cli.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/GenerateNU/sac/cli)](https://goreportcard.com/report/github.com/GenerateNU/sac/cli)
[![License](https://img.shields.io/github/license/GenerateNU/sac)](https://github.com/GenerateNU/sac/blob/main/LICENSE)



## Introduction
The SAC CLI is a tool designed to help manage and automate tasks for the GenerateNU Student Activity Calendar platform. It provides a comprehensive suite of commands for formatting, testing, linting, running, and managing databases and other configurations essential for the SAC platform.

## Table of Contents

1. [Installation](#installation)
2. [Usage](#usage)
3. [Available Commands](#available-commands)
4. [Command Aliases](#command-aliases)
5. [Dependencies](#dependencies)
6. [License](#license)

## Installation

To install the SAC CLI, clone the repository and run the following script from the root of the application:

```bash
./install_cli.sh
```

## Usage

To use the SAC CLI, you can run it with various flags or commands:

```
sac [flags]
sac [command]
```

## Available Commands

- **completion**: Generate the autocompletion script for the specified shell.
- **database**: Database management commands.
- **format**: Formatting commands.
- **help**: Help about any command.
- **lint**: Linting commands.
- **run**: Run commands for backend and frontend.
- **setup**: Installs and sets up the project. (WIP)
- **swagger**: Run swagger initialization for backend openapi spec.
- **test**: Testing commands.

For more detailed information about a command, use:

```bash
sac [command] --help
```

## Command Aliases

Several commands in the SAC CLI are equipped with short aliases to make them quicker and easier to use. For example, a command like `sac format frontend dashboard` which is pretty lengthy can be shortened to `sac f fe d`.

## Dependencies

- `golangci-lint`: A linter aggregator for Go, ensuring code quality.
- `gofumpt`: A stricter formatter for Go code, ensuring consistency.
