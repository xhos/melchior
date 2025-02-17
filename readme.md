# melchior  

a cli tool that translates natural language into shell commands using ai.  

> the project is currently in early stages of development  

## planned features  

- [x] ollama support  
- [ ] openai api support  
- [ ] gemini api support  
- [ ] optional verbose mode, detailed explanations for complex queries  
- [ ] package  

## description  

melchior is a command-line tool that converts natural language instructions into shell commands. no more searching through man pages or looking up commands online—just describe what you need, and melchior will generate the correct command.  

## installation  

wip  

## configuration  

create a `config.toml` file:  

```toml  
mode = "ollama"  
model = "mistral"  
ollama_url = "http://localhost:11434"
```

## prerequisites

- ollama running locally or accessible via network
- selected model pulled

## usage

> for convenience, it's recommended to create a short shell alias for melchior.

```bash
melchior <prompt>  
```

```bash
melchior initiate a docker swarm manager  
docker swarm init  

melchior generate ed25519 ssh key pair  
ssh-keygen -t ed25519  
```
