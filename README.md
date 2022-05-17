# gohub

Go で Web API を書くための下地プロジェクトです。
センスないなって思ったら適当にプルリク投げてください。

フレームワークは`net/http`を利用するつもりです。
Go のバージョンは 1.18

## Docker

### 「gohub api」のコンテナにアクセスする方法

```bash
# 矢印キーが効かないので「bin/bash」で入った方がいい
docker exec -it gohub_api /bin/sh
docker exec -it gohub_api /bin/bash
```

## 「gohub_db」 のコンテナにアクセスする方法

```bash
 # 矢印キーが効かないので「bin/bash」で入った方がいい
docker exec -it gohub_db /bin/sh
docker exec -it gohub_db /bin/bash

# psqlに入る
psql -h localhost -U postgres
```
