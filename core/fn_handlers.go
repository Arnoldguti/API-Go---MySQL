package core

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"database/sql"
	"fmt"
	"bytes"
	"net/http"
)

	//Connection to database server
	var db, err = sql.Open("mysql", InitConfig())
 func test() {
		if err != nil {
			fmt.Print(err.Error())
		}
		defer db.Close()

		// Connection available
		err = db.Ping()
		if err != nil {
			fmt.Print(err.Error())

		}

	}

 func GetPost (c *gin.Context) {
		var (
			post Data
			result gin.H
		)
		id := c.Param("id")
		row := db.QueryRow("select id, description, username, element_id,  date_added from posts where id = ?;", id)
		err = row.Scan(&post.Id, &post.Description, &post.Username, &post.Element, &post.Date)

		if err != nil {
			// If no results send null
			result = gin.H{
				"result": nil,
				"count":  0,
			}
		} else {
			result = gin.H{
				"result": post,
				//"count":  1,
			}
		}
		//	c.JSON(http.StatusOK, result)
		c.Render(1, render.IndentedJSON{Data: result})
	}

 func GetPosts(c *gin.Context) {
	var (
		post  Data
		posts []Data
	)
	rows, err := db.Query("select id, description, username, element_id, date_added from posts;")
	if err != nil {
		fmt.Print(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&post.Id, &post.Description, &post.Username, &post.Element, &post.Date)
		posts = append(posts, post)

		if err != nil {
			fmt.Print(err.Error())
		}
	}
	defer rows.Close()


	c.Render(1, render.IndentedJSON{Data: posts})

}

 func SetPost(c *gin.Context) {
	var buffer bytes.Buffer
	description := c.PostForm("description")
	username := c.PostForm("username")
	 element_id := c.PostForm("element_id")
	 date_added := c.PostForm("date_added")

	stmt, err := db.Prepare("insert into posts (description, username, element_id, date_added) values(?,?);")
	if err != nil {
		fmt.Print(err.Error())
	}

	_, err = stmt.Exec(description, username, element_id, date_added)
	if err != nil {
		fmt.Print(err.Error())
	}

	// Fastest way to append strings
	buffer.WriteString("Comment")
	buffer.WriteString(" ")
	buffer.WriteString("is")
	defer stmt.Close()
	name_created := buffer.String()
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf(" %s successfully created", name_created),
	})
}

func DelPost(c *gin.Context) {
	id := c.Query("id")
	stmt, err := db.Prepare("delete from posts where id= ?;")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully deleted comment"),
	})
}