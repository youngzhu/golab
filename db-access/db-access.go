package main

// https://golang.google.cn/doc/tutorial/database-access

import (
	"database/sql"
	"fmt"
	"github.com/youngzhu/golab/db-access/handle"
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

func AlbumByID(id int) (Album, error) {
	stmt, err := db.Prepare("select * from album where id=?")
	if err != nil {
		log.Fatal(err)
	}

	var alb Album

	err = stmt.QueryRow(id).Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumByID %d: no such album", id)
		}
		return alb, fmt.Errorf("albumByID %d: %v", id, err)
	}

	return alb, nil
}

// 如果插入成功，则返回新记录的ID
// 否则返回0
func addAlbum(alb Album) (int, error) {
	insert := "insert into album (title, artist, price) values (?, ?, ?)"
	result, err := db.Exec(insert, alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	// 获得新记录的ID
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return int(id), nil
}

func main() {
	var err error

	//db, err = handle.OpenWithString()
	db, err = handle.OpenWithProperties()
	// 错误：unknown collation
	// 不知道为啥。。
	//db, err = handle.OpenWithConnector()
	if err != nil {
		log.Fatal(err)
	}

	// Open 不一定立即连接数据库，取决于驱动的实现
	// 所以，用 Ping 确认连接
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

	alb, err := albumByID(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", alb)

	albID, err := addAlbum(Album{
		Title:  "Cafe or Tea",
		Artist: "Andy Lau",
		Price:  99.99,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albID)

	alb, err = AlbumByID(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", alb)
}
