# Golang + HTMX + Ollama Chat

This is the first app I've ever built using Golang. It is a simple interface to chat with the Ollama API hosted in a separate container. 

The server has two components: the ```server``` file handles all the requests from the browser and the ```ollama_client``` file handles all the communications with Ollama.

I tried to do as much from scratch as I could, so there's no library for HTML templates, just an html file with string format placeholders that's passed to fmt.Fprintf. Eventually I also want to write my own basic implementation of HTTP to replace the net/http library.

## Stack:
* Golang and HTMX server
* Ollama for LLMs
* Containerized using Docker

## Instructions:
1. Clone repo
2. Spin up containers: ```docker compose up```
3. In a new terminal window, install the necessary model in the ollama container: ```docker compose exec ollama ollama pull qwen2```
4. Server will be accessible in ```localhost:8080```

## Other commands
* To access the golang container, use ```docker compose exec server bash```
