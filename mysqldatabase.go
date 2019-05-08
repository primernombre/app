package main

import (
    "fmt"  
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)


func main() {

    fmt.Println("MySQL DataBase")
    
    db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/apptest")

    if err != nil {
    panic(err.Error()) 
    }
    defer db.Close()

    // Open doesn't open a connection. Validate DSN data:
    err = db.Ping()
    if err != nil {
        panic(err.Error()) 
    }

    // Execute the query
    var name string
    var color string
    err = db.QueryRow("select name, color from pet where species = ?", "Beagle").Scan(&name, &color)
    if err != nil {
         panic(err.Error())
    }
    
    fmt.Println("Mi perro beagle se llama: ")
    fmt.Println(name)
    fmt.Println(color)

  
} // End
