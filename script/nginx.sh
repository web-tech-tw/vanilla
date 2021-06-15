#!/bin/bash

## Install
apk add --no-cache nginx
mkdir -p /run/nginx

## Finish
echo "NGINX -> [Build] Success"
