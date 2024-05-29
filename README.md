# Go Authentication Demo

goでのユーザー認証とルーティング設定を把握しておくために作ったソースコード。
主にルーティングの設定の仕方が中心。標準パッケージだけでの試み。

![Gopher](https://www.clipartmax.com/png/middle/276-2767748_golang-gopher-jpg.png "Gopher")


# ディレクトリ構造

.
├── .env
├── .gitignore
├── cmd/
│ └── web/
│ └── main.go
├── constants/
│ └── constants.go
├── go.mod
├── go.sum
├── pkg/
│ ├── controllers/
│ │ ├── homeHandler.go
│ │ ├── mypageHandler.go
│ │ ├── router.go
│ │ ├── signinHandler.go
│ │ ├── signoutHandler.go
│ │ └── signupHandler.go
│ └── models/
│ ├── database.go
│ └── userRepository.go
├── public/
│ ├── css/
│ │ ├── form.css
│ │ ├── mypage.css
│ │ └── style.css
│ ├── img/
│ └── js/
├── README.md
└── views/
├── battle.gohtml
├── home.gohtml
├── layout.gohtml
├── mypage.gohtml
├── README.md
├── signin.gohtml
└── signup.gohtml

- `.env` : 環境変数を設定するファイル
- `cmd/web/main.go` : アプリケーションのエントリーポイント
- `constants/constants.go` : アプリケーション全体で使用する定数を定義するファイル
- `pkg/controllers/` : HTTPリクエストを処理するハンドラを定義するディレクトリ
- `pkg/models/` : データベースとのやり取りを行うモデルを定義するディレクトリ
- `public/` : 静的ファイル（CSS、画像、JavaScript）を格納するディレクトリ
- `views/` : HTMLテンプレートを格納するディレクトリ