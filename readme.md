# Melchior

A CLI tool that translates natural language into shell commands using AI.

> The project is currently in early stages of development

## Planned features

- [x] Ollama support
- [ ] OpenAI API support
- [ ] Gemeni API suppprt
- [ ] Optinal verbose mode, detailed explanations for complex queries
- [ ] Package

## Description

Melchior is a command-line tool that converts natural language instructions into shell commands. No more searching through man pages or looking up commands online - just describe what you need, and Melchior will generate the correct command.

## Installation

WIP

## Configuration

Create a `config.toml` file:

```toml
mode = "ollama"
model = "mistral"
ollama_url = "http://localhost:11434"
```

## Prerequisites

- Ollama running locally or accessible via network
- Selected model pulled

## Usage

> For convenience, it's recommended to create a short shell alias for Melchior.


```bash
melchior <prompt>
```

### Examples

```bash
melchior initiate a docker swarm manager
docker swarm init

melchior generatre ed25565 ssh key pair
ssh-keygen -t ed25519
```