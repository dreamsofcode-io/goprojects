#!/usr/bin/env bash

air -c ./.air.toml & \
tailwindcss \
  -i 'static/css/main.css' \
  -o 'static/css/style.css' \
  --watch #& \
#templ generate -watch
