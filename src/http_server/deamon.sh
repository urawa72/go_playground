#!/bin/sh

# apt-get install -y supervisor

go build -o /

echo "
[program:http_server]
command = /http_server
autostart = true
startsecs = 5
user = root
redirect_stderr = true
" >> /etc/supervisor/supervisord.conf
