package presenter

/*
presenter パッケージは，出力に対するアダプター

ここでは，アウトプットポートを実装(interfaceを満たすようにmethodを追加する)
port.UserOutputPort を実装
*/

import (
	"fmt"
	"net/http"

	"example.com/practice-clean-architecture/entity"
	"example.com/practice-clean-architecture/usecase/port"
)

type User struct {
	w http.ResponseWriter
}

// NewUserOutputPort はUserOutputPortを取得
func NewUserOutputPort(w http.ResponseWriter) port.UserOutputPort {
	return &User{
		w: w,
	}
}

// usecase.UserOutputPortを実装している
// Render はNameを出力
func (u *User) Render(user *entity.User) {
	u.w.WriteHeader(http.StatusOK)
	// httpでentity.User.Nameを出力
	fmt.Fprint(u.w, user.Name)
}

// RenderError はErrorを出力
func (u *User) RenderError(err error) {
	u.w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(u.w, err)
}
