package port

/*
port パッケージは，出力や入力などのポート(interface)を提供
*/

import (
	"context"

	"example.com/practice-clean-architecture/entity"
)

// interactorに実装される
// adapter/contorller に依存される
type UserInputPort interface {
	GetUserByID(ctx context.Context, userID string)
}

// interactorに依存される
// adapter/presenter に実装される
type UserOutputPort interface {
	Render(*entity.User)
	RenderError(error)
}

// userのCRUDに対するDB用のポート
// interactorに依存される
// adapter/gateway に実装される
type UserRepository interface {
	GetUserByID(ctx context.Context, userID string) (*entity.User, error)
}
