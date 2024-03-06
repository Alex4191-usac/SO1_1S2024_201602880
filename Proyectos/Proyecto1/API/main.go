package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Message struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type DataRecord struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var db *sql.DB

func main() {

	username := "user123"
	password := "user_password123"
	hostname := "localhost"
	portDB := "3306"
	dbname := "sample_db"

	// Initialize the database connection
	initDB(username, password, hostname, portDB, dbname)

	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/ram", infoRAMHandler)
	router.GET("/cpu", infoCPUHandler)
	router.GET("/data", getDataHandler)

	port := 8080
	router.Run(fmt.Sprintf(":%d", port))
}

func initDB(username, password, hostname, port, dbname string) {
	// Create a MySQL DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, dbname)

	// Open a connection to the database
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error opening database connection:", err)
		return
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging database:", err)
		return
	}

	fmt.Println("Connected to the MySQL database!")
}

func infoRAMHandler(c *gin.Context) {

	message := Message{
		Name: "getting RAM INFO",
		Id:   1,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func infoCPUHandler(c *gin.Context) {

	message := Message{
		Name: "getting CPU INFO",
		Id:   2,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func getDataHandler(c *gin.Context) {
	// Perform a SELECT query
	rows, err := db.Query("SELECT id, name FROM example_table")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying the database"})
		return
	}
	defer rows.Close()

	// Iterate over the result set and collect data
	var dataRecords []DataRecord
	for rows.Next() {
		var record DataRecord
		if err := rows.Scan(&record.ID, &record.Name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning database rows"})
			return
		}
		dataRecords = append(dataRecords, record)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": dataRecords,
	})
}
