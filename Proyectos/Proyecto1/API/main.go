package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strconv"
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

type CpuInfo struct {
	Cpu_total      uint64 `json:"cpu_total"`
	Cpu_en_uso     uint64 `json:"cpu_en_uso"`
	Cpu_porcentaje uint64 `json:"cpu_porcentaje"`
	Fecha          string `json:"fecha"`
}

type ListProcess struct {
	Process []Process `json:"processes"`
}

type Process struct {
	Pid   int            `json:"pid"`
	Name  string         `json:"name"`
	User  int            `json:"user"`
	State int            `json:"state"`
	Ram   int            `json:"ram"`
	Child []ChildProcess `json:"child"`
}

type ChildProcess struct {
	Pid      int    `json:"pid"`
	Name     string `json:"name"`
	State    int    `json:"state"`
	PidPadre int    `json:"pidPadre"`
}

var db *sql.DB
var ramDataChan = make(chan RamInfo)

// var cpuDataChan = make(chan CpuInfo)
var cpuDataChan = make(chan string)

func main() {

	username := "user123"
	password := "user_password123"
	hostname := "mysql"
	portDB := "3306"
	dbname := "sample_db"

	// Initialize the database connection
	initDB(username, password, hostname, portDB, dbname)

	defer db.Close()

	go assignToChannel(ramDataChan)
	go assignToChannelCpu(cpuDataChan)

	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/api/insertRam", handleRAM)
	router.GET("/api/getRam", getDataRamHandler)
	router.GET("/api/insertCpu", CpuHandler)
	router.GET("/api/getCpu", getDataCpuHandler)
	//PROCESS CREATIONS
	router.GET("/api/createProcess", createProcess)
	router.GET("/api/stopProcess", stopProcess)
	router.GET("/api/resumeProcess", resumeProcess)
	router.GET("/api/terminateProcess", terminateProcess)

	//LIST OF PROCESS
	router.GET("/api/listProcess", listProcessCpu)

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

/*METHODS FOR CPU USAGE*/

func assignToChannelCpu(ch chan string) {
	for {
		cmd := exec.Command("sh", "-c", "cat /proc/modulo_cpu")
		output, err := cmd.Output()
		if err != nil {
			fmt.Println("Error executing command")
			return
		}

		// Unmarshal the JSON output into RamInfo struct
		/*var cpuInfo CpuInfo
		err = json.Unmarshal(output, &cpuInfo)

		if err != nil {
			fmt.Println("Error unmarshalling JSON")
			return
		}
		log.Println("getting data from channel", cpuInfo)*/
		// Send RamInfo to the channel
		ch <- string(output)
		time.Sleep(500 * time.Millisecond)
	}

}

func CpuHandler(c *gin.Context) {
	cpuInfo := <-cpuDataChan

	// Unmarshal the JSON output into RamInfo struct
	var cpuData CpuInfo
	err := json.Unmarshal([]byte(cpuInfo), &cpuData)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error unmarshalling JSON"})
		return
	}

	if cpuData != (CpuInfo{}) {
		// perform the time of the transaction
		cpuData.Fecha = time.Now().Format("2006-01-02 15:04:05")
		insertSQL := "INSERT INTO cpu_module (percentage_used, created_at) VALUES (?,?)"
		log.Println(cpuData)
		_, err := db.Exec(insertSQL, cpuData.Cpu_porcentaje, cpuData.Fecha)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting data into database"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": cpuData})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting data from channel"})

	}

}

func reverseArrayCpu(arr []CpuInfo) []CpuInfo {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func getDataCpuHandler(c *gin.Context) {
	// Perform a SELECT query
	rows, err := db.Query("SELECT percentage_used, created_at FROM cpu_module ORDER BY id DESC LIMIT 10")
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	// Iterate over the result set and collect data
	var dataRecords []CpuInfo
	for rows.Next() {
		var record CpuInfo
		if err := rows.Scan(&record.Cpu_porcentaje, &record.Fecha); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning database rows"})
			return
		}
		dataRecords = append(dataRecords, record)
	}

	// Reverse the array
	dataRecords = reverseArrayCpu(dataRecords)

	c.JSON(http.StatusOK, gin.H{
		"data": dataRecords,
	})
}

/*Method for CREATE, STOP, TERMINATE PROCESS*/

func createProcess(c *gin.Context) {
	cmd := exec.Command("sleep", "infinity")
	err := cmd.Start()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating process"})
		return
	}

	pid := cmd.Process.Pid
	pidConv := strconv.Itoa(pid)

	c.JSON(http.StatusOK, gin.H{"message": pidConv})

}

func stopProcess(c *gin.Context) {
	pid := c.Request.URL.Query().Get("pid")

	if pid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'pid' parameter"})
		return
	}

	//check if pid is a number
	validPid, err := strconv.Atoi(pid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'pid' parameter"})
		return
	}

	cmd := exec.Command("kill", "-SIGSTOP", strconv.Itoa(validPid))
	err = cmd.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error stopping process"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Process stopped with pid: " + pid})
}

func resumeProcess(c *gin.Context) {
	pid := c.Request.URL.Query().Get("pid")

	if pid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'pid' parameter"})
		return
	}

	//check if pid is a number
	validPid, err := strconv.Atoi(pid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'pid' parameter"})
		return
	}

	cmd := exec.Command("kill", "-SIGCONT", strconv.Itoa(validPid))
	err = cmd.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error resuming process"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Process resumed with pid: " + pid})
}

func terminateProcess(c *gin.Context) {
	pid := c.Request.URL.Query().Get("pid")

	if pid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'pid' parameter"})
		return
	}

	//check if pid is a number
	validPid, err := strconv.Atoi(pid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'pid' parameter"})
		return
	}

	cmd := exec.Command("kill", "-9", strconv.Itoa(validPid))
	err = cmd.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error terminating process"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Process terminated with pid: " + pid})
}

// Handle List of Process
func listProcessCpu(c *gin.Context) {
	cpuInfo := <-cpuDataChan

	// Unmarshal the JSON into list of process
	var listProcess ListProcess
	err := json.Unmarshal([]byte(cpuInfo), &listProcess)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error unmarshalling JSON"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": listProcess})

	//

}
