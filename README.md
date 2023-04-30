# 概要
- Go言語を使ってCA・DDD・オブジェクト指向を試すディレクトリ。
- 目的は進めることではなく学ぶこと。理解や思考を優先して取り組む。
- 自分の中でGoに関するデファクトスタンダードを確立し、自分の指針となるコードにする。（これを元にGoの新規コードを爆速で開発できるようにする）

# ルール
- ドキュメント系は忘れるので全てREADMEに記述する。
- 将来的にpublicにする前提でコードを書く。

# 機能
- 環境変数の管理
- GoでDB管理・マイグレーション
- インフラ
  - TerraformでAWSのリソースを管理
    - Parameter Store
- APIを提供する
- docker
  - ホットリロード
- 認証
  - Auth0 or Amazon Cognito
- ルーティング
- ミドルウェア
- DBアクセス
  - MySQL
- ロギング
- CI/CD
- ディレクトリ構成
  - MVCとCAを同居させる？CAで試してみる？
  - できるだけ標準のGoの書き方で行う（汎用性を高めるために）


## タスク管理ツール
- タスク
- ユーザー
  - Email

## ロギング
- ログを出力する条件
  - err!=nilの場合
- ログの出力先
  - local
  - dev
  - prod
- ログレベル
  - info
  - warn
  - error

## 疑問
- mysqlの接続設定の種類
-  db.Close()のハンドリング

# TODO
- テーブル定義の管理
- マイグレーションの管理
- システム全体のエラーハンドリングのルール策定
- ロギング
- テストの導入
  - できる限り自動生成でテストコードを記述する。
  - 並列処理で実行時間を短縮する。
- CIの導入
  - PR時に自動でテストを実行する。
- Terraformの導入
- ルーティングの記述
- 作成するプログラムの要件定義

## ディレクトリ構成
### ルート配下
|    ディレクトリ名    |              詳細              |
|:-------------:|:----------------------------:|
| /environments |           環境変数を定義            |
|     /cmd      | アプリケーションのエントリポイント。ルーティングを行う。 |
|   /internal   |         アプリケーションを配置          |
|               |                              |

### /internal配下
|    ディレクトリ名    |               詳細                |
|:-------------:|:-------------------------------:|
|    /adapter    |        外部との接続を行う。APIなど。         |
|    /domain     |      ドメイン層。ビジネスロジックを記述する。       |