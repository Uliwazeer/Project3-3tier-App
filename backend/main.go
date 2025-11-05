package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
    "time"

    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
)

func connect() (*sql.DB, error) {
    password := os.Getenv("DB_PASSWORD")
    if password == "" {
        return nil, fmt.Errorf("DB_PASSWORD environment variable not set")
    }

    host := os.Getenv("DB_HOST")
    if host == "" {
        host = "mysql"
    }

    user := os.Getenv("DB_USER")
    if user == "" {
        user = "root"
    }

    dbName := os.Getenv("DB_NAME")
    if dbName == "" {
        dbName = "example"
    }

    dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, password, host, dbName)
    return sql.Open("mysql", dsn)
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
    db, err := connect()
    if err != nil {
        w.WriteHeader(500)
        log.Println("DB connection error:", err)
        return
    }
    defer db.Close()

    rows, err := db.Query("SELECT title FROM blog")
    if err != nil {
        w.WriteHeader(500)
        log.Println("Query error:", err)
        return
    }

    var titles []string
    for rows.Next() {
        var title string
        _ = rows.Scan(&title)
        titles = append(titles, title)
    }
    json.NewEncoder(w).Encode(titles)
}

func main() {
    log.Print("Prepare db...")
    if err := prepare(); err != nil {
        log.Fatal(err)
    }

    log.Print("Listening on port 8000")
    r := mux.NewRouter()
    r.HandleFunc("/", blogHandler)
    log.Fatal(http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, r)))
}

func prepare() error {
    db, err := connect()
    if err != nil {
        return err
    }
    defer db.Close()

    for i := 0; i < 60; i++ {
        if err := db.Ping(); err == nil {
            break
        }
        time.Sleep(time.Second)
    }

    if _, err := db.Exec("DROP TABLE IF EXISTS blog"); err != nil {
        return err
    }

    if _, err := db.Exec("CREATE TABLE IF NOT EXISTS blog (id int NOT NULL AUTO_INCREMENT, title varchar(255), PRIMARY KEY (id))"); err != nil {
        return err
    }

    for i := 0; i < 5; i++ {
        if _, err := db.Exec("INSERT INTO blog (title) VALUES (?);", fmt.Sprintf("Blog post #%d", i)); err != nil {
            return err
        }
    }
    return nil
}
