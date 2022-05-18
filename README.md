# gohub

Go で Web API を書くための下地プロジェクトです。
センスないなって思ったら適当にプルリク投げてください。

フレームワークは`net/http`を利用するつもりです。
Go のバージョンは 1.18

## Develop Flow

### Commands

```bash
docker compose build
docker volume create pgadmin4_volume
docker volume up
```

### URL

**SQL Client Tool**
http://localhost:8000

**API Application**
http://localhost:8080

### How to use PgAdmin4

1. access to url => http://localhost:8000
2. Login Info: gohub.api@sample.com/gohub_password
3. Click to [Add New Server]
4. Input info
   1. **General**
      - name: gohub_dev
   2. **Connection**
      - Host name/address: gohub_db
      - Port: 5432
      - Maintenance database: postgres
      - Username: postgres
      - Password: gohub_dev_password
      - Save Password: check <= as you like, but recommend to check

## Docker

### 「gohub api」のコンテナにアクセスする方法

```bash
# 矢印キーが効かないので「bin/bash」で入った方がいい
docker exec -it gohub_api /bin/sh
docker exec -it gohub_api /bin/bash
```

### 「gohub_db」 のコンテナにアクセスする方法

```bash
 # 矢印キーが効かないので「bin/bash」で入った方がいい
docker exec -it gohub_db /bin/sh
docker exec -it gohub_db /bin/bash

# psqlに入る
psql -h localhost -U postgres
```

### 環境変数を確認

```bash
docker exec gohub_api env
```
