# フロントエンド 開発仕様書
## 技術スタック
### 言語
TypeScript
### フレームワーク
Next.js
## 主要フォルダ
```plaintext
frontend
├── .devcontainer/              # devcontainer設定
├── docs/                       # ドキュメント
├── public/                     # 画像など静的ファイル
├── src/                        # アプリケーション本体
│   ├── app/                    # ページに関連するファイル（App Router）
│   ├── components/             # 汎用コンポーネント
│   ├── contexts/               # コンテキスト
│   ├── features/               # 機能別で扱うコンポーネント
│   ├── hooks/                  # カスタムフック
│   ├── lib/                    # 外部ライブラリとの連携
│   ├── styles/                 # テーマ設定などスタイル関係
│   └── types/                  # interfaceやtypeの定義
└── others                      # その他ファイル
```
## 開発環境のセットアップ
### DB設定
#### 環境変数の設定
.envファイルを作成し、環境変数を設定する。
