package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)



func main() {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/bodega")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM persona WHERE id_persona=?")

	res, err := stmt.Exec(1)

	affect, err := res.RowsAffected()

	fmt.Println(affect)



	
	stmt, err := db.Prepare("UPDATE persona SET nombre=? WHERE id_persona=?")

	res, err := stmt.Exec("mario", 1)

	affect, err := res.RowsAffected()

	fmt.Println(affect)










	var id_persona int
	var nombre, apellido, rut string

	err = db.QueryRow("SELECT * FROM persona").Scan(&id_persona, &nombre, &apellido, &rut)

	if err != nil{
		panic(err.Error())
	}

	fmt.Println(id_persona, nombre, apellido, rut)







/*
	insert, err := db.Query("INSERT INTO persona (nombre, apellido, rut) VALUES ('','qwer','aasdf3432')")

	if err != nil {
		panic(err.Error())
	}

	insert.Close()*/

	/*result, err := 	db.Query("SELECT * FROM persona")

	if err != nil {
		panic(err.Error())
	}

	for result.Next(){
		var id_persona int
		var nombre, apellido, rut string

		err = result.Scan(&id_persona, &nombre, &apellido, &rut)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(id_persona)
		fmt.Println(nombre)
		fmt.Println(apellido)
		fmt.Println(rut)

	}*/



}