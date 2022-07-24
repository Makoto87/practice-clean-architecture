package interactor

/*
interactor パッケージは，インプットポートとアウトプットポートを繋げる責務を持つ

interactorはアウトプットポートに依存(importする)
依存先
- usecase/port/userRepository
- usecase/port/outputPort

インプットポートを実装(interfaceを満たすようにmethodを追加する)
usecase/port の UserInputPort を実装
*/

import (
	"context"

	"github.com/arkuchy/clean-architecture-sample-sample/usecase/port"
)

type User struct {
	OutputPort port.UserOutputPort
	UserRepo   port.UserRepository
}

// NewUserInputPort は User (UserInputPort を実装したもの) を取得します．
func NewUserInputPort(outputPort port.UserOutputPort, userRepository port.UserRepository) port.UserInputPort {
	return &User{
		OutputPort: outputPort,
		UserRepo:   userRepository,
	}
}

// usecase.UserInputPortを実装
// GetUserByID は，UserRepo.GetUserByIDを呼び出し，その結果をOutputPort.Render or OutputPort.RenderErrorに渡す
func (u *User) GetUserByID(ctx context.Context, userID string) {
	user, err := u.UserRepo.GetUserByID(ctx, userID)
	if err != nil {
		u.OutputPort.RenderError(err)
		return
	}
	u.OutputPort.Render(user)
}