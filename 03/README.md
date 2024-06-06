# 03 自作のスキーマを使ってGraphQLサーバーを作ろう
GitHub API v4 を真似してサーバーサイドのコードを作る

## gqlgen で生成されるディレクトリ構造をカスタマイズする

今回の最終的なディレクトリ構成は以下のようになる．
```
 .
+├─ internal
+│   └─ generated.go # このファイルの中身は編集しない
 ├─ graph
-│   ├─ generated.go # このファイルの中身は編集しない
 │   ├─ model
 │   │   └─ models_gen.go # 定義した型が構造体として定義される
 │   ├─ resolver.go
-│   ├─ schema.graphqls # スキーマ定義
 │   └─ schema.resolvers.go # この中に、各queryやmutationのビジネスロジックを書く
+├─ schema.graphqls # スキーマ定義
 ├─ gqlgen.yml # gqlgenの設定ファイル
 ├─ server.go # エントリポイント
 ├─ go.mod
 └─ go.sum
```

このような構成で自動生成させるために，`gqlgen.yml`を書き換える．
```
# 自動生成コードの元となるGraphQLスキーマがどこに配置してあるか
schema:
-  - graph/*.graphqls
+  - ./*.graphqls

# 自動生成されるgeneated.goの置き場所
exec:
-  filename: graph/generated.go
-  package: graph
+  filename: internal/generated.go
+  package: internal

# スキーマオブジェクトに対応するGo構造体の置き場所
model:
  filename: graph/model/models_gen.go
  package: model

# リゾルバコードの置き場所
resolver:
  layout: follow-schema
  dir: graph
  package: graph
```

`schema.graphqls`ファイルを作成する．GraphQLスキーマを定義


ここまできたら，`go mod init`でプロジェクトを準備し，`go run github.com/99designs/gqlgen generate`でコードを自動生成する．

