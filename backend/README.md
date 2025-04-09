# ファイル構成
## internal
- domain
- usecase
- controller
- infrastructure
### domain
アプリケーションの中心となるエンティティを定義する。どこにも依存しないコードを格納する。
### usecase
アプリケーションの操作の流れを定義する。domainで定義されたデータ構造を使用して実装する。
#### repository
usecaseのデータ操作を行う部分。infrastructureのデータ制御部分に依存するが、usecaseのdomain以外に依存する部分はrepositoryに隔離する。
#### service
usecaseの外部に向けて操作を行う部分。データ操作はrepositoryを使って行い、controllerではservice内の関数を呼び出して使う。
### controller
外部からのリクエストを受け取り、適切なusecaseを呼び出すフォルダ。主にAPIのエンドポイントを定義し、リクエストのバリデーションやレスポンスの整形を行う。
#### dto
リクエスト・レスポンスの際に用いるデータ構造を格納する。
#### converter
domainで定義されたデータ構造と、dtoのデータ構造とを変換するコードを格納する。
#### generated
OpenAPI関連のツールで生成されたコードを格納する。
### infrastructure
データベース操作や外部サービスとの連携など、アプリケーションの実装依存部分をまとめたフォルダ。usecaseから利用される。
#### db
データベースへの接続や設定を行うコードを格納するフォルダ。
#### M304
M304の制御を行うコードを格納するフォルダ。
#### orm
ORM等のデータ操作ツールに依存するコードを格納するフォルダ。
### utils
