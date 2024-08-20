package app

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func NewDB() *sql.DB {
	fmt.Println("Get Connection Database Running")
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go_restful?parseTime=true")
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected!!")
	db.SetConnMaxLifetime(time.Minute * 30)
	db.SetConnMaxIdleTime(time.Minute * 3)
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(10)
	return db
}

func StartServer(router http.Handler) {
	fmt.Println("Server running on : " + BaseUrl)
	server := http.Server{
		Addr: Host + ":" + strconv.Itoa(Port),
		// Handler: mux,
		Handler: router,
	}
	server.ListenAndServe()
}
