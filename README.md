# 概要
- Go言語を使ってCA・DDD・オブジェクト指向を試すディレクトリ。
- 目的は進めることではなく学ぶこと。理解や思考を優先して取り組む。
- 自分の中でGoに関するデファクトスタンダードを確立し、自分の指針となるコードにする。（これを元にGoの新規コードを爆速で開発できるようにする）

## 実行
```bash
go run cmd/main.go
localhost:8081/tasks
```

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
- 共通処理に関してApplicationディレクトリにServiceを用意する。
- domain層のディレクトリ構成に関して考える。
- tasksの命名に関して修正をする。
- 同パッケージ内のset可能な問題に関してIFを使って対応できないかを検討する。
- プレゼンターを用意する。
- レスポンスの処理がもっと簡潔になるように変更をする。
- DIが美しくないので修正をする。
- Factoryに関して確認をする。
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

## アーキテクチャ
Clean Architectureを意識したディレクトリ構成になります。

ルール
- 図の上から上位レイヤーと呼ぶ。
- 下位レイヤーから上位レイヤーへ依存を行う。同レイヤーへの依存はOK。上位レイヤーから下位レイヤーへ依存する場合はinterfaceを用いて依存性逆転の法則（DIP）を行う。


### Clean Architecture
Clean Architecture観点で考えた時の本ディレクトリ構成


|            レイヤー名            |    ディレクトリ名     |
|:---------------------------:|:--------------:|
| Enterprise Business Rules	  |     domain     |
| Application Business Rules	 |  application   |
|     Interface Adapters	     |    adaptor     |
|    Frameworks & Drivers	    | infrastructure |

### レイヤードアーキテクチャ
レイヤードアーキテクチャ観点で考えた時の本ディレクトリ構成

- 依存関係は上の層から下の層のみ許可する。
- 本プロジェクトはドメイン層からインフラ層への依存はDIPを使って逆転させているのでオニオンアーキテクチャになっている。

|    レイヤー名    |    ディレクトリ名     |
|:-----------:|:--------------:|
| プレゼンテーション層	 |    adaptor     |
|  アプリケーション層  |  application   |
|    ドメイン層    |     domain     |
|    インフラ層    | infrastructure |

## ディレクトリ構成 ルート配下
|    ディレクトリ名    |              詳細              |
|:-------------:|:----------------------------:|
| /environments |           環境変数を定義            |
|     /cmd      | アプリケーションのエントリポイント。ルーティングを行う。 |
|   /internal   |         アプリケーションを配置          |
|               |                              |

## ディレクトリ構成 /internal配下
|       ディレクトリ名       |                     詳細                     |
|:-------------------:|:------------------------------------------:|
|      /adapter       |              外部との接続を行う。APIなど。              |
| /adapter/controller | エンドポイント定義。HTTP Requestとのマッピング。入力値のバリデーション。 |

```shell
.
└── app
    ├── adaptor
    │   ├── controller
    │   │   └── task
    │   │       ├── tasks_controller.go
    │   └── presenter
    │       └── task
    │           └── tasks_presenter.go
    ├── application
    │   ├── service
    │   └── usecase
    │       └── task
    │           ├── get_tasks_input_port.go
    │           ├── get_tasks_output_port.go
    │           └── get_tasks_usecase.go
    ├── domain
    │   └── task
    │       ├── reward.go（Value Object）
    │       ├── task.go（Entity）
    │       ├── task_service.go
    │       └── tasks_repository.go
    └── infrastructure
        ├── database
        │   └── database.go
        ├── queryservice
        ├── register（DI）
        │   └── task.go
        └── repository
            └── task
                └── tasks_repository.go

```

### adapter/controller

### domain/entity.go
- ファイル名は `ディレクトリ名.go` とする。
- 単数系で命名する。
- `New構造体名` でコンストラクタを用意する。
- コンストラクタはただ初期化するだけではなく、不正値をバリデーションするようにする。
- 基本的にはポインタ型で返却をする。 ← 検討中
- Getterの命名はフィールド名の先頭大文字と一致させる。
- 構造体のフィールドはprivateで定義をする。 ← 同階層からの破壊的変更に関しての対策は検討中。
- 構造体のフィールドに対して基本的にタグを使わない。（ドメインは特定の技術に依存しないように作成する）

### domain/value_object.go
- ファイル名は `フィールド名.go` とする。
- 単数系で命名する。
- `Newフィールド名` でコンストラクタを用意する。
- コンストラクタはただ初期化するだけではなく、不正値をバリデーションするようにする。
- 不変である。そのため自身の状態を変更する関数は定義しない。値を変更する場合は新しくコンストラクタにて交換をする。

### domain/repository
- ファイル名は `ディレクトリ名(s)_repository.go` とする。
- 構造体名は `ディレクトリ名(s)Repository` とする。
- データ層とのやり取りを行うためのインターフェースを定義する。

### domain/service

## 用語解説
- DTO
- DataModel
- QueryService

### DTO
- Data Transfer Objectの略。
- CQRSを実現するために用い、複数の集約をまたがる参照系の処理を保存する構造体として使う。
- Application層のusecaseでIFを定義し、実装はinfrastructure層で行う。
- ユースケースごとに発生する集約に対応するため、複数箇所で使い回しは行わない。最大公約数的な使い方をしない。

### DataModel
- 本プロジェクト固有の命名かもしれない。（適切な命名があればそちらに移管予定）
- DBとのアクセスに用いる構造体。
- DataModelにデータを詰めた後、domain層のオブジェクトに詰め替えて上位レイヤーへ返却する。

### QueryService
- CQRSを実現するために用いる概念であり、usecaseにてIFを定義し、infrastructure層で実装する。
- 参照系専用のオブジェクト。

## コーディングルール
- パッケージ名はシンプルな名詞を使う。
  - ×：valueObject
  - △：value_object
  - ○：valueobject
- レイヤー間のパッケージ参照時にレイヤー名を付与する。命名は `ディレクトリ名+レイヤー名`
  - ○：TaskDomain

## メモ
- 1アプリケーションで複数の境界づけられたコンテキストを扱う場合は、`internal` 配下にディレクトリを作成し、その中にコンテキストごとにディレクトリを作成する。
  - internal/app
  - internal/someapp