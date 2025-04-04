# バックエンド 開発仕様書
## 技術スタック
### 言語
Go
### フレームワーク
Gin
### データベース関係
- MySQL
- sqlc
- golang-migrate
### API関係
- OpenAPI
- oapi-codegen
## 主要フォルダ
```
backend
├── .devcontainer/              # devcontainer設定
├── cmd/                        # エントリーポイント
├── db/                         # DB操作用SQLファイル
├── docs/                       # ドキュメント
├── internal/                   # 内部ロジック
│   ├── controller/             # リクエストを受け取り、ユースケースを呼び出す
│   ├── domain/                 # エンティティ
│   ├── infrastructure/         # 外部ツールを利用した実装
│   ├── testdata/               # テスト用のダミーデータ
│   ├── usecase/                # アプリケーションの操作
│   │   ├── interfases/         # インタフェース
│   │   ├── repository/         # データベース操作
│   │   └── service/            # ビジネスロジックの実行
│   └── utils/                  # 汎用的なコード
├── openapi/                    # API仕様書（OpenAPI）
├── test/                       # テスト用コンテナ設定
└── others                      # その他ファイル
```
## 開発環境のセットアップ
golang-migrateでのマイグレーション実行コマンド
```
migrate -source file://db/migration -database "mysql://user名:パスワード@tcp(コンテナ名:3306)/DB名" up
```
