#!/bin/bash
CGO_ENABLED=O GOOS=linux go build -a -installsuffix cgo -o app app.go

