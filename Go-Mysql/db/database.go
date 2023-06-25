package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//username:password@tcp(localhost:3306)/database
const url = "root:valen1402@tcp(localhost:3306)/goweb_db"

// guarda la conexion a la base de datos
var db *sql.DB

// funcion que se conecta a la base de datos
func Connect() {
	connection, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Succes Conection")
	db = connection

}

// funcion que cierra la conexion a la base de datos
func Close() {
	db.Close()
}

// verificar la conexion
func Ping() {
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Succes Ping")
}

//verifica si existe la tabla

func ExistTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := Query(sql)
	if err != nil {
		panic(err)
	} else {
		return rows.Next()
	}
}

// Crea una tabla
func CreateTable(schema string, name string) {
	if !ExistTable(name) {
		_, err := db.Exec(schema)
		if err != nil {
			fmt.Println(err)
		}
	}
}

//Reiniciar registro de una tabla
func TruncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE TABLE %s", tableName)
	_, err := Exec(sql)
	if err != nil {
		fmt.Println(err)
	}
}



//Poliformismo de Exec
func Exec(query string, args ...any) (sql.Result, error){
	
	result, err:= db.Exec(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	return result, nil
}

//Poliformismo de Query
func Query(query string, args ...any) (*sql.Rows, error){
	rows, err:= db.Query(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	return rows, nil
}


