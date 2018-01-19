#!/usr/bin/bash
env GOOS=linux GOARCH=386 go build -o bin/gost2netpro_lin
env GOOS=windows GOARCH=386 go build -o bin/gost2netpro_win