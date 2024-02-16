package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
)

// App struct
type App struct {
	ctx context.Context
}

type RamInfo struct {
	TotalRam     uint64 `json:"totalRam"`
	MemoriaEnUso uint64 `json:"memoriaEnUso"`
	Porcentaje   uint64 `json:"porcentaje"`
	Libre        uint64 `json:"libre"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Sum(augend int, addend int) string {
	return fmt.Sprintf("The sum of %d and %d is %d", augend, addend, augend+addend)
}

func (a *App) ReadRam() RamInfo {

	cmd := exec.Command("sh", "-c", "cat /proc/modulo_ram")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return RamInfo{}
	}

	// Unmarshal the JSON output into RamInfo struct
	var ramInfo RamInfo
	err = json.Unmarshal(output, &ramInfo)

	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return RamInfo{}
	}

	fmt.Printf("Total RAM: %d\n", ramInfo.TotalRam)
	fmt.Printf("Memory in use: %d\n", ramInfo.MemoriaEnUso)
	fmt.Printf("Percentage used: %d%%\n", ramInfo.Porcentaje)
	fmt.Printf("Free memory: %d\n", ramInfo.Libre)

	return ramInfo

}
