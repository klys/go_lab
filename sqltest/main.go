package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

type User struct {
    ID   int64
    Name string
}

type Test struct {
	id int
	test string
}

func main() {
    db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Create a new user
    stmt, err := db.Prepare("INSERT INTO test (test) VALUES (?)")
    if err != nil {
        fmt.Println(err)
        return
    }
    _, err = stmt.Exec("Bowser")
    if err != nil {
        fmt.Println(err)
        return
    }

    // Read a user by ID
    row := db.QueryRow("SELECT id, test FROM test WHERE id = ?", 4)
    var u Test
    err = row.Scan(&u.id, &u.test)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("%+v\n", u)

    // Update a user by ID
    stmt1, err := db.Prepare("UPDATE test SET test = ? WHERE id = ?")
    if err != nil {
        fmt.Println(err)
        return
    }
    _, err = stmt1.Exec("Juliet", u.id)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Delete a user by ID
    stmt2, err := db.Prepare("DELETE FROM test WHERE id = ?")
    if err != nil {
        fmt.Println(err)
        return
    }
    _, err = stmt2.Exec(3)
    if err != nil {
        fmt.Println(err)
        return
    }
}