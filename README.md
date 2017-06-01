```
docker run -d --restart always -p 127.0.0.1:3306:3306 -e MYSQL_DATABASE=lightningtalk -e MYSQL_ROOT_PASSWORD=123456 --name local-mysql mysql:latest
```
