package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called students
	db, err := sql.Open("mysql", "sun01822:@Sun2021Yes#@tcp(127.0.0.1:3306)/students")

	// if there is an error opening the connection, handle it
	if err != nil{
		panic(err.Error())
	}else{
		fmt.Println("Database connected successfully")
	}


	// defer the close till after the main function has finished
	defer db.Close()

}