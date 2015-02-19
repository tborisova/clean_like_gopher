package main

import (
    "database/sql"
    // "fmt"
    _ "github.com/go-sql-driver/mysql"
)


func main(){
  db, err := sql.Open("mysql", "root:@/golangtest")
  if err != nil {
      panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
  }
  defer db.Close()
  
  // var age int
  // name := "foo"
  

  // row := db.QueryRow("SELECT age FROM users WHERE name = ?", name)
  // err = row.Scan(&age)


/*
  check how many rows
  var count int
  row := db.QueryRow("SELECT count(id) as count FROM users")
  err = row.Scan(&count)
  fmt.Printf("The age of foo is: %d", count)
*/
  // insert
  // stmtIns, err := db.Prepare("INSERT INTO users VALUES( ?, ?, ? )") // ? = placeholder
  // if err != nil {
  //     panic(err.Error()) // proper error handling instead of panic in your app
  // }
  // defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

  // Insert square numbers for 0-24 in the database
  // for i := 0; i < 25; i++ {
      // _, err = stmtIns.Exec(2, 14, "bar") // Insert tuples (i, i^2)
      // if err != nil {
      //     panic(err.Error()) // proper error handling instead of panic in your app
      // }
  // }

  // delete
  // stmtIns, err := db.Prepare("DELETE FROM users") // ? = placeholder
  // if err != nil {
  //     panic(err.Error()) // proper error handling instead of panic in your app
  // }
  // defer stmtIns.Close() // Close the statement when we leave main() / the program terminates
  // _, err = stmtIns.Exec() 
  // if err != nil {
  //         panic(err.Error()) // proper error handling instead of panic in your app
  //     }

  // TRUNCATE
 // stmtIns, err := db.Prepare("TRUNCATE users") // ? = placeholder
  // if err != nil {
  //     panic(err.Error()) // proper error handling instead of panic in your app
  // }
  // defer stmtIns.Close() // Close the statement when we leave main() / the program terminates
  // _, err = stmtIns.Exec() 
  // if err != nil {
  //         panic(err.Error()) // proper error handling instead of panic in your app
  //     }
   
}