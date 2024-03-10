package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type RamInfo struct {
	TotalRam     uint64 `json:"totalRam"`
	MemoriaEnUso uint64 `json:"memoriaEnUso"`
	Porcentaje   uint64 `json:"porcentaje"`
	Libre        uint64 `json:"libre"`
	Fecha        string `json:"fecha"`
}

var db *sql.DB
var ramDataChan = make(chan RamInfo)

func main() {

	username := "user123"
	password := "user_password123"
	hostname := "localhost"
	portDB := "3306"
	dbname := "sample_db"

	// Initialize the database connection
	initDB(username, password, hostname, portDB, dbname)

	defer db.Close()

	go assignToChannel(ramDataChan)

	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/ram", infoRAMHandler)
	router.GET("/insertRam", handleRAM)
	router.GET("/getRam", getDataRamHandler)

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

	cmd := exec.Command("sh", "-c", "cat /proc/modulo_ram")
	output, err := cmd.Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error executing command"})
	}

	// Unmarshal the JSON output into RamInfo struct
	var ramInfo RamInfo
	err = json.Unmarshal(output, &ramInfo)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error unmarshalling JSON"})
	}

	c.JSON(http.StatusOK, gin.H{
		"ramInfo": ramInfo,
	})
}

func assignToChannel(ch chan RamInfo) {
	for {
		cmd := exec.Command("sh", "-c", "cat /proc/modulo_ram")
		output, err := cmd.Output()
		if err != nil {
			fmt.Println("Error executing command")
			return
		}

		// Unmarshal the JSON output into RamInfo struct
		var ramInfo RamInfo
		err = json.Unmarshal(output, &ramInfo)
		if err != nil {
			fmt.Println("Error unmarshalling JSON")
			return
		}
		log.Println("getting data from channel", ramInfo)
		// Send RamInfo to the channel
		ch <- ramInfo
		time.Sleep(500 * time.Millisecond)
	}
}

func handleRAM(c *gin.Context) {
	ramData := <-ramDataChan

	if ramData != (RamInfo{}) {
		// perform the time of the transaction
		ramData.Fecha = time.Now().Format("2006-01-02 15:04:05")
		insertSQL := "INSERT INTO ram_module (total_memory, used_memory, free_memory, percentage_used, created_at) VALUES (?,?,?,?,?)"
		log.Println(ramData)
		_, err := db.Exec(insertSQL, ramData.TotalRam, ramData.MemoriaEnUso, ramData.Libre, ramData.Porcentaje, ramData.Fecha)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting data into database"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": ramData})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting data from channel"})

	}

}

/*func insertRamHandler(c *gin.Context) {

	cmd := exec.Command("sh", "-c", "cat /proc/modulo_ram")
	output, err := cmd.Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error executing command"})
	}

	// Unmarshal the JSON output into RamInfo struct
	var ramInfo RamInfo
	err = json.Unmarshal(output, &ramInfo)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error unmarshalling JSON"})
	}

	// Perform an INSERT query

	result, err := db.Exec("INSERT INTO ram_module (total_memory, used_memory, free_memory, percentage_used) VALUES (?,?,?,?)", ramInfo.TotalRam, ramInfo.MemoriaEnUso, ramInfo.Libre, ramInfo.Porcentaje)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting data into the database"})
		return
	}

	// Get the ID of the inserted record
	insertedID, _ := result.LastInsertId()

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Data inserted with ID: %d", insertedID),
	})

}*/

func getDataRamHandler(c *gin.Context) {
	// Perform a SELECT query
	rows, err := db.Query("SELECT total_memory, used_memory, free_memory, percentage_used, created_at FROM ram_module ORDER BY id DESC LIMIT 10")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying the database"})
		return
	}
	defer rows.Close()

	// Iterate over the result set and collect data
	var dataRecords []RamInfo
	for rows.Next() {
		var record RamInfo
		if err := rows.Scan(&record.TotalRam, &record.MemoriaEnUso, &record.Libre, &record.Porcentaje, &record.Fecha); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning database rows"})
			return
		}
		dataRecords = append(dataRecords, record)
	}

	// Reverse the array
	dataRecords = reverseArray(dataRecords)

	c.JSON(http.StatusOK, gin.H{
		"data": dataRecords,
	})
}

func reverseArray(arr []RamInfo) []RamInfo {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
