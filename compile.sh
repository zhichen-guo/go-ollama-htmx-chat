#!/bin/bash

go install
env GOOS=darwin GOARCH=arm64 go install