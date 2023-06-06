package main

import (
	"database/sql"
	"fmt"
	"net/http"

	//"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

var DB *sql.DB

type Link struct {
	ID        int       `json:"id"`
	URL       string    `json:"url"`
	Domain    string    `json:"domain"`
	CreatedAt time.Time `json:"date"`
	ValidUrl  bool      `json:"validurl"`
}

func main() {
	createDBConnection()
	defer DB.Close()
	r := gin.Default()
	setupRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
func setupRoutes(r *gin.Engine) {

	r.POST("/url", SaveLongLink)
	// r.GET("/:id", redirectHandler)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

}

// POST
func SaveLongLink(c *gin.Context) {
	reqBody := Link{}
	err := c.Bind(&reqBody)
	if err != nil {
		res := gin.H{
			"error": "invalid request body",
		}
		c.Writer.Header().Set("Content-Type", "application/json")
		c.JSON(http.StatusBadRequest, res)

		return
	}
	reqBody.CreatedAt = time.Now()
	reqBody.Domain, err = getdomain(reqBody.URL)
	if err != nil {
		// res := gin.H{
		// 	"error": "invalid url",
		// }
		// c.Writer.Header().Set("Content-Type", "application/json")
		// c.JSON(http.StatusBadRequest, res)
		// return
		reqBody.ValidUrl = false
	} else {
		reqBody.ValidUrl = true
	}
	//reqBody.ValidUrl = validurl(reqBody.URL)

	// Data[lastID] = reqBody
	fmt.Println(reqBody)
	res, err := DB.Exec(`INSERT INTO "url" ("url", "domain", "created_at", "validurl")
	VALUES ( $1, $2, $3, $4)`, reqBody.URL, reqBody.Domain, reqBody.CreatedAt, reqBody.ValidUrl)
	if err != nil {
		fmt.Println("err inserting data: ", err)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	lastInsID, _ := res.LastInsertId()
	reqBody.ID = int(lastInsID)
	fmt.Println("res: ", lastInsID)
	c.JSON(http.StatusOK, reqBody)
	c.Writer.Header().Set("Content-Type", "application/json")
}
