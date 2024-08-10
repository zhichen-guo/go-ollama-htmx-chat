# Golang + HTMX + Ollama Chat

This is the first app I've ever built using Golang. It is a simple interface to chat with an Ollama API. I tried to do as much from scratch as I could, so there's no special library for templates, just an html file with format placeholders that's passed to fmt.Fprintf. Eventually I also want to write my own basic implementation of HTTP to replace the net/http library.

## Stack:
* Golang and HTMX server
* Ollama for LLMs
* Containerized using Docker

## Instructions:
1. Clone repo
2. Spin up containers: ```docker compose up```
3. In a new terminal window, install the necessary model in the ollama container: ```docker compose exec ollama ollama pull qwen2```
4. Server will be accessible in ```localhost:8080```
