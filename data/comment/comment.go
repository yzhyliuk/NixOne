package comment

import (
	"fmt"
	"nixLevelFour/datasource/mysqldb"
)
var (
	dbName = "blog"
	insertCommentQuery = "INSERT INTO comments (post_id,id,name,email,body) VALUES (?,?,?,?,?)"
)

type Comment struct {
	PostID	int		`json:"postId"`
	ID 		int		`json:"id"`
	Name 	string	`json:"name"`
	Email 	string	`json:"email"`
	body 	string	`json:"body"`
}


func (c Comment) Save() {
	db, err := mysqldb.InitDBConnection(dbName)
	if err != nil {
		fmt.Printf("Error connecting to data source: %s \n",err)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare(insertCommentQuery)
	if err != nil {
		fmt.Printf("Error preparing statement: %s \n",err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(c.PostID,c.ID,c.Name,c.Email,c.body)
	if err != nil {
		fmt.Printf("Error executing INSERT Query: %s \n",err)
		return
	}
	return
}
