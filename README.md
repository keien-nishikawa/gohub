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
