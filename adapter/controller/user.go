package controller

/*
controller パッケージは，入力に対するアダプター

ここではインプットポートとアウトプットポートを組み立てて，インプットポートを実行
ユースケースレイヤからの戻り値を受け取って出力する必要はなく，
純粋にhttpを受け取り，ユースケースを実行
*/

import (
	"database/sql"
	"net/http"
	"strings"

	"example.com/practice-clean-architecture/usecase/port"
)

type User struct {
	OutputFactory func(w http.ResponseWriter) port.UserOutputPort
	// -> presenter.NewUserOutputPort
	InputFactory func(o port.UserOutputPort, u port.UserRepository) port.UserInputPort
	// -> interactor.NewUserInputPort
	RepoFactory func(c *sql.DB) port.UserRepository
	// -> gateway.NewUserRepository
	Conn *sql.DB
}

// GetUserByID は，httpを受け取り，portを組み立てて，inputPort.GetUserByIDを呼び出します．
func (u *User) GetUserByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// /users/ より右のパスを取得 = user id 取得
	userID := strings.TrimPrefix(r.URL.Path, "/user/")
	// 抽象化されているが、後で実装されたものを利用する
	outputPort := u.OutputFactory(w)
	repository := u.RepoFactory(u.Conn)
	inputPort := u.InputFactory(outputPort, repository)
	inputPort.GetUserByID(ctx, userID)
}
