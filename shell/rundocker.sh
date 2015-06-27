#!/bin/bash

docker run --name blog-redis -v /opt/db/redis:/data -p 6379:6379 -d ckeyer/redis:3.0

docker run --name blog-mysql -v /opt/db/mysql:/var/lib/mysql  -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -d ckeyer/mysql:5.7

docker run --name blog-nginx -p 80:80 -p 443:443 -d ckeyer/nginx:localhost
