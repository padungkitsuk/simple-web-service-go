package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/microsoft/go-mssqldb"
	"database/sql"
	"context"
	"log"
	"fmt"
	// "errors"
)

var db *sql.DB
var server = "hfympcdb.database.windows.net"
var port = 1433
var user = "streammpc"
var password = "P@ssw0rd"
var database = "hfylnssqdb001dev"

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/pong", func(c *gin.Context) {
		conn()
		var rw,err = Read()

		fmt.Printf("rows %d err %s",rw,err)


		c.JSON(http.StatusOK, gin.H{
			"message": "ping",
		})
	})

	

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func conn(){
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
	server, user, password, port, database)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
	log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
	log.Fatal(err.Error())
	}
	fmt.Printf("Connected!")
}

// Read records
func Read() (int, error) {
	
    ctx := context.Background()

    // Check if database is alive.
    err := db.PingContext(ctx)
    if err != nil {
        return -1, err
    }
	
    tsql := fmt.Sprintf("SELECT CODE, VALUE FROM SYSTEM_CONFIGURATION;")

	
    // Execute query
    rows, err := db.QueryContext(ctx, tsql)
    if err != nil {
        return -1, err
    }
	
    defer rows.Close()

    var count int
	
    // Iterate through the result set.
    for rows.Next() {
        var code, value string
        // var id int
		
        // Get values from row.
        err := rows.Scan(&code, &value)
        if err != nil {
            return -1, err
        }

        fmt.Printf("code: %s, value: %s\n", code, value)
        count++
    }
	fmt.Printf("Read Success!\n")
    return count, nil
}
