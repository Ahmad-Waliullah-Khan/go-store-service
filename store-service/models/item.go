package models

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:PASSWORD@localhost/beego-store?sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database: ", err)
	}
	log.Println("Database connection successful")
}

type Item struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func AddItem(item Item) error {
	_, err := db.Exec("INSERT INTO items (name, price) VALUES ($1, $2)", item.Name, item.Price)
	return err
}

func RemoveItem(id int) error {
	_, err := db.Exec("DELETE FROM items WHERE id = $1", id)
	return err
}

func ListItems(page, pageSize int) ([]Item, error) {
	rows, err := db.Query("SELECT id, name, price FROM items LIMIT $1 OFFSET $2", pageSize, (page-1)*pageSize)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.ID, &item.Name, &item.Price)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func GetItem(id int) (Item, error) {
	var item Item
	err := db.QueryRow("SELECT id, name, price FROM items WHERE id = $1", id).Scan(&item.ID, &item.Name, &item.Price)
	if err != nil {
		return Item{}, err
	}
	return item, nil
}

func UpdateItem(id int, updatedItem Item) error {
	_, err := db.Exec("UPDATE items SET name = $1, price = $2 WHERE id = $3", updatedItem.Name, updatedItem.Price, id)
	return err
}
