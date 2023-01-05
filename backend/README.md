# backend

## golang-migrate

### コンテナに入る
```
docker exec -it {コンテナ名} sh
```

### マイグレーションファイルの作成
```
migrate create --ext {拡張子} --dir {パス} --seq {マイグレーション名}
```

### マイグレーションファイルの実行
```
migrate --database "mysql://{ユーザー名}:{パスワード}@tcp({DBのコンテナ名}:{ポート番号})/{DB名}" --path {マイグレーションフィルのパス} up
```

### マイグレーションファイルを戻す
```
migrate --database "mysql://{ユーザー名}:{パスワード}@tcp({DBのコンテナ名}:{ポート番号})/{DB名}" --path {マイグレーションフィルのパス} down
```