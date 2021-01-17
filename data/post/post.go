package post

import (
	"fmt"
	"nixLevelFour/datasource/mysqldb"
)

var (
	dbName = "blog"
	insertPostQuery = "INSERT INTO posts (user_id,id,title,body) VALUES (?,?,?,?)"
)


type Post struct {
	UserID int		`json:"userId"`
	ID int			`json:"id"`
	Title string	`json:"title"`
	Body string 	`json:"body"`
}

func (p Post) Save() {
	db, err := mysqldb.InitDBConnection(dbName)
	if err != nil {
		fmt.Printf("Error connecting to data source: %s \n",err)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare(insertPostQuery)
	if err != nil {
		fmt.Printf("Error preparing statement: %s \n",err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.UserID,p.ID,p.Title,p.Body)
	if err != nil {
		fmt.Printf("Error executing INSERT Query: %s \n",err)
		return
	}
	return
}

