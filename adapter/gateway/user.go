package gateway

/*
gateway パッケージは，DB操作に対するアダプター（永続化に関するアダプター）
port.UserRepository interface の実装を行う
*/

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"example.com/practice-clean-architecture/entity"
	"example.com/practice-clean-architecture/usecase/port"
)

type UserRepository struct {
	conn *sql.DB
}

// NewUserRepository はUserRepositoryを返します．
func NewUserRepository(conn *sql.DB) port.UserRepository {
	return &UserRepository{
		conn: conn,
	}
}

// GetUserByID は、userIDをもとにDBからデータを取得します．
func (u *UserRepository) GetUserByID(ctx context.Context, userID string) (*entity.User, error) {
	conn := u.GetDBConn()
	row := conn.QueryRowContext(ctx, "SELECT * FROM `user` WHERE id=?", userID)
	user := entity.User{}
	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found. UserID = %s", userID)
		}
		log.Println(err)
		return nil, errors.New("internal server error. adapter/gateway/GetUserByID")
	}
	return &user, nil
}

// GetDBConn はconnectionを取得します．
func (u *UserRepository) GetDBConn() *sql.DB {
	return u.conn
}
