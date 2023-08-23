package services

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

    "fmt"
    "time"
)

var DB *sqlx.DB;

func Init_DB () {
    var err error;
    // if database already exists (not nil), return the instance
    // kinda singleton but imporved
    if DB != nil {
        fmt.Println("Database pool already exists");
        return
    }
    // try 5 times before returning an error
    for i := 0; i < 5; i++ {
        DB, err = open_database();
        if err == nil {
            break
        } else {
            fmt.Println(err)
            time.Sleep(time.Second * 5)
        }
    }
}


func open_database() (*sqlx.DB, error) {

    db, err := sqlx.Connect("sqlite3", "database.db")
    if err != nil {
        fmt.Println(err)
        fmt.Println("Coudln't connect to database")
        return nil, err
    }


    // setting up the DB as a connection pool
    db.DB.SetMaxOpenConns(1000) // The default is 0 (unlimited)
	db.DB.SetMaxIdleConns(10) // defaultMaxIdleConns = 2
	db.DB.SetConnMaxLifetime(0) // 0, connections are reused forever.
// for now the DB will be a global variable so every package can use it

    return db, nil
}
