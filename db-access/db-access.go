package main

// https://golang.google.cn/doc/tutorial/database-access

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

type Album struct {
	ID     int
	Title  string
	Artist string
	Price  float32
}

func albumsByArtist(name string) ([]Album, error) {
	var albums []Album
	sqlStr := "select * from album where artist=?"
	rows, err := db.Query(sqlStr, name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()

	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	if albums == nil {
		return nil, fmt.Errorf("albumsByArtist %q: Not found", name)
	}

	return albums, nil
}

func albumByID(id int) (Album, error) {
	var alb Album

	sqlStr := "select * from album where id=?"
	row := db.QueryRow(sqlStr, id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumByID %d: no such album", id)
		}
		return alb, fmt.Errorf("albumByID %d: %v", id, err)
	}

	return alb, nil
}

func main() {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "golab",
	}
	var err error
	// Open 不一定立即连接数据库，取决于驱动的实现
	// 所以，用 Ping 确认连接
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")

	artist := "Jackson"
	artist = "John Coltrane"
	albums, err := albumsByArtist(artist)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	alb, err := albumByID(12)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", alb)
}
