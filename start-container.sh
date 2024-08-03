#!/bin/bash

docker run -it --rm -p 8080:8080 --name golang -v .:/go/go-http golang
