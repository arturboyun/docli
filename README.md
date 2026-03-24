# Docli

Docli is a command-line interface (CLI) tool for managing docker containers.

## Installation

To install Docli, you can use pip:

```bash
go build -o bin/docli main.go
```

## Usage

To use Docli, simply run the following command:

```bash
./bin/docli [command] [options]
```

## Commands

- `ps`: List all running containers
- `logs <container_id>`: View logs of a container
- `cp <container_id>:<source_path> <destination_path>`: Copy files from a container to the host machine
