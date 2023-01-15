curl -X POST -d 'username=test&password=123456' http://127.0.0.1:8888/v1/login
curl -X POST -d 'username=test' http://127.0.0.1:8888/v1/logout
curl http://127.0.0.1:8888/v1/users