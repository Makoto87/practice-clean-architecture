package driver

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"example.com/practice-clean-architecture/adapter/controller"
	"example.com/practice-clean-architecture/adapter/gateway"
	"example.com/practice-clean-architecture/adapter/presenter"
	"example.com/practice-clean-architecture/usecase/interactor"
	_ "github.com/go-sql-driver/mysql"
)

// Serve はserverを起動
func Serve(addr string) {
	// DB接続
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DATABASE"))

	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
		return
	}

	// controller 使用
	user := controller.User{
		OutputFactory: presenter.NewUserOutputPort,
		InputFactory:  interactor.NewUserInputPort,
		RepoFactory:   gateway.NewUserRepository,
		Conn:          conn,
	}
	http.HandleFunc("/user/", user.GetUserByID)
	// port 指定してサーバーが起動
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}
