package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
type User struct{
	 Name string `json:"name"`
	 Age uint16 `json:"age"` }
func main() {


	dsn := "root:@tcp(127.0.0.1:3306)/golang?parseTime=true&charset=utf8mb4&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Baglanyşygy barla
	err = db.Ping()
	if err != nil {
		panic(fmt.Sprintf("Cannot connect to DB: %s", err))
	}

	fmt.Println("Connected to MySQL!")

	// // Test INSERT
	// _, err = db.Exec("INSERT INTO users(name, age) VALUES(?, ?)", "Es", 23)
	// if err != nil {
	// 	panic(fmt.Sprintf("Insert failed: %s", err))
	// }

	res, err:=db.Query("SELECT `name`,	`age` from `users`")
	if err != nil {
		panic(fmt.Sprintf("Cannect : %s", err))
	}
	for res.Next(){
		var user User
		err= res.Scan(&user.Name,&user.Age)
			if err != nil {
		panic(fmt.Sprintf("Cannect : %s", err))
	}
	fmt.Println(fmt.Sprintf("User: %s with age %d",user.Name, user.Age))
	}

}
