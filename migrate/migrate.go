// NOTE: main関数コードを使用する
package main

import (
	"fmt"
	"go-todo-api/db"
	"go-todo-api/models"
)


func main(){
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated!")
	defer db.CloseDB(dbConn)
	//GORMに対してUserモデルとTodoモデルの実際のデータ構造を変更するよう指示。
	//これらのモデルがデータベース内のテーブルとしてどのカラムを持つか、どのデータ型を使用するかを更新
	//dbConn.AutoMigrate(models.User{}, models.Todo{})の場合、GORMに対して実際のデータベースのスキーマを変更する指示が通らない。
	//代わりに、UserとTodoモデルのコピーを作成しそれらのコピーに対してスキーマ変更を行っているが、これらのコピーは元のモデルとは独立していて実際のデータベーステーブルに変更を加えない。結果として、データベースに変更が反映されない
	dbConn.AutoMigrate(&models.User{}, &models.Todo{})

}