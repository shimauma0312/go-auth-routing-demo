package controllers

import (
	tmp "go-auth-demo/constants"
	"net/http"
	"text/template"
)

// サインインリクエストを処理
func (app *App) signinHandler(w http.ResponseWriter, r *http.Request) {
	// GETリクエストの場合、サインインページを表示
	if r.Method == "GET" {
		// サインインページのテンプレートを読み込み
		t, err := template.ParseFiles(tmp.Dir + tmp.Signin)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// テンプレートを実行し、結果をレスポンスに書き出し。
		err = t.ExecuteTemplate(w, tmp.Signin, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	// POSTリクエストの場合、サインイン処理
	if r.Method == "POST" {
		// リクエストからフォームデータを解析
		r.ParseForm()
		// フォームからユーザー名とパスワードを取得
		formUsername := r.FormValue("username")
		formPassword := r.FormValue("password")
		// データベースからユーザー情報を取得
		dbUser, err := app.DB.GetUser(formUsername)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// ユーザー一致をcheck
		if dbUser.UserID == formUsername && dbUser.Password == formPassword {
			session, _ := app.Get(r, "user")
			session.Values["id"] = dbUser.UserID
			session.Save(r, w)
			http.Redirect(w, r, "/mypage", http.StatusFound)
		}
	}
}
