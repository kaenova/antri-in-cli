package main

import (
	"antri-in-cli/requirement"
	"antri-in-cli/utils"
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

var (
	OS_NAME_WINDOWS = "windows"
	OS_NAME_LINUX   = "linux"
	PORT_80         = "80"
	PORT_80_443     = "80, 443"
)

func main() {
	var (
		osName string = ""
		// portUsed        string = ""
		// ngrokAuth       string = ""
		// localOrServerIP string = ""
		// domainName      string = ""
	)

	myFigure := figure.NewFigure("Antri.In", "5lineoblique", true)
	myFigure.Print()
	fmt.Println("========")
	fmt.Println("Script Instalasi Antri.In")
	fmt.Println("by Yaudahlah Teams")
	fmt.Println("(Masukkan link github antri-in)")
	fmt.Println("========")
	fmt.Println("Memeriksa Prasyarat")

	// Periksa OS
	data, _ := utils.SpinnerCheck("Memeriksa OS", requirement.CheckOS)
	osName = data.(string)
	fmt.Println(osName) // Nanti dihapus

	// Periksa Port 80
	data1, _ := utils.SpinnerCheck("Memeriksa Ketersediaan Port 80", requirement.CheckPort80)
	port1 := data1.(string)
	fmt.Println(port1) // Nanti dihapus

	// Periksa Port 80
	data2, _ := utils.SpinnerCheck("Memeriksa Ketersediaan Port 443", requirement.CheckPort443)
	port2 := data2.(string)
	fmt.Println(port2) // Nanti dihapus

	// Periksa Docker
	utils.SpinnerCheck("Memeriksa Ketersediaan Port 80", requirement.CheckDocker)

	// Periksa Docker-Compose
	utils.SpinnerCheck("Memeriksa Ketersediaan Port 80", requirement.CheckDocker)
}
