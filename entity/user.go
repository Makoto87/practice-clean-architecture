package entity

/*
entity パッケージは，ドメインモデルを実装
他のレイヤに依存することはない
DB操作のような技術的な実装もない
*/

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
