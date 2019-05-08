package main

import (
    "log"
    "net/http"
    "html/template"
    "fmt"  
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
      "time"
)


/////////////////////////////
type Pet struct { 
    Name  string
    Color string
    Date string
    Time string
}
///////////////////////// 

func HomePage(w http.ResponseWriter, r *http.Request){
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
    
    now := time.Now()
    pet:= Pet{}
    pet.Name = name
    pet.Color = color
    pet.Date = now.Format("02-01-2006")
    pet.Time = now.Format("15:04:05")
    
    t, err := template.ParseFiles("homepage.html") //parse the html file homepage.html
    if err != nil { // if there is an error
  	  log.Print("template parsing error: ", err) // log it
  	}
    err = t.Execute(w, pet) //execute the template and pass it the pet struct to fill in the gaps
    if err != nil { // if there is an error
  	  log.Print("template executing error: ", err) //log it
  	}
    
}
//////////////////
func main() {
    
   log.Println("Abrir home page en: http://localhost:8080")
   fmt.Printf("Iniciando server HTTP ...\n")   
   http.HandleFunc("/", HomePage)
   log.Fatal(http.ListenAndServe(":8080", nil))
    
//////////////////

    
  
} // End




