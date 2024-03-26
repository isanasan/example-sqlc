# example-sqlc

## setUp

ホストマシン上にて

```bash
docker run --name mysql -e MYSQL_ALLOW_EMPTY_PASSWORD=true -d -p 3306:3306 mysql:8.0
docker exec -it mysql bash
```

mysqlコンテナ内にて

```bash
mysql
CREATE DATABASE test CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
use test;
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    age INT NOT NULL
);
```

## execute

```bash
go run tutorial.go
```

## result

```bash
2024/03/26 14:19:06 1
2024/03/26 14:19:06 {1 test test@test.com 20}
2024/03/26 14:19:06 {1 test test@test.com 40}
```
